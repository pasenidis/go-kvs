package main

import (
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/{key}", KeyReaderHandler).Methods("GET")
	r.HandleFunc("/{key}", KeyWriterHandler).Methods("GET")

	r.Use(mux.CORSMethodMiddleware(r))

	log.Fatal(http.ListenAndServe(":8080", r))
}

// HomeHandler is the home route located at /
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	json := simplejson.New()
	json.Set("name", "go-kvs")
	json.Set("version", 1.0)

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Write(payload)
}

// KeyReaderHandler is the key route located at /{key}
func KeyReaderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	json := simplejson.New()
	json.Set("key", vars["key"])

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Write(payload)
}

// KeyWriterHandler is the key route creator located at /{key}
func KeyWriterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json := simplejson.New()

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Write(payload)
}
