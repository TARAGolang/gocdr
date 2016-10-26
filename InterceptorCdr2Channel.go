package gocdr

import (
	"github.com/fulldump/golax"

	"github.com/smartdigits/gocdr/model"
)

// InterceptorCdr2Channel push CDRs to a channel. This interceptor should wrap
// `InterceptorCdr`.
//
// Note that this interceptor do not extract the CDRs from the channel, that should
// be done by other task. There is a util included (`Chan2Mongo`) to do that.
//
// Typical usage:
//
// 	channel_cdrs := make(chan *model.CDR, 100)  // Buffered channel, 100 items
//
// 	// assume `mongo_db` already exists
//
// 	Chan2Mongo(channel_cdrs, mongo_db) // do the job: channel -> mongo
//
// 	a := golax.NewApi()
//
// 	a.Root.
// 	    Interceptor(InterceptorCdr2Channel(channel_cdrs)). // Pass created channel
// 	    Interceptor(InterceptorCdr("invented-service")).
// 	    Method("GET", func(c *golax.Context) {
// 	        // Implement your API here
// 	    })
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
