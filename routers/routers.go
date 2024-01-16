package routers

import (
	"database/sql"
	"github.com/KaynHvH/achievify/ai"
	"github.com/KaynHvH/achievify/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/generate/{content}", func(w http.ResponseWriter, r *http.Request) {
		ai.GenerateResponseHandler(w, r, db)
	}).Methods("GET")

	router.HandleFunc("/response/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetResponseHandler(w, r, db)
	}).Methods("GET")

	return router
}
