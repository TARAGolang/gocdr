package model

// Service identifies the process and other information related to the service.
type Service struct {

	// Name is the service name.
	Name string `json:"name" bson:"name"`

	// Version is the service version, for example: 1.2, 0.0.1, 7.3.
	Version string `json:"version" bson:"version"`

	// Commit is the short commit number to identify exactly the code base
	// that is being executed
	Commit string `json:"commit" bson:"commit"`
}
