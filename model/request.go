package model

// Request represents a CDR Request.
type Request struct {

	// Method is the HTTP verb
	Method string `json:"method" bson:"method"`

	// URI is the raw URL (the same one sent by the client)
	URI string `json:"uri" bson:"uri"`

	// Handler is the URL before replacing parameters
	Handler string `json:"handler" bson:"handler"`

	// Parameters is the URL parameters sent by client
	Parameters map[string]string `json:"parameters" bson:"parameters"`

	// Query is the URL query params
	Query map[string][]string `json:"query" bson:"query"`

	// Length is the client request body length
	Length int64 `json:"length" bson:"length"`
}
