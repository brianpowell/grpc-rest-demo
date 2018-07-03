package main

import (
	"github.com/brianpowell/grpc-rest-demo/grpcServer"
	"github.com/brianpowell/grpc-rest-demo/restServer"
)

var (
	key  = "localhost.key"
	cert = "localhost.crt"
)

func main() {
	var block = make(chan struct{})
	go grcpful.ServerGRPC("localhost:3000", cert, key)
	go restful.ServerREST("localhost:3001", cert, key)
	<-block
}
