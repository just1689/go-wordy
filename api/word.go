package api

import (
	"encoding/json"
	"fmt"
	"github.com/just1689/go-wordy/controller"
	"github.com/just1689/go-wordy/model"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//IsWord checks if a word is a word
func IsWord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var word model.Word
	err := decoder.Decode(&word)
	log.Infoln(fmt.Sprintf("Looking up: %s", word.ID))
	if err != nil {
		log.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", err))
		json.NewEncoder(w).Encode(model.Result{Exists: false, Error: err})
		return
	}
	e, err := controller.Exists(word.ID)
	log.Infoln(fmt.Sprintf("Found: %v", e))
	json.NewEncoder(w).Encode(model.Result{Exists: e, Error: err})
}

//CORSHandler does CORS check
func CORSHandler(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-type")
		w.Header().Set("Allow", "GET,POST,OPTIONS")

		if r.Method == "OPTIONS" {
			fmt.Fprintf(w, "Options")
		} else {
			f.ServeHTTP(w, r)
		}
	}
}
