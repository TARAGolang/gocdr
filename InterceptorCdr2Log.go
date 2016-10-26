package gocdr

import (
	"encoding/json"
	"log"

	"github.com/fulldump/golax"
)

// InterceptorCdr2Log log CDR to `stdout`. This interceptor should wrap `InterceptorCdr`.
//
// Typical usage:
//
// 	a := golax.NewApi()
//
// 	a.Root.
// 	    Interceptor(InterceptorCdr2Log()).
// 	    Interceptor(InterceptorCdr("invented-service")).
// 	    Method("GET", func(c *golax.Context) {
// 	        // Implement your API here
// 	    })
//
// Here is the sample output:
//
// 	2016/10/24 19:49:19 CDR {"id":"580a465bce507629a613107c","version":"1.0.0","consumer_id":"my-consumer-id","origin":"127.0.0.1","session_id":"","service":"invented-service","entry_date":"2016-10-21T18:46:19.820423299+02:00","entry_timestamp":1.4770683798204234e+09,"elapsed_seconds":1.621246337890625e-05,"request":{"method":"POST","uri":"/value-1/value-2/test-node?query_a=aaa\u0026query_b=bbb","handler":"/{param1}/{param2}/test-node","args":{"query_a":["aaa"],"query_b":["bbb"]},"length":2},"response":{"status_code":222,"length":5,"error":{"code":27,"description":"my-error-description"}},"read_access":["other-involved-consumer-id","my-consumer-id"],"custom":{"a":20,"b":55}}

func InterceptorCdr2Log() *golax.Interceptor {
	return &golax.Interceptor{
		Documentation: golax.Doc{
			Name: "Cdr2Log",
			Description: `
			Store cdrs into log.
		`,
		},
		After: func(c *golax.Context) {

			cdr := GetCdr(c)

			serialized, err := json.Marshal(cdr)

			if nil != err {
				// TODO: Log this somewere
			}

			log.Printf(
				"%s\t%s",
				"CDR",
				serialized,
			)
		},
	}
}
