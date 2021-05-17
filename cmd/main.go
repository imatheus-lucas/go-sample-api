package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/imatheus-lucas/go-sample-api/pkg/config"
	"github.com/imatheus-lucas/go-sample-api/pkg/http/routes"
	"github.com/imatheus-lucas/go-sample-api/pkg/utils"
)

func main() {
	router := mux.NewRouter()

	routes.Register(
		router.PathPrefix("/api").Subrouter(),
	)
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		json.NewEncoder(w).Encode(utils.ResponseError("error 404"))
	})

	server := http.Server{
		Addr:         ":3333",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      config.Cors(router),
	}
	log.Fatal(server.ListenAndServe())
}
