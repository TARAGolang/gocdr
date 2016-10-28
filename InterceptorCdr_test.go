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
	consumerID := "my-consumer-id"
	requestBody := "hi"
	responseBody := "hello"

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

		fmt.Fprintf(c.Response, responseBody)
	})

	s := apitest.New(a)

	s.Request(method, url).
		WithHeader("X-Consumer-Id", consumerID).
		WithBodyString(requestBody).
		Do()

	cdr := cdrtest.Memory[0]

	if consumerID != cdr.ConsumerId {
		t.Error("consumerID")
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

	if int64(len(requestBody)) != cdr.Request.Length {
		t.Error("request.length")
	}

	query := map[string][]string{
		"query_a": []string{"aaa"},
		"query_b": []string{"bbb"},
	}

	if !reflect.DeepEqual(query, cdr.Request.Query) {
		t.Error("request.query")
	}

	parameters := map[string]string{
		"param1": "value-1",
		"param2": "value-2",
	}

	if !reflect.DeepEqual(parameters, cdr.Request.Parameters) {
		t.Error("request.parameters")
	}

	if int64(len(responseBody)) != cdr.Response.Length {
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

	access := []string{"other-involved-consumer-id", consumerID}
	if !reflect.DeepEqual(access, cdr.ReadAccess) {
		t.Error("read_access")
	}

}
