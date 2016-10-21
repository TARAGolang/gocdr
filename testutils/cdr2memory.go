package testutils

import (
	"github.com/fulldump/golax"

	"github.com/smartdigits/gocdr/model"
)

type TestCDR struct {
	Memory []*model.CDR
}

func (t *TestCDR) Reset() {
	t.Memory = []*model.CDR{}
}

func NewTestCDR() *TestCDR {
	t := &TestCDR{}
	t.Reset()

	return t
}

func (t *TestCDR) InterceptorCdr2Memory() *golax.Interceptor {
	return &golax.Interceptor{
		Documentation: golax.Doc{
			Name: "Cdr2Memory",
			Description: `
			Save cdrs into memory.
		`,
		},
		After: func(c *golax.Context) {
			cdr := GetCdr(c)
			t.Memory = append(t.Memory, cdr)
		},
	}
}
