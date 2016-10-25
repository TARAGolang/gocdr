package gocdr

import (
	"github.com/fulldump/golax"

	"github.com/smartdigits/gocdr/model"
)

func InterceptorCdr2Channel(s chan *model.CDR) *golax.Interceptor {
	return &golax.Interceptor{
		Documentation: golax.Doc{
			Name: "Cdr2Channel",
			Description: `
			Push CDRs to a channel.
		`,
		},
		After: func(c *golax.Context) {
			cdr := GetCdr(c)
			s <- cdr
		},
	}
}
