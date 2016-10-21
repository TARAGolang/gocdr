package model

type Request struct {
	Method  string              `json:"method"   bson:"method"`
	URI     string              `json:"uri"      bson:"uri"`
	Handler string              `json:"handler"  bson:"handler"`
	Args    map[string][]string `json:"args"     bson:"args"`
	Length  int64               `json:"length"   bson:"length"`
}
