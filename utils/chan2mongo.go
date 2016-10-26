package utils

import (
	"time"

	"gopkg.in/mgo.v2"

	"github.com/smartdigits/gocdr/model"
)

// Chan2Mongo is a util to extract CDRs from a channel and store to a MongoDB
// collection. This interceptor should wrap `InterceptorCdr`.
//
// Note that this interceptor do not extract the CDRs from the channel, that should
// be done by other task. There is a util included (`Chan2Mongo`) to do that.
//
// Typical usage:
// 	channel_cdrs := make(chan *model.CDR, 100)  // Buffered channel, 100 items
//
// 	collection_cdrs := mongo_db.C("cdrs") // assume `mongo_db` already exists
//
// 	Chan2Mongo(channel_cdrs, collection_cdrs) // do the job: channel -> mongo
//
// 	a := golax.NewApi()
//
// 	a.Root.
// 	    Interceptor(InterceptorCdr2Channel(channel_cdrs)). // Pass created channel
// 	    Interceptor(InterceptorCdr("invented-service")).
// 	    Method("GET", func(c *golax.Context) {
// 	        // Implement your API here
// 	    })
//
func Chan2Mongo(s chan *model.CDR, m *mgo.Collection) {

	go func() {

		for {
			cdr := <-s

			if nil == cdr {
				break
			}

			err := m.Insert(cdr)
			if nil != err {
				// TODO: alarm warning?
				s <- cdr
				time.Sleep(1 * time.Second)
			}
		}

	}()

}
