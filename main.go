package main

import (
	"github.com/KaynHvH/achievify/ai"
	"github.com/KaynHvH/achievify/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/generate/{content}", ai.GenerateResponse).Methods("GET")
	router.HandleFunc("/response/{id}", handlers.GetResponse).Methods("GET")

	log.Println("Server works properly")
	if err := http.ListenAndServe(":3030", router); err != nil {
		log.Fatal(err)
	}
}
