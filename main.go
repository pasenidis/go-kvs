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
	r.HandleFunc("/{key}", KeyHandler)

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

// KeyHandler is the key route located at /{key}
func KeyHandler(w http.ResponseWriter, r *http.Request) {
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
