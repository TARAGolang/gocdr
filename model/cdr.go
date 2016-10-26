package model

import "time"

// CDR represents a "Call Data Record" information. The purpose of this kind
// of records is access audits and billing.
type CDR struct {

	// Id is a random string to identify a specific CDR.
	Id string `json:"id" bson:"_id,omitempty"`

	// Version is the specification version used in this CDR.
	Version string `json:"version" bson:"version"`

	// ConsumerId is the header-injected `consumer_id`.
	ConsumerId string `json:"consumer_id" bson:"consumer_id"`

	// Origin is the real client IP (v4 and v6 are supported). If the service is
	// behind a proxy, the real IP should be forwarded in the header
	// `X-Forwarded-For`.
	Origin string `json:"origin" bson:"origin"`

	// SessionId stores cookie session, in case the client was using a session.
	SessionId string `json:"session_id" bson:"session_id"`

	// Service indicates the service/server name being provided.
	Service string `json:"service" bson:"service"`

	// EntryDate stores a `Time` object with the starting request timestamp.
	EntryDate time.Time `json:"entry_date" bson:"entry_date"`

	// EntryTimestamp stores `EntryDate` in UNIX time format in seconds.
	EntryTimestamp float64 `json:"entry_timestamp" bson:"entry_timestamp,minsize"`

	// ElapsedSeconds is the real time the request has taken in seconds.
	ElapsedSeconds float64 `json:"elapsed_seconds" bson:"elapsed_seconds,minsize"`

	// Request is a struct with standard HTTP client request information
	// (automatically filled up).
	Request Request `json:"request" bson:"request"`

	// Response is a struct with standard HTTP response information.
	Response Response `json:"response" bson:"response"`

	// Array with all consumerIds allowed to read this CDR.
	ReadAccess []string `json:"read_access" bson:"read_access"`

	// Custom stores specific service information that is service-dependent in
	// any format.
	Custom interface{} `json:"custom" bson:"custom"`
}
