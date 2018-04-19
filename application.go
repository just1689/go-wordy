package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/just1689/go-wordy/api"
	"github.com/just1689/go-wordy/controller"
	"github.com/just1689/go-wordy/mem"
	log "github.com/sirupsen/logrus"
)

const port = 8080

func main() {

	log.SetFormatter(&log.TextFormatter{})
	mem.Init()
	log.Info(fmt.Sprintf("Loading database"))
	controller.LoadAll()
	log.Info(fmt.Sprintf("Database loaded!"))

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/isWord", api.CORSHandler(api.IsWord)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/health", api.CORSHandler(api.Health)).Methods("POST", "OPTIONS")
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), router))
	log.Info(fmt.Sprintf("Listening [:%v]", port))
}
