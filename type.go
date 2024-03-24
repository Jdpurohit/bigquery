package main

type LogsTable struct {
	UId      int64  `json:"uid,string"`
	URL      string `json:"url"`
	Event    string `json:"event"`
	Unit     string `json:"unit"`
	ConfigId int64  `json:"config_id,string"`
	Details  string `json:"details"`

	// below parameter are expect in general log v2 only
	EV  string `json:"ev"`
	UN  string `json:"un"`
	CId int64  `json:"cid,string"`
	Dt  string `json:"dt"`
}

type LogsHBTable struct {
	TY            string  `json:"ty"`
	IntegrationId int64   `json:"iid,string"`
	ConfigId      int64   `json:"cid,string"`
	Device        string  `json:"dv"`
	Geo           string  `json:"geo"`
	CreativeSize  string  `json:"cs"`
	Partner       string  `json:"ptn"`
	Revenue       float64 `json:"rev,string"`
	Currency      string  `json:"cur"`
	S2S           bool    `json:"s2s"`
}

type LogsFCTable struct {
	TY           string `json:"ty"`
	ConfigId     int64  `json:"cid,string"`
	CreativeId   int64  `json:"fid,string"`
	Device       string `json:"dv"`
	Geo          string `json:"geo"`
	CreativeSize string `json:"cs"`
}

type LogsBDTable struct {
	TY            string  `json:"ty"`
	IntegrationId int64   `json:"iid,string"`
	ConfigId      int64   `json:"cid,string"`
	Device        string  `json:"dv"`
	Geo           string  `json:"geo"`
	CreativeSize  string  `json:"cs"`
	Partner       string  `json:"ptn"`
	Revenue       float64 `json:"rev,string"`
	Currency      string  `json:"cur"`
	S2S           bool    `json:"s2s"`
}
