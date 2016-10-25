package gocdr

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fulldump/apitest"
	"github.com/fulldump/golax"

	"github.com/smartdigits/gocdr/testutils"
)

func Test_Cdr(t *testing.T) {

	// Test configuration
	method := "POST"
	url := "/value-1/value-2/test-node?query_a=aaa&query_b=bbb"
	consumer_id := "my-consumer-id"
	request_body := "hi"
	response_body := "hello"

	cdrtest := testutils.NewTestCDR()

	a := golax.NewApi()

	a.Root.
		Interceptor(cdrtest.InterceptorCdr2Memory()).
		Interceptor(InterceptorCdr2Log()).
		Interceptor(InterceptorCdr("invented-service")).
		Node("{param1}").
		Node("{param2}").
		Node("test-node").
		Method("POST", func(c *golax.Context) {

		cdr := GetCdr(c)
		cdr.Custom = map[string]interface{}{
			"a": 20,
			"b": 55,
		}
		cdr.AddReadAccess("other-involved-consumer-id")

		c.Error(222, "my-error-description").ErrorCode = 27

		fmt.Fprintf(c.Response, response_body)
	})

	s := apitest.New(a)

	s.Request(method, url).
		WithHeader("X-Consumer-Id", consumer_id).
		WithBodyString(request_body).
		Do()

	cdr := cdrtest.Memory[0]

	if consumer_id != cdr.ConsumerId {
		t.Error("consumer_id")
	}

	if "invented-service" != cdr.Service {
		t.Error("service")
	}

	if method != cdr.Request.Method {
		t.Error("request.method")
	}

	if url != cdr.Request.URI {
		t.Error("request.uri")
	}

	if "/{param1}/{param2}/test-node" != cdr.Request.Handler {
		t.Error("request.handler")
	}

	if int64(len(request_body)) != cdr.Request.Length {
		t.Error("request.length")
	}

	args := map[string][]string{
		"query_a": []string{"aaa"},
		"query_b": []string{"bbb"},
	}

	if !reflect.DeepEqual(args, cdr.Request.Args) {
		t.Error("request.args")
	}

	if int64(len(response_body)) != cdr.Response.Length {
		t.Error("response.length")
	}

	if 222 != cdr.Response.StatusCode {
		t.Error("response.status_code")
	}

	if 27 != cdr.Response.Error.Code {
		t.Error("response.error.code")
	}

	if "my-error-description" != cdr.Response.Error.Description {
		t.Error("response.error.description")
	}

	access := []string{"other-involved-consumer-id", consumer_id}
	if !reflect.DeepEqual(access, cdr.ReadAccess) {
		t.Error("read_access")
	}

}
