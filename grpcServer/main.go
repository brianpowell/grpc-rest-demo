package grcpful

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"

	"github.com/brianpowell/grpc-rest-demo/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Make the signature match
type vehicleServer struct {
	models.VehicleServerServer
}

// ServerGRPC - Main function to start the Server
func ServerGRPC(addr, certFile, keyFile string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	s := grpc.NewServer(grpc.Creds(credentials.NewTLS(config)))

	models.RegisterVehicleServerServer(s, &vehicleServer{})
	fmt.Println("RGPC Server: ", addr)
	s.Serve(lis)
}

// Get - implements our VehicleServer
func (s *vehicleServer) Get(ctx context.Context, in *models.VehicleQuery) (*models.VehicleReply, error) {

	// Validate the Query
	if err := in.Validate(); err != nil {
		return &models.VehicleReply{
			Success: false,
			Message: err.Error(),
		}, err
	}

	fmt.Println("GET RGPC: ", in.GetId())

	// To Do: Grab Vehicle from DB (not for this demo)
	veh := &models.Vehicle{
		Manufacturer: "Audi",
		Model:        "A4",
		Price:        24999.99,
		Mileage:      50000,
	}

	// Return the Reply
	return &models.VehicleReply{
		Success: true,
		Message: "Vehicle found",
		Data:    veh,
	}, nil
}

// Post - implements our VehicleServer
func (s *vehicleServer) Post(ctx context.Context, in *models.Vehicle) (*models.VehicleReply, error) {

	// Validate the Post Vehicle
	if err := in.Validate(false); err != nil {
		return &models.VehicleReply{
			Success: false,
			Message: err.Error(),
		}, err
	}

	// To Do: Insert into DB (not for this demo)

	// Return the Reply
	return &models.VehicleReply{
		Success: true,
		Message: "Vehicle created",
		Data:    in,
	}, nil
}

// Put - implements our VehicleServer
func (s *vehicleServer) Put(ctx context.Context, in *models.Vehicle) (*models.VehicleReply, error) {

	// Validate the Put Vehicle
	if err := in.Validate(true); err != nil {
		return &models.VehicleReply{
			Success: false,
			Message: err.Error(),
		}, err
	}

	// To Do: Update in DB (not for this demo)

	// Return the Reply
	return &models.VehicleReply{
		Success: true,
		Message: "Vehicle updated",
		Data:    in,
	}, nil
}

// Delete - implements our VehicleServer
func (s *vehicleServer) Del(ctx context.Context, in *models.VehicleQuery) (*models.VehicleReply, error) {

	// Validate the Query
	if err := in.Validate(); err != nil {
		return &models.VehicleReply{
			Success: false,
			Message: err.Error(),
		}, err
	}

	// To Do: Delete Vehicle from DB (not for this demo)

	// Return the Reply
	return &models.VehicleReply{
		Success: true,
		Message: "Vehicle deleted",
	}, nil
}
