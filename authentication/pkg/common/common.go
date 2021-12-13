package handler

import (
	"encoding/json"
	"net/http"
)

func ServeJSON(wr http.ResponseWriter, value interface{}) {

	jsonRecord, err := json.Marshal(value)

	if err != nil {
		http.Error(wr, "Internal server error", http.StatusInternalServerError)
		return
	}

	wr.Header().Set("Content-Type", "application/json")
	wr.Write(jsonRecord)
}
