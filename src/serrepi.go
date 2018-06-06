package main

import (
	"dht"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Automatisation de la serre")

	//gestion exit
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	tickerDht := time.NewTicker(15 * time.Minute)
	r := mux.NewRouter()
	srv := &http.Server{Addr: ":3000", Handler: r}

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		tickerDht.Stop()
		srv.Shutdown(nil)
		done <- true
	}()

	fmt.Println("Demarrage de lecture periodique de la temperature/humidite")
	data := dht.ReadDht()
	go func() {
		for range tickerDht.C {
			data = dht.ReadDht()
		}
	}()

	//serveur web

	r.HandleFunc("/dht", func(w http.ResponseWriter, r *http.Request) {
		b, _ := json.Marshal(data)
		w.Write(b)
	}).Methods("GET")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	<-done
	fmt.Println("exiting")
}
