package main

import (
	"github.com/brianpowell/demo/grpcServer"
	"github.com/brianpowell/demo/restServer"
)

func main() {
	var block = make(chan struct{})
	go grcpful.ServerGRPC("localhost:3000")
	go restful.ServerREST("localhost:3001")
	<-block
}
