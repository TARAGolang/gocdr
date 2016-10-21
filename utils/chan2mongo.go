package utils

import (
	"time"

	"vendor/gopkg.in/mgo.v2"
)

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
