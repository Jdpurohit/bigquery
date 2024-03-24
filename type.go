package main

import (
	"encoding/json"
	"fmt"
)

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
	CreativeSize string `json:"cs"`
}

type LogsBDTable struct {
	TY            string  `json:"ty"`
	IntegrationId int64   `json:"iid,string"`
	ConfigId      int64   `json:"cid,string"`
	Device        string  `json:"dv"`
	CreativeSize  string  `json:"cs"`
	Partner       string  `json:"ptn"`
	Revenue       float64 `json:"rev,string"`
	Currency      string  `json:"cur"`
	S2S           bool    `json:"s2s"`
}

type Payload struct {
	Entries []map[string]interface{}
}

// UnmarshalJSON implements json.Unmarshaler
func (p *Payload) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return fmt.Errorf("no bytes to unmarshal")
	}
	// See if we can guess based on the first character
	switch b[0] {
	case '{':
		return p.unmarshalSingle(b)
	case '[':
		return p.unmarshalMany(b)
	}

	return fmt.Errorf("unexpected payload")
}

func (p *Payload) unmarshalSingle(b []byte) error {
	var e map[string]interface{}
	err := json.Unmarshal(b, &e)
	if err != nil {
		return err
	}
	p.Entries = []map[string]interface{}{e}
	return nil
}

func (p *Payload) unmarshalMany(b []byte) error {
	var entries []map[string]interface{}
	err := json.Unmarshal(b, &entries)
	if err != nil {
		return err
	}
	p.Entries = entries
	return nil
}
