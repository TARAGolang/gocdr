package model

// Error represents a response error object
type Error struct {

	// Code is a specific application error code (integer). It could be any
	// integer number: `200`, `396`, `2`, `3213254`, ...
	Code int `json:"code" bson:"code"`

	// Description is a human readable description for the error. It is a
	// specific application domain description, to developers.
	Description string `json:"description" bson:"description"`
}
