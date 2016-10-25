package gocdr

import (
	"encoding/json"
	"log"

	"github.com/fulldump/golax"
)

func InterceptorCdr2Log() *golax.Interceptor {
	return &golax.Interceptor{
		Documentation: golax.Doc{
			Name: "Cdr2Log",
			Description: `
			Store cdrs into log.
		`,
		},
		After: func(c *golax.Context) {

			cdr := GetCdr(c)

			serialized, err := json.Marshal(cdr)

			if nil != err {
				// TODO: Log this somewere
			}

			log.Printf(
				"%s\t%s",
				"CDR",
				serialized,
			)
		},
	}
}
