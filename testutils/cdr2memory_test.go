package testutils

import (
	"reflect"
	"testing"

	"github.com/fulldump/apitest"
	"github.com/fulldump/golax"

	"github.com/smartdigits/gocdr/model"
	"github.com/smartdigits/gocdr/testutils"
)

func Test_Cdr2memory(t *testing.T) {

	cdrtest := testutils.NewTestCDR()

	a := golax.NewApi()

	a.Root.
		Interceptor(cdrtest.InterceptorCdr2Memory()).
		Interceptor(InterceptorCdr2Log()).
		Interceptor(InterceptorCdr("invented-service")).
		Node("api").
		Method("GET", func(c *golax.Context) {

		name := c.Request.URL.Query().Get("name")

		cdr := GetCdr(c)
		cdr.Custom = map[string]interface{}{
			"name": name,
		}

	})

	s := apitest.New(a)

	method := "GET"
	url := "/api?name="
	consumer_id := "my-consumer-id"

	s.Request(method, url+"one").
		WithHeader("X-Consumer-Id", consumer_id).
		Do()

	s.Request(method, url+"two").
		WithHeader("X-Consumer-Id", consumer_id).
		Do()

	if cdrtest.Memory[0].Custom.(map[string]interface{})["name"] != "one" {
		t.Error("CDR 'one' not found in memory (first position)")
	}

	if cdrtest.Memory[1].Custom.(map[string]interface{})["name"] != "two" {
		t.Error("CDR 'two' not found in memory (second position)")
	}

	// Check Reset()
	cdrtest.Reset()
	if !reflect.DeepEqual(cdrtest.Memory, []*model.CDR{}) {
		t.Error("`TestCDR.Reset()` should empty the `Memory` array.")
		return
	}

}
