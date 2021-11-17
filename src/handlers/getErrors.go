package handlers

import (
	"encoding/json"
	"evaluationapp/src/services"
	"net/http"
)

func GetErrors(s services.ICollectorService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		list := s.GetAllErrors()
		jsonBytes, err := json.Marshal(list)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	}
}
