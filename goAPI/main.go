package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Car struct {
	ID      string `json:"id"`
	Brand   string `json:"brand"`
	Model   string `json:"model"`
	Mileage string `json:"mileage"`
	Owners  string `json:"owners"`
}

var cars []Car

func getCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func getCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range cars {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Car{})
}

func createCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var car Car
	_ = json.NewDecoder(r.Body).Decode(&car)
	car.ID = strconv.Itoa(rand.Intn(1000000))
	cars = append(cars, car)
	json.NewEncoder(w).Encode(car)
}

func updateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range cars {
		if item.ID == params["id"] {
			cars = append(cars[:index], cars[index+1:]...)
			var car Car
			_ = json.NewDecoder(r.Body).Decode(&car)
			car.ID = params["id"]
			cars = append(cars, car)
			json.NewEncoder(w).Encode(car)
			return
		}
	}
	json.NewEncoder(w).Encode(cars)
}

func deleteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range cars {
		if item.ID == params["id"] {
			cars = append(cars[:index], cars[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(cars)
}

func main() {
	r := mux.NewRouter()
	cars = append(cars, Car{ID: "1", Brand: "Lada", Model: "Kalina", Mileage: "100", Owners: "1"})
	cars = append(cars, Car{ID: "2", Brand: "Audi", Model: "911", Mileage: "210", Owners: "3"})
	r.HandleFunc("/cars", getCars).Methods("GET")
	r.HandleFunc("/cars/{id}", getCar).Methods("GET")
	r.HandleFunc("/cars", createCar).Methods("POST")
	r.HandleFunc("/cars/{id}", updateCar).Methods("PUT")
	r.HandleFunc("/cars/{id}", deleteCar).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
