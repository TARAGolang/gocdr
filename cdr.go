package gocdr

import (
	"net/http"
	"strings"
	"time"

	"github.com/fulldump/golax"
	"gopkg.in/mgo.v2/bson"
)

const _CONTEXT_KEY = "cdr-a0f1360d58b5e766ec1adb9bcd42a954"

func InterceptorCdr(service string) *golax.Interceptor {
	return &golax.Interceptor{
		Documentation: golax.Doc{
			Name: "InterceptorCDR",
			Description: `
			Collect standard request information and custom service information
			and makeis available for processing it.

			Sample CDR:
			´´´json

			{
				"id": "57e26675ce50761c7950116d",
				"consumer_id": "a",
				"origin": "127.0.0.1",
				"session_id": "",
				"service": "dormer",
				"entry_timestamp": 1474455157.9033415,
				"exit_timestamp": 1474455157.904661,
				"elapsed_time": 0.0013194084167480469,
				"request": {
					"method": "GET",
					"uri": "/dormer/v1/exec/meteosat/weather/hello",
					"handler": "/dormer/v1/exec/{organization_name}/{dataset_name}/{operation_name}",
					"args": {
						"organization_name": "meteosat",
						"dataset_name": "weather",
						"operation_name": "hello"
					},
					"length": 0
				},
				"response": {
					"status_code": 200,
					"length": 153,
					"error": null
				},
				"read_access": ["a"],
				"custom": {}
			}

			´´´
		`,
		},
		Before: func(c *golax.Context) {

			cdr := &Definition{
				Id:             bson.NewObjectId(),
				Service:        service,
				Origin:         formatRemoteAddr(c.Request),
				EntryTimestamp: float64(time.Now().UnixNano()) / 1000000000,
				ElapsedTime:    0,
				Request: Request{
					Method: c.Request.Method,
					URI:    c.Request.RequestURI,
					Args:   c.Request.URL.Query(),
				},
				Custom: map[string]interface{}{},
			}

			c.Set(_CONTEXT_KEY, cdr)

		},
		After: func(c *golax.Context) {

			cdr := GetCdr(c)

			cdr.ExitTimestamp = float64(time.Now().UnixNano()) / 1000000000
			cdr.ElapsedTime = cdr.ExitTimestamp - cdr.EntryTimestamp

			consumer_id := c.Request.Header.Get("X-Consumer-Id")

			cdr.ConsumerId = consumer_id
			cdr.AddReadAccess(consumer_id)

			cdr.Request.Length = int64(c.Response.Length)
			cdr.Request.Handler = c.PathHandlers

			cdr.Response.StatusCode = c.Response.StatusCode
			cdr.Response.Length = c.Response.Length

			err := c.LastError
			if nil != err {
				cdr.SetError(err.ErrorCode, err.Description)
			}

		},
	}
}

func formatRemoteAddr(r *http.Request) string {
	xorigin := strings.TrimSpace(strings.Split(
		r.Header.Get("X-Forwarded-For"), ",")[0])
	if xorigin != "" {
		return xorigin
	} else {
		return r.RemoteAddr[0:strings.LastIndex(r.RemoteAddr, ":")]
	}
}

func GetCdr(c *golax.Context) *Definition {
	v, exists := c.Get(_CONTEXT_KEY)

	if !exists {
		return nil
	}
	return v.(*Definition)
}
