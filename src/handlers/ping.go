package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

func GetPing() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonBytes, _ := json.Marshal(time.Now())
		w.Write(jsonBytes)
	}

}
