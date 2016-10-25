package gocdr

import (
	"net/http"
	"strings"
	"time"

	"github.com/fulldump/golax"
	"gopkg.in/mgo.v2/bson"

	"github.com/smartdigits/gocdr/constants"
	"github.com/smartdigits/gocdr/model"
)

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
				"id": "580a465bce507629a613107c",
				"version": "1.0.0",
				"consumer_id": "my-consumer-id",
				"origin": "127.0.0.1",
				"session_id": "",
				"service": "invented-service",
				"entry_date": "2016-10-21T18:46:19.820423299+02:00",
				"entry_timestamp": 1.4770683798204234e+09,
				"elapsed_seconds": 1.621246337890625e-05,
				"request": {
					"method": "POST",
					"uri": "/value-1/value-2/test-node?query_a=aaa\u0026query_b=bbb",
					"handler": "/{param1}/{param2}/test-node",
					"args": {
						"query_a": ["aaa"],
						"query_b": ["bbb"]
					},
					"length": 2
				},
				"response": {
					"status_code": 222,
					"length": 5,
					"error": {
						"code": 27,
						"description": "my-error-description"
					}
				},
				"read_access": ["other-involved-consumer-id", "my-consumer-id"],
				"custom": {
					"a": 20,
					"b": 55
				}
			}
			´´´
		`,
		},
		Before: func(c *golax.Context) {

			cdr := &model.CDR{
				Id:             bson.NewObjectId(),
				Version:        constants.VERSION,
				Service:        service,
				Origin:         formatRemoteAddr(c.Request),
				EntryDate:      time.Now(),
				EntryTimestamp: float64(time.Now().UnixNano()) / 1000000000,
				ElapsedSeconds: 0,
				Request: model.Request{
					Length: c.Request.ContentLength,
					Method: c.Request.Method,
					URI:    c.Request.RequestURI,
					Args:   c.Request.URL.Query(),
				},
				Custom: map[string]interface{}{},
			}

			c.Set(constants.CONTEXT_KEY, cdr)

		},
		After: func(c *golax.Context) {

			cdr := GetCdr(c)

			exit_timestamp := float64(time.Now().UnixNano()) / 1000000000
			cdr.ElapsedSeconds = exit_timestamp - cdr.EntryTimestamp

			consumer_id := c.Request.Header.Get("X-Consumer-Id")

			cdr.ConsumerId = consumer_id
			cdr.AddReadAccess(consumer_id)

			cdr.Request.Handler = c.PathHandlers

			cdr.Response.StatusCode = c.Response.StatusCode
			cdr.Response.Length = int64(c.Response.Length)

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