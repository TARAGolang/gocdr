package model

// Response represents a CDR response.
type Response struct {

	// StatusCode is the returned status code number.
	StatusCode int `json:"status_code" bson:"status_code"`

	// Length is the response body length.
	Length int64 `json:"length" bson:"length"`

	// Error represents an error object.
	Error *Error `json:"error" bson:"error"`
}
