package gocdr

import (
	"encoding/json"

	"github.com/fulldump/golax"
)

func InterceptorCdr2Header() *golax.Interceptor {
	return &golax.Interceptor{
		Documentation: golax.Doc{
			Name: "InterceptorCdr2Log",
			Description: `
			Return Cdr in a header ´X-Cdr´.

			This interceptor should wrap ´InterceptorCdr´
		`,
		},
		After: func(c *golax.Context) {

			cdr := GetCdr(c)

			serialized, err := json.Marshal(cdr)
			if nil != err {
				panic(err)
			}

			c.Response.Header().Add("X-Cdr", string(serialized))

		},
	}
}
