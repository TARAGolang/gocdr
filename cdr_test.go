package gocdr

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/fulldump/apitest"
	"github.com/fulldump/golax"
)

func Test_API_HappyPath(t *testing.T) {

	a := golax.NewApi()

	a.Root.
		Interceptor(InterceptorCdr2Header()).
		Interceptor(InterceptorCdr("invented-service")).
		Node("{param1}").
		Node("{param2}").
		Node("test-node").
		Method("GET", func(c *golax.Context) {

		// Don't write nothing to request !!

		cdr := GetCdr(c)
		cdr.Custom = map[string]interface{}{
			"a": 20,
			"b": 55,
		}

	})

	s := apitest.New(a)

	method := "GET"
	url := "/value-1/value-2/test-node?query_a=aaa&query_b=bbb"
	consumer_id := "my-consumer-id"

	r := s.Request(method, url).
		WithHeader("X-Consumer-Id", consumer_id).
		Do()

	json_cdr := r.Header.Get("X-Cdr")
	fmt.Println(json_cdr)

	cdr := &Definition{}
	json.Unmarshal([]byte(json_cdr), cdr)

	if consumer_id != cdr.ConsumerId {
		t.Error("consumer_id")
	}

	if "invented-service" != cdr.Service {
		t.Error("service")
	}

	if "GET" != cdr.Request.Method {
		t.Error("request.method")
	}

	if url != cdr.Request.URI {
		t.Error("request.uri")
	}

	if "/{param1}/{param2}/test-node" != cdr.Request.Handler {
		t.Error("request.handler")
	}

	if 0 != cdr.Request.Length {
		t.Error("request.length")
	}

	args := map[string][]string{
		"query_a": []string{"aaa"},
		"query_b": []string{"bbb"},
	}

	if !reflect.DeepEqual(args, cdr.Request.Args) {
		t.Error("request.args")
	}

	if 0 != cdr.Response.Length {
		t.Error("response.length")
	}

	if 200 != cdr.Response.StatusCode {
		t.Error("response.status_code")
	}

}
