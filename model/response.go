package model

type Response struct {
	StatusCode int    `json:"status_code" bson:"status_code"`
	Length     int64  `json:"length"      bson:"length"`
	Error      *Error `json:"error"       bson:"error"`
}
