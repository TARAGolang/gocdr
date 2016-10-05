package gocdr

import "github.com/fulldump/golax"

func InterceptorCdr2Mongo() *golax.Interceptor {
	return &golax.Interceptor{
		Documentation: golax.Doc{
			Name: "Cdr2Mongo",
			Description: `
			Store cdrs into a mongo db.
		`,
		},
		After: func(c *golax.Context) {
			cdr := GetCdr(c)
			_ = cdr
			// TODO: Write to mongo collection in a goroutine / channel :)
		},
	}
}
