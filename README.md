# Introduction

A very simple payments REST API. The API is backed by an in memory store for persisting data. It goes without saying this won't
survive a server restart but it shows an example of the persistence interface that other stores could implement.

## Getting Started

Have:
- Go version 1.11 (it may work with earlier versions but it was built with this)
- Something to send test requests. `curl` would do!

## Run

Steps are to clone the repo and then build the code.

```
cd $GOPATH/src
mkdir -p github.com/cdempsie
cd github.com/cdempsie
git clone https://github.com/cdempsie/payments-example
go get -t ./...
cd server
go run server.go
```

The server will start on port 8000 by default, to change this do:

```
go run server.go -port 8888
```

## Run The Tests

You can run the unit tests using (server does not need to be running):

```
go test -cover ./...
```
