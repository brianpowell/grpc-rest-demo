package restful

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/brianpowell/grpc-rest-demo/models"
	"github.com/go-chi/chi"
)

// ServerREST - Main function to start the Server
func ServerREST(addr string) {
	r := chi.NewRouter()
	r.Get("/vehicle/{id}", Get)
	r.Post("/vehicle", Post)
	r.Put("/vehicle/{id}", Put)
	r.Delete("/vehicle/{id}", Del)
	fmt.Println("REST Server: ", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

// Get - Rest HTTP Handler
func Get(w http.ResponseWriter, r *http.Request) {

	// Stub the response values
	code := 200
	resp := models.VehicleReply{
		Success: true,
		Message: "Vehicle deleted",
	}

	// Grab URL value
	id := chi.URLParam(r, "id")
	if id == "" {
		code = 404
		resp.Success = false
		resp.Message = "Id must be present"
	} else {
		// To Do: Grab Vehicle from DB (not for this demo)
		resp.Data = &models.Vehicle{
			Manufacturer: "Audi",
			Model:        "A4",
			Price:        24999.99,
			Mileage:      50000,
		}
	}

	// Kick out the response
	respBytes, _ := json.Marshal(resp)
	w.WriteHeader(code)
	w.Write(respBytes)
}

// Post - Rest HTTP Handler
func Post(w http.ResponseWriter, r *http.Request) {

	veh := models.Vehicle{}

	// decode input
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&veh)
	r.Body.Close()

	// Stub the response values
	code := 200
	resp := models.VehicleReply{
		Success: true,
		Message: "Vehicle created",
	}

	// Validate input
	if err := veh.Validate(false); err != nil {
		code = 500
		resp.Success = false
		resp.Message = err.Error()
	} else {
		resp.Data = &veh
	}

	// To Do: Insert into DB (not for this demo)

	// Kick out the response
	respBytes, _ := json.Marshal(resp)
	w.WriteHeader(code)
	w.Write(respBytes)
}

// Put - Rest HTTP Handler
func Put(w http.ResponseWriter, r *http.Request) {

	veh := models.Vehicle{}

	// decode input
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&veh)
	r.Body.Close()

	// Stub the response values
	code := 200
	resp := models.VehicleReply{
		Success: true,
		Message: "Vehicle updated",
	}

	// Validate input
	if err := veh.Validate(true); err != nil {
		code = 404
		resp.Success = false
		resp.Message = err.Error()
	} else {
		resp.Data = &veh
	}

	// To Do: Update in DB (not for this demo)

	// Kick out the response
	respBytes, _ := json.Marshal(resp)
	w.WriteHeader(code)
	w.Write(respBytes)
}

// Del - Rest HTTP Handler
func Del(w http.ResponseWriter, r *http.Request) {

	// Stub the response values
	code := 200
	resp := models.VehicleReply{
		Success: true,
		Message: "Vehicle deleted",
	}

	// Grab URL value
	id := chi.URLParam(r, "id")
	if id == "" {
		code = 404
		resp.Success = false
		resp.Message = "Id must be present"
	} else {
		// To Do: Delete Vehicle from DB (not for this demo)
	}

	// Kick out the response
	respBytes, _ := json.Marshal(resp)
	w.WriteHeader(code)
	w.Write(respBytes)
}
