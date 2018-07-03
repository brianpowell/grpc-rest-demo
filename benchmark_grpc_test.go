package main

import (
	"crypto/tls"
	"log"
	"testing"

	"github.com/brianpowell/grpc-rest-demo/models"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// BenchmarkGRPCGet - Test the GRPC POST connection
func BenchmarkGRPCGet(b *testing.B) {

	// Set up a connection to the server.
	config := &tls.Config{}
	config.InsecureSkipVerify = true
	conn, err := grpc.Dial("localhost:3000", grpc.WithTransportCredentials(credentials.NewTLS(config)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := models.NewVehicleServerClient(conn)

	// run grpc calls against it
	for i := 0; i < b.N; i++ {
		client.Get(context.Background(), &models.VehicleQuery{
			Id: "12345",
		})
	}
}

// BenchmarkGRPCPost - Test the GRPC POST connection
func BenchmarkGRPCPost(b *testing.B) {

	// Set up a connection to the server.
	config := &tls.Config{}
	config.InsecureSkipVerify = true
	conn, err := grpc.Dial("localhost:3000", grpc.WithTransportCredentials(credentials.NewTLS(config)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := models.NewVehicleServerClient(conn)

	// run grpc calls against it
	for i := 0; i < b.N; i++ {
		client.Post(context.Background(), &models.Vehicle{
			Manufacturer: "Audi",
			Model:        "A4",
			Price:        24999.99,
			Mileage:      50000,
		})
	}
}

// BenchmarkGRPCPut - Test the GRPC POST connection
func BenchmarkGRPCPut(b *testing.B) {

	// Set up a connection to the server.
	config := &tls.Config{}
	config.InsecureSkipVerify = true
	conn, err := grpc.Dial("localhost:3000", grpc.WithTransportCredentials(credentials.NewTLS(config)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := models.NewVehicleServerClient(conn)

	// run grpc calls against it
	for i := 0; i < b.N; i++ {
		client.Put(context.Background(), &models.Vehicle{
			Id:           "12345",
			Manufacturer: "Audi",
			Model:        "A4",
			Price:        24999.99,
			Mileage:      50000,
		})
	}
}

// BenchmarkGRPCDel - Test the GRPC POST connection
func BenchmarkGRPCDel(b *testing.B) {

	// Set up a connection to the server.
	config := &tls.Config{}
	config.InsecureSkipVerify = true
	conn, err := grpc.Dial("localhost:3000", grpc.WithTransportCredentials(credentials.NewTLS(config)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := models.NewVehicleServerClient(conn)

	// run grpc calls against it
	for i := 0; i < b.N; i++ {
		client.Del(context.Background(), &models.VehicleQuery{
			Id: "12345",
		})
	}
}
