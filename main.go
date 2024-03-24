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

	var entries []map[string]interface{}
	err = json.Unmarshal(body, &entries)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var (
		bdLogs []*pb.LogsBDTable
		hbLogs []*pb.LogsHBTable
		fcLogs []*pb.LogsFCTable
		gnLogs []*pb.LogsTable
	)

	for _, entry := range entries {
		if entry["ev"] == "hb" {
			hbRecord := hbLogging(entry)
			bdRecord := bdLogging(entry)
			if hbRecord != nil {
				hbLogs = append(hbLogs, hbRecord)
			}
			if bdRecord != nil {
				bdLogs = append(bdLogs, bdRecord)
			}
		} else if entry["ev"] == "fc" {
			fcRecord := fcLogging(entry)
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

	// Now, send the test rows grouped into in a single append for gnLogs.
	var gnLogsData [][]byte
	for k, mesg := range gnLogs {
		b, err := proto.Marshal(mesg)
		if err != nil {
			log.Printf("gnLogs: failed to marshal message %d: %v", k, err)
		}
		gnLogsData = append(gnLogsData, b)
	}

	log.Println("calling saveRecords for gnLogs")
	gn := &pb.LogsTable{}
	if err := saveRecords(ctx, client, protodesc.ToDescriptorProto(gn.ProtoReflect().Descriptor()), datasetID, "logs", gnLogsData); err != nil {
		log.Println("Error saving general logs:", err)
	}

	// Now, send the test rows grouped into in a single append for fcLogs.
	var fcLogsData [][]byte
	for k, mesg := range fcLogs {
		b, err := proto.Marshal(mesg)
		if err != nil {
			log.Printf("fcLogs: failed to marshal message %d: %v", k, err)
		}
		gnLogsData = append(gnLogsData, b)
	}

	log.Println("calling saveRecords for fcLogs")
	fc := &pb.LogsFCTable{}
	if err := saveRecords(ctx, client, protodesc.ToDescriptorProto(fc.ProtoReflect().Descriptor()), datasetID, "logs_fc", fcLogsData); err != nil {
		log.Println("Error saving fallback creative logs:", err)
	}

	// Now, send the test rows grouped into in a single append for hbLogs.
	var hbLogsData [][]byte
	for k, mesg := range hbLogs {
		b, err := proto.Marshal(mesg)
		if err != nil {
			log.Printf("hbLogs: failed to marshal message %d: %v", k, err)
		}
		gnLogsData = append(gnLogsData, b)
	}

	log.Println("calling saveRecords for hbLogs")
	hb := &pb.LogsHBTable{}
	if err := saveRecords(ctx, client, protodesc.ToDescriptorProto(hb.ProtoReflect().Descriptor()), datasetID, "logs_hb", hbLogsData); err != nil {
		log.Println("Error saving header bidding logs:", err)
	}

	// Now, send the test rows grouped into in a single append.
	var bdLogsData [][]byte
	for k, mesg := range bdLogs {
		b, err := proto.Marshal(mesg)
		if err != nil {
			log.Printf("bdLogs: failed to marshal message %d: %v", k, err)
		}
		bdLogsData = append(bdLogsData, b)
	}

	log.Println("calling saveRecords for bdLogs")
	bd := &pb.LogsBDTable{}
	if err := saveRecords(ctx, client, protodesc.ToDescriptorProto(bd.ProtoReflect().Descriptor()), datasetID, "bid_data", bdLogsData); err != nil {
		log.Println("Error saving bid data logs:", err)
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
		log.Println("saveRecords: NewManagedStream:", err)
		return err
	}

	result, err := ms.AppendRows(ctx, rowsData)
	if err != nil {
		log.Println("saveRecords: grouped-row append failed:", err)
		return err
	}

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
	requiredentry := []string{"ev", "uid", "url"}
	allPresent := enforceRequiredParams(requiredentry, entry)
	if !allPresent {
		return nil
	}

	if _, ok := entry["uid"].(int64); !ok {
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
	event, ok := eventMap[entry["ev"].(string)]
	if !ok {
		return nil
	}

	row := &pb.LogsTable{
		Uid:      entry["uid"].(int64),
		Datetime: logDate,
		Hash:     hash,
		Url:      entry["url"].(string),
		Event:    event,
	}

	// Add optional fields here...
	if unit, ok := entry["un"].(string); ok && unit != "" {
		row.Unit = unit
		parts := strings.Split(unit, "__")
		if len(parts) > 1 {
			parts = parts[:len(parts)-1]
		}
		row.UnitName = strings.Join(parts, "__")
	}
	if cid, ok := entry["cid"].(string); ok && cid != "" {
		row.ConfigId = cid
	}
	if dt, ok := entry["dt"].(string); ok && dt != "" {
		row.Details = dt
	}
	return row
}

func gnLogging(entry map[string]interface{}, logDate, hash string) *pb.LogsTable {
	requiredentry := []string{"event", "uid", "url"}
	allPresent := enforceRequiredParams(requiredentry, entry)
	if !allPresent {
		return nil
	}

	if _, ok := entry["uid"].(int64); !ok {
		return nil
	}

	row := &pb.LogsTable{
		Uid:      entry["uid"].(int64),
		Datetime: logDate,
		Hash:     hash,
		Url:      entry["url"].(string),
		Event:    entry["event"].(string),
	}

	// Add optional fields here...
	if unit, ok := entry["unit"].(string); ok && unit != "" {
		row.Unit = unit
		parts := strings.Split(unit, "__")
		if len(parts) > 1 {
			parts = parts[:len(parts)-1]
		}
		row.UnitName = strings.Join(parts, "__")
	}
	if cid, ok := entry["config_id"].(string); ok && cid != "" {
		row.ConfigId = cid
	}
	if dt, ok := entry["details"].(string); ok && dt != "" {
		row.Details = dt
	}
	return row
}

func hbLogging(entry map[string]interface{}) *pb.LogsHBTable {
	requiredentry := []string{"cid", "iid", "ty"}
	allPresent := enforceRequiredParams(requiredentry, entry)
	if !allPresent {
		return nil
	}

	eventMap := map[string]string{
		"bw": "bid_won",
		"ar": "ad_rendered",
		"vi": "viewable_impression",
		"ac": "click",
	}
	event, ok := eventMap[entry["ty"].(string)]
	if !ok {
		return nil
	}

	return &pb.LogsHBTable{
		Date:          time.Now().UTC().Format("2006-01-02"),
		Event:         event,
		IntegrationId: entry["iid"].(int64),
		ConfigId:      entry["cid"].(int64),
		Device:        entry["dv"].(string),
		Geo:           entry["geo"].(string),
		CreativeSize:  entry["cs"].(string),
		Partner:       entry["ptn"].(string),
		Revenue:       entry["rev"].(float64),
		Currency:      entry["cur"].(string),
		S2S:           entry["s2s"].(bool),
	}
}

func fcLogging(entry map[string]interface{}) *pb.LogsFCTable {
	requiredentry := []string{"cid", "ty", "fid"}
	allPresent := enforceRequiredParams(requiredentry, entry)
	if !allPresent {
		return nil
	}

	eventMap := map[string]string{
		"cl": "click",
		"im": "impression",
		"vi": "viewable_impression",
	}
	event, ok := eventMap[entry["ty"].(string)]
	if !ok {
		return nil
	}

	return &pb.LogsFCTable{
		Date:         time.Now().UTC().Format("2006-01-02"),
		Event:        event,
		ConfigId:     entry["cid"].(int64),
		Device:       entry["dv"].(string),
		Geo:          entry["geo"].(string),
		CreativeSize: entry["cs"].(string),
		CreativeId:   entry["fid"].(int64),
	}
}

func bdLogging(entry map[string]interface{}) *pb.LogsBDTable {
	requiredentry := []string{"cid", "iid", "ty"}
	allPresent := enforceRequiredParams(requiredentry, entry)
	if !allPresent {
		return nil
	}

	eventMap := map[string]string{
		"bd": "bid",
	}
	event, ok := eventMap[entry["ty"].(string)]
	if !ok {
		return nil
	}

	return &pb.LogsBDTable{
		Timestamp:     time.Now().UTC().Format("2006-01-02"),
		Event:         event,
		IntegrationId: entry["iid"].(int64),
		ConfigId:      entry["cid"].(int64),
		Device:        entry["dv"].(string),
		Geo:           entry["geo"].(string),
		CreativeSize:  entry["cs"].(string),
		Partner:       entry["ptn"].(string),
		Revenue:       entry["rev"].(float64),
		Currency:      entry["cur"].(string),
		S2S:           entry["s2s"].(bool),
	}
}

func enforceRequiredParams(requiredentry []string, fields map[string]interface{}) bool {
	for _, v := range requiredentry {
		val, ok := fields[v]
		if !ok || val == nil || val == "" {
			return false
		}
	}
	return true
}
