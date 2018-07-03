This project strives to show matching gRPC and RESTful API implementations that allow users to see the benchmarks via each approach.

## Generate a Self-signed Certificate
```
$ cd /grpc-rest-demo
$ openssl req  -newkey rsa:2048 -nodes -keyout localhost.key -x509 -days 365 -out localhost.crt
```
*Make sure the `Common Name:` is `localhost`*

This will produce a `localhost.key` and `localhost.crt` in the root directory of the project. These will be used with the connections to provide TLS services.

Now that we are

## Run Bechmark Tests
`go test -run=Bench* -benchtime=2s -bench=.`
