PROJECT=github.com/smartdigits/gocdr
GOPATH=$(shell pwd)
GO=go
GOCMD=GOPATH=$(GOPATH) $(GO)

.PHONY: test all clean dependencies setup example

all: test

clean:
	rm -fr src/
	rm -fr bin/
	rm -fr pkg/

setup:
	mkdir -p src/$(PROJECT)
	rm -fr src/$(PROJECT)
	ln -s ../../.. src/$(PROJECT)
	$(GOCMD) get github.com/fulldump/golax
	$(GOCMD) get github.com/fulldump/apitest
	$(GOCMD) get gopkg.in/mgo.v2

test:
	$(GOCMD) test $(PROJECT)
