This project strives to show matching gRPC and RESTful API implementations that allow users to see the benchmarks via each approach. The approach assumes the need to Marshal the HTTP calls into structs via their JSON definitions.

## Download the project
`$ go get github.com/brianpowell/grpc-rest-demo`

Head into the Go Project folder
## Step into the directory
`$ cd $GOPATH/src/github.com/brianpowell/grpc-rest-demo`

We need some certs to make the TLS work for the Routers.
## Generate a Self-signed Certificate
`$ openssl req  -newkey rsa:2048 -nodes -keyout localhost.key -x509 -days 365 -out localhost.crt`

**Make sure the `Common Name:` is `localhost`**

This will produce a `localhost.key` and `localhost.crt` in the root directory of the project. These will be used with the connections to provide TLS services.

Now that we have all the dependencies, we need to start the servers. We will need two terminal windows.
## Run the Server (window 1)
`$ go run main.go`

Now that we have the servers up, we can run the benchmarks to compare the performance.
## Run Bechmark Tests (window 2)
`$ go test -run=Bench* -benchtime=2s -bench=.`
