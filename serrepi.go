package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/sync/singleflight"
)

//GetDht dht data : get greenhouse temperature and humidity
func GetDht(w http.ResponseWriter, r *http.Request) {
	var dhtGroup singleflight.Group
	data, err, shared := dhtGroup.Do("dhtRead", func() (interface{}, error) {
		return ReadDht()
	})

	fmt.Printf("data shared %t", shared)

	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), 500)
	}

	json.NewEncoder(w).Encode(data.(DhtData))
}

func main() {
	fmt.Println("Automatisation de la serre")
	r := mux.NewRouter()

	r.HandleFunc("/dht", GetDht).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
