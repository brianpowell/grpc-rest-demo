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

var (
	grpcAddr      = "localhost:3000"
	grpcTLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
)

// BenchmarkGRPCGet - Test the GRPC GET connection
func BenchmarkGRPCGet(b *testing.B) {

	// Connect to the Server
	conn, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(credentials.NewTLS(grpcTLSConfig)))
	if err != nil {
		log.Fatalf("Connection Error: %v", err)
	}
	defer conn.Close()
	client := models.NewVehicleServerClient(conn)

	// Run the gRPC calls
	for i := 0; i < b.N; i++ {
		client.Get(context.Background(), &models.VehicleQuery{
			Id: "12345",
		})
	}
}

// BenchmarkGRPCPost - Test the GRPC POST connection
func BenchmarkGRPCPost(b *testing.B) {

	// Connect to the Server
	conn, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(credentials.NewTLS(grpcTLSConfig)))
	if err != nil {
		log.Fatalf("Connection Error: %v", err)
	}
	defer conn.Close()
	client := models.NewVehicleServerClient(conn)

	// Run the gRPC calls
	for i := 0; i < b.N; i++ {
		client.Post(context.Background(), &models.Vehicle{
			Manufacturer: "Audi",
			Model:        "A4",
			Price:        24999.99,
			Mileage:      50000,
		})
	}
}

// BenchmarkGRPCPut - Test the GRPC PUT connection
func BenchmarkGRPCPut(b *testing.B) {

	// Connect to the Server
	conn, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(credentials.NewTLS(grpcTLSConfig)))
	if err != nil {
		log.Fatalf("Connection Error: %v", err)
	}
	defer conn.Close()
	client := models.NewVehicleServerClient(conn)

	// Run the gRPC calls
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

// BenchmarkGRPCDel - Test the GRPC DELETE connection
func BenchmarkGRPCDel(b *testing.B) {

	// Connect to the Server
	conn, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(credentials.NewTLS(grpcTLSConfig)))
	if err != nil {
		log.Fatalf("Connection Error: %v", err)
	}
	defer conn.Close()
	client := models.NewVehicleServerClient(conn)

	// Run the gRPC calls
	for i := 0; i < b.N; i++ {
		client.Del(context.Background(), &models.VehicleQuery{
			Id: "12345",
		})
	}
}
