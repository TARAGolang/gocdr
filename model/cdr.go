package model

import "time"

type CDR struct {
	Id             string      `json:"id"               bson:"_id,omitempty"`
	Version        string      `json:"version"          bson:"version"`
	ConsumerId     string      `json:"consumer_id"      bson:"consumer_id"`
	Origin         string      `json:"origin"           bson:"origin"`
	SessionId      string      `json:"session_id"       bson:"session_id"`
	Service        string      `json:"service"          bson:"service"`
	EntryDate      time.Time   `json:"entry_date"       bson:"entry_date"`
	EntryTimestamp float64     `json:"entry_timestamp"  bson:"entry_timestamp,minsize"`
	ElapsedSeconds float64     `json:"elapsed_seconds"  bson:"elapsed_seconds,minsize"`
	Request        Request     `json:"request"          bson:"request"`
	Response       Response    `json:"response"         bson:"response"`
	ReadAccess     []string    `json:"read_access"      bson:"read_access"`
	Custom         interface{} `json:"custom"           bson:"custom"`
}
