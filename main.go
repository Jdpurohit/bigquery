package main

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/bigquery/storage/managedwriter"
	"github.com/Jdpurohit/bigquery/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/types/descriptorpb"
)

var (
	projectID = "ci-infrastructure"
	datasetID = "cipt"
)

func main() {
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		handleOptions(w, r)
	case http.MethodPost:
		handlePost(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
	w.WriteHeader(http.StatusOK)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	logDate := time.Now().UTC().Format("2006-01-02 15:04:05")
	hash := getHash(r)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var pLoad Payload
	err = json.Unmarshal(body, &pLoad)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	xAppEngineCountry := "unknown"
	if r.Header.Get("X-Appengine-Country") != "" {
		xAppEngineCountry = r.Header.Get("X-Appengine-Country")
	}

	var (
		bdLogs []*pb.LogsBDTable
		hbLogs []*pb.LogsHBTable
		fcLogs []*pb.LogsFCTable
		gnLogs []*pb.LogsTable
	)

	for _, entry := range pLoad.Entries {
		if entry["ev"] == "hb" {
			hbRecord := hbLogging(entry, xAppEngineCountry)
			bdRecord := bdLogging(entry, xAppEngineCountry)
			if hbRecord != nil {
				hbLogs = append(hbLogs, hbRecord)
			}
			if bdRecord != nil {
				bdLogs = append(bdLogs, bdRecord)
			}
		} else if entry["ev"] == "fc" {
			fcRecord := fcLogging(entry, xAppEngineCountry)
			if fcRecord != nil {
				fcLogs = append(fcLogs, fcRecord)
			}
		} else if entry["ev"] != nil {
			gnRecord := gnLoggingV2(entry, logDate, hash)
			if gnRecord != nil {
				gnLogs = append(gnLogs, gnRecord)
			}
		} else {
			gnRecord := gnLogging(entry, logDate, hash)
			if gnRecord != nil {
				gnLogs = append(gnLogs, gnRecord)
			}
		}
	}

	log.Println("initializing managed writer new client")
	ctx := context.Background()
	client, err := managedwriter.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// ==>> LogsTable

	var gnLogsData [][]byte
	for k, mesg := range gnLogs {
		b, err := proto.Marshal(mesg)
		if err != nil {
			log.Printf("gnLogs: failed to marshal message %d: %v", k, err)
		}
		gnLogsData = append(gnLogsData, b)
	}

	if len(gnLogs) > 0 {
		log.Printf("pushing %d records to gnLogs table", len(gnLogs))
		gn := &pb.LogsTable{}
		if err := saveRecords(ctx, client, protodesc.ToDescriptorProto(gn.ProtoReflect().Descriptor()), datasetID, "logs", gnLogsData); err != nil {
			log.Println("Error saving general logs:", err)
		}
	}

	// ==>> LogsFCTable

	var fcLogsData [][]byte
	for k, mesg := range fcLogs {
		b, err := proto.Marshal(mesg)
		if err != nil {
			log.Printf("fcLogs: failed to marshal message %d: %v", k, err)
		}
		fcLogsData = append(fcLogsData, b)
	}

	if len(fcLogs) > 0 {
		log.Printf("pushing %d records to fcLogs table", len(fcLogs))
		fc := &pb.LogsFCTable{}
		if err := saveRecords(ctx, client, protodesc.ToDescriptorProto(fc.ProtoReflect().Descriptor()), datasetID, "logs_fc", fcLogsData); err != nil {
			log.Println("Error saving fallback creative logs:", err)
		}
	}

	// ==>> LogsHBTable

	var hbLogsData [][]byte
	for k, mesg := range hbLogs {
		b, err := proto.Marshal(mesg)
		if err != nil {
			log.Printf("hbLogs: failed to marshal message %d: %v", k, err)
		}
		hbLogsData = append(hbLogsData, b)
	}

	if len(hbLogs) > 0 {
		log.Printf("pushing %d records to hbLogs table", len(hbLogs))
		hb := &pb.LogsHBTable{}
		if err := saveRecords(ctx, client, protodesc.ToDescriptorProto(hb.ProtoReflect().Descriptor()), datasetID, "logs_hb", hbLogsData); err != nil {
			log.Println("Error saving header bidding logs:", err)
		}
	}

	// ==>> LogsBDTable

	var bdLogsData [][]byte
	for k, mesg := range bdLogs {
		b, err := proto.Marshal(mesg)
		if err != nil {
			log.Printf("bdLogs: failed to marshal message %d: %v", k, err)
		}
		bdLogsData = append(bdLogsData, b)
	}

	if len(bdLogs) > 0 {
		log.Printf("pushing %d records to bdLogs table", len(bdLogs))
		bd := &pb.LogsBDTable{}
		if err := saveRecords(ctx, client, protodesc.ToDescriptorProto(bd.ProtoReflect().Descriptor()), datasetID, "bid_data", bdLogsData); err != nil {
			log.Println("Error saving bid data logs:", err)
		}
	}
	fmt.Fprint(w, "")
}

func saveRecords(ctx context.Context, client *managedwriter.Client, dp *descriptorpb.DescriptorProto, datasetID, tableID string, rowsData [][]byte) error {
	// setup a new stream.
	ms, err := client.NewManagedStream(ctx,
		managedwriter.WithDestinationTable(managedwriter.TableParentFromParts(projectID, datasetID, tableID)),
		managedwriter.WithType(managedwriter.DefaultStream),
		managedwriter.WithSchemaDescriptor(dp),
	)
	if err != nil {
		return fmt.Errorf("saveRecords: NewManagedStream: %v", err)
	}
	defer ms.Close()

	result, err := ms.AppendRows(ctx, rowsData)
	if err != nil {
		return fmt.Errorf("saveRecords: grouped-row append failed: %v", err)
	}

	res, err := result.FullResponse(ctx)
	if err != nil {
		return fmt.Errorf("saveRecords: FullResponse error: %v", err)
	}
	log.Printf("saveRecords: Row Wise Error: %v", res.GetRowErrors())

	_, err = result.GetResult(ctx)
	return err
}

func getHash(r *http.Request) string {
	aip := getIP(r)
	userAgent := r.Header.Get("User-Agent")
	setHash := fmt.Sprintf("CI://%s%s", aip, userAgent)
	hash := md5.Sum([]byte(setHash))
	return fmt.Sprintf("%x", hash)[:9]
}

func getIP(r *http.Request) string {
	if xUserIP := r.Header.Get("X-Appengine-User-IP"); xUserIP != "" {
		return xUserIP
	}
	if xForwardedFor := r.Header.Get("X-Forwarded-For"); xForwardedFor != "" {
		return strings.Split(xForwardedFor, ",")[0]
	}
	if clientIP := r.Header.Get("Client-Ip"); clientIP != "" {
		return clientIP
	}
	if realIP := r.Header.Get("X-Real-Ip"); realIP != "" {
		return realIP
	}
	return r.RemoteAddr
}

func gnLoggingV2(entry map[string]interface{}, logDate, hash string) *pb.LogsTable {
	log.Println("gnLoggingV2 entry: ", entry)
	jsonString, err := json.Marshal(entry)
	if err != nil {
		log.Printf("Unable to marshal JSON due to %s", err)
		return nil
	}
	var target LogsTable
	err = json.Unmarshal(jsonString, &target)
	if err != nil {
		log.Printf("Unable to unmarshal JSON due to %s", err)
		return nil
	}

	log.Println("gnLoggingV2 target: ", target)
	uid, err := target.UId.Int64()
	if err != nil {
		return nil
	}
	cid, _ := target.CId.Int64() // non-mandatory field
	if uid == 0 || target.URL == "" || target.EV == "" {
		return nil
	}

	eventMap := map[string]string{
		"ai":   "ad_impression",
		"al":   "ad_load",
		"als":  "ad_load_saas",
		"alt":  "ad_latency",
		"am":   "ad_miss",
		"ams":  "ad_miss_saas",
		"aul":  "ad_unit_load",
		"auls": "ad_unit_load_saas",
		"cad":  "mobile_ci_ad_density",
		"cd":   "consent_denied",
		"cg":   "consent_given",
		"cnf":  "cmp_not_found",
		"dl":   "dom_latency",
		"ha":   "heavy_ad",
		"oad":  "mobile_other_ad_density",
		"pi":   "page_impression",
		"pl":   "page_latency",
		"te":   "tag_error",
		"tl":   "tag_latency",
		"ul":   "unit_latency",
		"vc":   "ag_video_click",
		"vi":   "ag_video_impression",
	}
	event, ok := eventMap[target.EV]
	if !ok {
		return nil
	}

	row := &pb.LogsTable{
		Uid:      uid,
		Datetime: logDate,
		Hash:     hash,
		Url:      target.URL,
		Event:    event,
	}

	// Add optional fields here...
	if target.UN != "" {
		row.Unit = target.UN
		parts := strings.Split(target.UN, "__")
		if len(parts) > 1 {
			parts = parts[:len(parts)-1]
		}
		row.UnitName = strings.Join(parts, "__")
	}
	if cid != 0 {
		row.ConfigId = cid
	}
	if target.Dt != "" {
		row.Details = target.Dt
	}
	return row
}

func gnLogging(entry map[string]interface{}, logDate, hash string) *pb.LogsTable {
	log.Println("gnLogging entry: ", entry)
	jsonString, err := json.Marshal(entry)
	if err != nil {
		log.Printf("Unable to marshal JSON due to %s", err)
		return nil
	}
	var target LogsTable
	err = json.Unmarshal(jsonString, &target)
	if err != nil {
		log.Printf("Unable to unmarshal JSON due to %s", err)
		return nil
	}

	log.Println("gnLogging target: ", target)
	uid, err := target.UId.Int64()
	if err != nil {
		return nil
	}
	cid, _ := target.ConfigId.Int64() // non-mandatory field
	if uid == 0 || target.URL == "" || target.Event == "" {
		return nil
	}

	row := &pb.LogsTable{
		Uid:      uid,
		Datetime: logDate,
		Hash:     hash,
		Url:      target.URL,
		Event:    target.Event,
	}

	// Add optional fields here...
	if target.Unit != "" {
		row.Unit = target.Unit
		parts := strings.Split(target.Unit, "__")
		if len(parts) > 1 {
			parts = parts[:len(parts)-1]
		}
		row.UnitName = strings.Join(parts, "__")
	}
	if cid != 0 {
		row.ConfigId = cid
	}
	if target.Details != "" {
		row.Details = target.Details
	}
	return row
}

func hbLogging(entry map[string]interface{}, country string) *pb.LogsHBTable {
	log.Println("hbLogging entry: ", entry)
	jsonString, err := json.Marshal(entry)
	if err != nil {
		log.Printf("Unable to marshal JSON due to %s", err)
		return nil
	}
	var target LogsHBTable
	err = json.Unmarshal(jsonString, &target)
	if err != nil {
		log.Printf("Unable to unmarshal JSON due to %s", err)
		return nil
	}
	log.Println("hbLogging target: ", target)

	iid, err := target.IntegrationId.Int64()
	if err != nil {
		return nil
	}
	cid, err := target.ConfigId.Int64()
	if err != nil {
		return nil
	}
	rev, _ := target.Revenue.Float64() // non-mandatory field
	if cid == 0 || iid == 0 || target.TY == "" {
		return nil
	}

	eventMap := map[string]string{
		"bw": "bid_won",
		"ar": "ad_rendered",
		"vi": "viewable_impression",
		"ac": "click",
	}
	event, ok := eventMap[target.TY]
	if !ok {
		return nil
	}

	if target.Device == "" {
		target.Device = "unknown"
	}
	if target.CreativeSize == "" {
		target.CreativeSize = "unknown"
	}
	if target.Partner == "" {
		target.Partner = "unknown"
	}
	if target.Currency == "" {
		target.Currency = "unknown"
	}
	return &pb.LogsHBTable{
		Date:          time.Now().UTC().Format("2006-01-02"),
		Event:         event,
		IntegrationId: iid,
		ConfigId:      cid,
		Device:        target.Device,
		Geo:           country,
		CreativeSize:  target.CreativeSize,
		Partner:       target.Partner,
		Revenue:       rev,
		Currency:      target.Currency,
		S2S:           target.S2S,
	}
}

func fcLogging(entry map[string]interface{}, country string) *pb.LogsFCTable {
	log.Println("fcLogging entry: ", entry)
	jsonString, err := json.Marshal(entry)
	if err != nil {
		log.Printf("Unable to marshal JSON due to %s", err)
		return nil
	}
	var target LogsFCTable
	err = json.Unmarshal(jsonString, &target)
	if err != nil {
		log.Printf("Unable to unmarshal JSON due to %s", err)
		return nil
	}

	log.Println("fcLogging target: ", target)
	creativeId, err := target.CreativeId.Int64()
	if err != nil {
		return nil
	}
	cid, err := target.ConfigId.Int64()
	if err != nil {
		return nil
	}
	if cid == 0 || creativeId == 0 || target.TY == "" {
		return nil
	}

	eventMap := map[string]string{
		"cl": "click",
		"im": "impression",
		"vi": "viewable_impression",
	}
	event, ok := eventMap[target.TY]
	if !ok {
		return nil
	}

	if target.Device == "" {
		target.Device = "unknown"
	}
	if target.CreativeSize == "" {
		target.CreativeSize = "unknown"
	}

	return &pb.LogsFCTable{
		Date:         time.Now().UTC().Format("2006-01-02"),
		Event:        event,
		ConfigId:     cid,
		Device:       target.Device,
		Geo:          country,
		CreativeSize: target.CreativeSize,
		CreativeId:   creativeId,
	}
}

func bdLogging(entry map[string]interface{}, country string) *pb.LogsBDTable {
	log.Println("bdLogging entry: ", entry)
	jsonString, err := json.Marshal(entry)
	if err != nil {
		log.Printf("Unable to marshal JSON due to %s", err)
		return nil
	}
	var target LogsBDTable
	err = json.Unmarshal(jsonString, &target)
	if err != nil {
		log.Printf("Unable to unmarshal JSON due to %s", err)
		return nil
	}
	log.Println("bdLogging target: ", target)
	iid, err := target.IntegrationId.Int64()
	if err != nil {
		return nil
	}
	cid, err := target.ConfigId.Int64()
	if err != nil {
		return nil
	}
	rev, _ := target.Revenue.Float64() // non-mandatory field
	if cid == 0 || iid == 0 || target.TY == "" {
		return nil
	}

	eventMap := map[string]string{
		"bd": "bid",
	}
	event, ok := eventMap[target.TY]
	if !ok {
		return nil
	}

	if target.Device == "" {
		target.Device = "unknown"
	}
	if target.CreativeSize == "" {
		target.CreativeSize = "unknown"
	}
	if target.Partner == "" {
		target.Partner = "unknown"
	}
	if target.Currency == "" {
		target.Currency = "unknown"
	}

	return &pb.LogsBDTable{
		Timestamp:     time.Now().UnixMicro(),
		Event:         event,
		IntegrationId: iid,
		ConfigId:      cid,
		Device:        target.Device,
		Geo:           country,
		CreativeSize:  target.CreativeSize,
		Partner:       target.Partner,
		Revenue:       rev,
		Currency:      target.Currency,
		S2S:           target.S2S,
	}
}
