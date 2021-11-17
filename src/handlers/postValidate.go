package handlers

import (
	"encoding/json"
	"evaluationapp/src/requests"
	"evaluationapp/src/response"
	"evaluationapp/src/services"
	"io/ioutil"
	"net/http"
)

func PostValidate(s services.IEvaluationService, c services.ICollectorService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		defer r.Body.Close()

		var req requests.MathExpression
		err1 := json.Unmarshal(bodyBytes, &req)
		if err1 != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		res, err := s.ValidateExpression(req.Expression)

		var result response.ValidatedMathExpression

		if err != nil {
			url := r.URL.Path
			c.AddError(req.Expression, err.Error(), url)
			result.Reason = err.Error()
		}

		result.Valid = res

		jsonBytes, err := json.Marshal(result)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	}
}
