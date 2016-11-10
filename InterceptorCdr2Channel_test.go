package gocdr

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/fulldump/apitest"
	"github.com/fulldump/golax"

	"github.com/smartdigits/gocdr/model"
	"github.com/smartdigits/gocdr/testutils"
)

func Test_Cdr2channel(t *testing.T) {

	cdrs := make(chan *model.CDR, 10)

	cdrtest := testutils.NewTestCDR()

	a := golax.NewApi()

	a.Root.
		Interceptor(cdrtest.InterceptorCdr2Memory()).
		Interceptor(InterceptorCdr2Channel(cdrs)).
		Interceptor(InterceptorCdr(&model.Service{Name: "invented-service"})).
		Node("api").
		Method("GET", func(c *golax.Context) {

		GetCdr(c).Custom = map[string]interface{}{
			"name": c.Request.URL.Query().Get("name"),
		}

	})

	s := apitest.New(a)

	// Do sample SYNCed requests
	for i := 0; i < 10; i++ {
		s.Request("GET", "/api?name="+strconv.Itoa(i)).
			WithHeader("X-Consumer-Id", "my-consumer-id").
			Do()
	}

	// Extract from channel and compare to memory
	for _, memoryCdr := range cdrtest.Memory {
		channelCdr := <-cdrs
		if !reflect.DeepEqual(channelCdr, memoryCdr) {
			t.Error("Channel CDR does not match with memory CDR")
		}
	}

}
