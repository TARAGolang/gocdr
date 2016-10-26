package testutils

import (
	"github.com/fulldump/golax"

	"github.com/smartdigits/gocdr/model"
)

// TestCDR is a util for testing CDRs. It can store CDRs to memory and read
// them at testing time.
type TestCDR struct {

	// Memory is the secuence of CDRs in a test
	Memory []*model.CDR
}

// Reset clean all CDRs from memory. SetUp test is a good place to use it.
func (t *TestCDR) Reset() {
	t.Memory = []*model.CDR{}
}

// NewTestCDR Create a new TestCDR
//
// Typical usage:
//
// 	func Test_ExampleBilling(t *testing.T) {
//
// 	    cdrtest := testutils.NewTestCDR() // IMPORTANT
//
// 	    a := golax.NewApi()
// 	    a.Root.Interceptor(cdrtest.InterceptorCdr2Memory()) // IMPORTANT
//
// 	    BuildYourApi(a)
//
// 	    s := apitest.New(a)
// 	    s.Request("GET", "/my-url-to-test").Do()
//
// 	    // IMPORTANT: Do things with cdrtest.Memory[i], for example:
// 	    if 200 != cdrtest.Memory[0].Response.StatusCode {
// 	        t.Error("blah blah blah...")
// 	    }
// 	}
func NewTestCDR() *TestCDR {
	t := &TestCDR{}
	t.Reset()

	return t
}

// InterceptorCdr2Memory return the interceptor you should put in your api to
// capture CDRs
func (t *TestCDR) InterceptorCdr2Memory() *golax.Interceptor {
	return &golax.Interceptor{
		Documentation: golax.Doc{
			Name: "Cdr2Memory",
			Description: `
			Save cdrs into memory.
		`,
		},
		After: func(c *golax.Context) {
			cdr := getCdr(c)
			t.Memory = append(t.Memory, cdr)
		},
	}
}
