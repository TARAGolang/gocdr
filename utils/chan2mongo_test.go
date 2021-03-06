package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/fulldump/apitest"
	"github.com/fulldump/golax"
	uuid "github.com/satori/go.uuid"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/smartdigits/gocdr"
	"github.com/smartdigits/gocdr/model"
	"github.com/smartdigits/gocdr/testutils"
)

func Test_Chan2Mongo(t *testing.T) {

	dbName := "gocdr-" + uuid.NewV4().String()

	session, _ := mgo.Dial("localhost")
	session.SetMode(mgo.Monotonic, true)
	session.SetSyncTimeout(100 * time.Millisecond) // Insert fast fail

	db := session.DB(dbName)

	defer func(d *mgo.Database) {
		//d.DropDatabase()
	}(db)

	cdrs := make(chan *model.CDR, 10) // #1 make channel

	Chan2Mongo(cdrs, db) // #2 dump channel to mongo

	cdrtest := testutils.NewTestCDR()

	a := golax.NewApi()
	a.Root.
		Interceptor(cdrtest.InterceptorCdr2Memory()).
		Interceptor(gocdr.InterceptorCdr2Channel(cdrs)). // #3 put CDRs in channel
		Interceptor(gocdr.InterceptorCdr(nil)).
		Node("api").
		Method("GET", func(c *golax.Context) {})

	s := apitest.New(a)

	// Do sample SYNCed requests
	for i := 0; i < 10; i++ {
		s.Request("GET", "/api?name="+strconv.Itoa(i)).
			WithHeader("X-Consumer-Id", "my-consumer-id").
			Do()
	}

	// wait all cdrs to be stored in mongo
	for len(cdrs) != 0 {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("chan.length:", len(cdrs))
	}

	// Extract from mongo and compare to memory
	for i, memoryCdr := range cdrtest.Memory {

		collection := db.C(GetCollectionName(memoryCdr))

		mongoCdr := &model.CDR{}
		err := collection.Find(bson.M{"request.query.name": strconv.Itoa(i)}).One(mongoCdr)
		if nil != err {
			panic(err)
		}

		// Tweak differences between mongo representation and memory:
		mongoCdr.Id = memoryCdr.Id
		mongoCdr.Custom = map[string]interface{}{}
		mongoCdr.EntryDate = memoryCdr.EntryDate

		if !reflect.DeepEqual(mongoCdr, memoryCdr) {
			t.Error("Channel CDR does not match with memory CDR")
		}
	}

}
