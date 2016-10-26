package model

type Request struct {
	Method     string              `json:"method"     bson:"method"`
	URI        string              `json:"uri"        bson:"uri"`
	Handler    string              `json:"handler"    bson:"handler"`
	Query      map[string][]string `json:"query"      bson:"query"`
	Parameters map[string]string   `json:"parameters" bson:"parameters"`
	Length     int64               `json:"length"     bson:"length"`
}
