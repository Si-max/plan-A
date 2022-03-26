package main

/*
Simple implementation of the task 1, could have been constructed better with the structs and pointers.
*/

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {

	// get hostname
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	// get request ip

	ip := make(http.Header)
	ip.Set("X-FORWARDED-FOR", "127.0.0.1")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// serve time in RFC format
	json.NewEncoder(w).Encode(time.Now().Format(time.RFC850))

	// serve hostname
	json.NewEncoder(w).Encode(hostname)

	// serve runtime
	json.NewEncoder(w).Encode(runtime.GOOS)

	// serve request ip
	json.NewEncoder(w).Encode(ip.Get("X-Forwarded-For"))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", home)
	err := http.ListenAndServe(":8081", r)
	log.Fatal(err)
}
