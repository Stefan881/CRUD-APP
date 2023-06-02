package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"log"

	"github.com/gorilla/mux"
)

type Vehicle struct {
	Id    int    `json:"id"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var Vehicles = []Vehicle{
	{1, "toyota", "corolla", 10000},
	{2, "toyota", "canry", 20000},
	{3, "toyota", "civic", 15000},
}

func returnAllCars(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Vehicles)

}

func returnCarsByBrand(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	carM := vars["make"]
	cars := &[]Vehicle{}
	for _, car := range Vehicles {
		if car.Make == carM {
			*cars = append(*cars, car)
		}
	}

	json.NewEncoder(w).Encode(cars)
}

func returnCarsById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	carId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Print("Unable to convert to string")
	}
	for _, car := range Vehicles {
		if car.Id == carId {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(car)

		}
	}

}
func updateCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId,err := strconv.Atoi(vars["id"])
	if err != nil{
		fmt.Println("Unable to convert to string")
	}
	var updatedCar Vehicle
	json.NewDecoder(r.Body).Decode(&updatedCar)
	for k,v := range Vehicles{
		if v.Id == carId{
			Vehicles = append(Vehicles[:k],Vehicles[k + 1:]...)
			Vehicles = append(Vehicles,updatedCar)
		}
	}
	json.NewEncoder(w).Encode(Vehicles)
	w.WriteHeader(http.StatusOK)

}
func createCar(w http.ResponseWriter, r *http.Request) {
	var newCar Vehicle
	json.NewDecoder(r.Body).Decode(&newCar)
	Vehicles = append(Vehicles, newCar)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Vehicles)

}
func removeCarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Unable to convert to string")
	}
	for k,v := range Vehicles{
		if v.Id == carId{
			Vehicles = append(Vehicles[:k],Vehicles[k+1:]...)

		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Vehicles)


}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/cars", returnAllCars).Methods("GET")
	router.HandleFunc("/cars/make/{make}", returnCarsByBrand).Methods("GET")
	router.HandleFunc("/cars/{id}", returnCarsById).Methods("GET")
	router.HandleFunc("/cars/{id}", updateCar).Methods("PUT")
	router.HandleFunc("/cars", createCar).Methods("POST")
	router.HandleFunc("/cars/{id}", removeCarById).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", router))

}
