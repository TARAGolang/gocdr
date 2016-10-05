package gocdr

import "gopkg.in/mgo.v2/bson"

type Definition struct {
	Id             bson.ObjectId `json:"id"               bson:"_id"`
	ConsumerId     string        `json:"consumer_id"      bson:"consumer_id"`
	Origin         string        `json:"origin"           bson:"origin"`
	SessionId      string        `json:"session_id"       bson:"session_id"`
	Service        string        `json:"service"          bson:"service"`
	EntryTimestamp float64       `json:"entry_timestamp"  bson:"entry_timestamp,minsize"`
	ExitTimestamp  float64       `json:"exit_timestamp"   bson:"exit_timestamp,minsize"`
	ElapsedTime    float64       `json:"elapsed_time"     bson:"elapsed_time,minsize"`
	Request        Request       `json:"request"          bson:"request"`
	Response       Response      `json:"response"         bson:"response"`
	ReadAccess     []string      `json:"read_access"      bson:"read_access"`
	Custom         interface{}   `json:"custom"           bson:"custom"`
}

type Request struct {
	Method  string              `json:"method"   bson:"method"`
	URI     string              `json:"uri"      bson:"uri"`
	Handler string              `json:"handler"  bson:"handler"`
	Args    map[string][]string `json:"args"     bson:"args"`
	Length  int64               `json:"length"   bson:"length"`
}

type Response struct {
	StatusCode int    `json:"status_code" bson:"status_code"`
	Length     int    `json:"length"      bson:"length"`
	Error      *Error `json:"error"       bson:"error"`
}

type Error struct {
	Code        int    `json:"code" bson:"code"`
	Description string `json:"description" bson:"description"`
}

func (cdr *Definition) SetError(code int, desc string) {
	cdr.Response.Error = &Error{
		Code:        code,
		Description: desc,
	}
}

func (cdr *Definition) AddReadAccess(consumer_id string) bool {
	if "" == consumer_id {
		return false
	}

	for _, element := range cdr.ReadAccess {
		if consumer_id == element {
			return false
		}
	}
	cdr.ReadAccess = append(cdr.ReadAccess, consumer_id)
	return true
}
