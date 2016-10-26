# gocdr

[![Build Status](https://travis-ci.org/smartdigits/gocdr.svg?branch=master)](https://travis-ci.org/smartdigits/gocdr)
[![Go report card](http://goreportcard.com/badge/smartdigits/gocdr)](https://goreportcard.com/report/smartdigits/gocdr)
[![GoDoc](https://godoc.org/github.com/smartdigits/gocdr?status.svg)](https://godoc.org/github.com/smartdigits/gocdr)

<sup>Tested for Go 1.5, 1.6, 1.7, tip</sup>

CDRs SDK for SmartDigits billing system

<!-- MarkdownTOC autolink=true bracket=round depth=4 -->

- [InterceptorCdr](#interceptorcdr)
	- [Custom info](#custom-info)
	- [Access](#access)
	- [Sample CDR](#sample-cdr)
- [InterceptorCdr2Log](#interceptorcdr2log)
- [InterceptorCdr2Channel](#interceptorcdr2channel)
- [GetCdr](#getcdr)
- [NewTestCDR](#newtestcdr)
- [Dependencies](#dependencies)
- [Testing](#testing)

<!-- /MarkdownTOC -->


Gocdr is a set of interceptors to generate and store CDRs:

* [InterceptorCdr](#interceptorcdr) - Generate and fill a CDR
* [InterceptorCdr2Log](#interceptorcdr2log) - Log CDRs
* [InterceptorCdr2Channel](#interceptorcdr2channel) - Collect all CDRs in a channel

And some helpers:

* [GetCdr](#getcdr) - Retrieve CDR from context
* [NewTestCDR](#newtestcdr) - Store CDRs in memory to allow unit testing


## InterceptorCdr

Attach a new CDR to the context and populate it automatically.

Typical usage:

```go
a := golax.NewApi()

a.Root.
	Interceptor(InterceptorCdr2Log()).
	Interceptor(InterceptorCdr("invented-service")).
	Method("GET", func(c *golax.Context) {
		// Implement your API here
	})

```

### Custom info

CDR allow to store custom service information:

```go
cdr := GetCdr(c)
cdr.Custom = map[string]interface{}{
	"a": 20,
	"b": 55,
}
```

### Access

Typically only the consumer using the service can read the CDR but sometimes, a
third party is involved (for example, somebody sharing a bucket with me).

In that case, the service should add the third consumer_id to the access list:

```go
cdr := GetCdr(c)
cdr.AddReadAccess("other-involved-consumer-id")
```

### Sample CDR


A sample CDR in JSON:
```json
{
	"id": "580a465bce507629a613107c",
	"version": "1.0.0",
	"consumer_id": "my-consumer-id",
	"origin": "127.0.0.1",
	"session_id": "",
	"service": "invented-service",
	"entry_date": "2016-10-21T18:46:19.820423299+02:00",
	"entry_timestamp": 1.4770683798204234e+09,
	"elapsed_seconds": 1.621246337890625e-05,
	"request": {
		"method": "POST",
		"uri": "/value-1/value-2/test-node?query_a=aaa\u0026query_b=bbb",
		"handler": "/{param1}/{param2}/test-node",
		"args": {
			"query_a": ["aaa"],
			"query_b": ["bbb"]
		},
		"length": 2
	},
	"response": {
		"status_code": 222,
		"length": 5,
		"error": {
			"code": 27,
			"description": "my-error-description"
		}
	},
	"read_access": ["other-involved-consumer-id", "my-consumer-id"],
	"custom": {
		"a": 20,
		"b": 55
	}
}
```


```
TODO: Explain all fields in a CDR
```

## InterceptorCdr2Log

Log CDR to `stdout`. This interceptor should wrap `InterceptorCdr`.

Typical usage:

```go
a := golax.NewApi()

a.Root.
	Interceptor(InterceptorCdr2Log()).
	Interceptor(InterceptorCdr("invented-service")).
	Method("GET", func(c *golax.Context) {
		// Implement your API here
	})

```

Here is the sample output:

```txt
2016/10/24 19:49:19 CDR {"id":"580a465bce507629a613107c","version":"1.0.0","consumer_id":"my-consumer-id","origin":"127.0.0.1","session_id":"","service":"invented-service","entry_date":"2016-10-21T18:46:19.820423299+02:00","entry_timestamp":1.4770683798204234e+09,"elapsed_seconds":1.621246337890625e-05,"request":{"method":"POST","uri":"/value-1/value-2/test-node?query_a=aaa\u0026query_b=bbb","handler":"/{param1}/{param2}/test-node","args":{"query_a":["aaa"],"query_b":["bbb"]},"length":2},"response":{"status_code":222,"length":5,"error":{"code":27,"description":"my-error-description"}},"read_access":["other-involved-consumer-id","my-consumer-id"],"custom":{"a":20,"b":55}}
```


## InterceptorCdr2Channel

Push CDRs to a channel. This interceptor should wrap `InterceptorCdr`.

Note that this interceptor do not extract the CDRs from the channel, that should
be done by other task. There is a util included (`Chan2Mongo`) to do that.

Typical usage:
```go
channel_cdrs := make(chan *model.CDR, 100)  // Buffered channel, 100 items

collection_cdrs := mongo_db.C("cdrs") // assume `mongo_db` already exists

Chan2Mongo(channel_cdrs, collection_cdrs) // do the job: channel -> mongo

a := golax.NewApi()

a.Root.
	Interceptor(InterceptorCdr2Channel(channel_cdrs)). // Pass created channel
	Interceptor(InterceptorCdr("invented-service")).
	Method("GET", func(c *golax.Context) {
		// Implement your API here
	})
```

## GetCdr

Get CDR object from context.

Typical usage:

```go
func MyHandler(c *golax.Context) {
	// ...
	cdr := gocdr.GetCdr(c)
	// ...
}
```

## NewTestCDR

Store CDRs in memory to be readed in your tests.

Typical usage:

```go
func Test_ExampleBilling(t *testing.T) {

	cdrtest := testutils.NewTestCDR() // IMPORTANT

	a := golax.NewApi()
	a.Root.Interceptor(cdrtest.InterceptorCdr2Memory()) // IMPORTANT

	BuildYourApi(a)

	s := apitest.New(a)
	s.Request("GET", "/my-url-to-test").Do()

	// IMPORTANT: Do things with cdrtest.Memory[i], for example:
	if 200 != cdrtest.Memory[0].Response.StatusCode {
		t.Error("blah blah blah...")
	}
}
```

## Dependencies

Dependencies for testing are:

* github.com/fulldump/apitest
* github.com/fulldump/golax
* github.com/satori/go.uuid
* gopkg.in/mgo.v2

NOTE: Pinned versions are included in `vendor/*` so tests will always run.

Transitive dependencies for runtime are:

* github.com/fulldump/golax

If `utils.Chan2Mongo` is used, as you can expect, it will be needed also:

* gopkg.in/mgo.v2


## Testing

As simple as:

```sh
git clone "<this-repo>"
make
```
