package model

type Error struct {
	Code        int    `json:"code"        bson:"code"`
	Description string `json:"description" bson:"description"`
}
