package main

import (
	"evaluationapp/src/handlers"
	"evaluationapp/src/services"
	"net/http"
	"os"
)

func main() {

	srv := services.NewEvaluationService()
	c := services.NewCollectService()

	os.Setenv("PORTENV", ":5000")
	PORT := os.Getenv("PORTENV")

	http.HandleFunc("/ping", handlers.GetPing())
	http.HandleFunc("/errors", handlers.GetErrors(c))
	http.HandleFunc("/evaluate", handlers.PostEvaluate(srv, c))
	http.HandleFunc("/validate", handlers.PostValidate(srv, c))
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		panic(err)
	}
}
