package api

import (
	"net/http"
	"encoding/json"
)

//Health shows if all is ok
func Health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := struct {
		OK bool
	}{
		OK: true,
	}
	json.NewEncoder(w).Encode(&data)

}
