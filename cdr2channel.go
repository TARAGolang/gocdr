package gocdr

import "github.com/fulldump/golax"

func InterceptorCdr2Channel(s chan *Definition) *golax.Interceptor {
	return &golax.Interceptor{
		Documentation: golax.Doc{
			Name: "Cdr2Channel",
			Description: `
			Push CDRs to a channel.
		`,
		},
		After: func(c *golax.Context) {
			cdr := GetCdr(c)
			_ = cdr
			// TODO: Write to channel :)
		},
	}
}
