# gocdr

CDRs SDK for SmartDigits billing system


<!-- MarkdownTOC autolink=true bracket=round depth=4 -->

- [InterceptorCdr](#interceptorcdr)
	- [Custom info](#custom-info)
	- [Access](#access)
	- [Sample CDR](#sample-cdr)
- [InterceptorCdr2Log](#interceptorcdr2log)

<!-- /MarkdownTOC -->


Gocdr is a set of interceptors to generate and store CDRs:

* [InterceptorCdr](#interceptorcdr) - Generate and fill a CDR
* InterceptorCdr2Log - Log CDRs
* InterceptorCdr2Channel - Collect all CDRs in a channel

And some helpers:

* GetCdr - Retrieve CDR from context
* NewTestCDR - Store CDRs in memory to allow unit testing


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
TODO: PUT TIMESTAMP HERE
CDR {"id":"580a465bce507629a613107c","version":"1.0.0","consumer_id":"my-consumer-id","origin":"127.0.0.1","session_id":"","service":"invented-service","entry_date":"2016-10-21T18:46:19.820423299+02:00","entry_timestamp":1.4770683798204234e+09,"elapsed_seconds":1.621246337890625e-05,"request":{"method":"POST","uri":"/value-1/value-2/test-node?query_a=aaa\u0026query_b=bbb","handler":"/{param1}/{param2}/test-node","args":{"query_a":["aaa"],"query_b":["bbb"]},"length":2},"response":{"status_code":222,"length":5,"error":{"code":27,"description":"my-error-description"}},"read_access":["other-involved-consumer-id","my-consumer-id"],"custom":{"a":20,"b":55}}
```