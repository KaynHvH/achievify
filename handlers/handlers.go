package handlers

import (
	"database/sql"
	"github.com/KaynHvH/achievify/database"
	"github.com/KaynHvH/achievify/models"
	"github.com/KaynHvH/achievify/utils"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func GetResponseHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	response, err := database.GetResponseByID(db, id)
	if err != nil {
		log.Println("Error getting response from database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	points := utils.SplitResponse(response)

	var pointList []models.Point
	for i, point := range points {
		pointList = append(pointList, models.Point{ID: i + 1, Point: point})
	}

	tmpl, err := template.ParseFiles("./static/html/response.html")
	if err != nil {
		log.Println("Error parsing HTML template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, pointList)
	if err != nil {
		log.Println("Error executing HTML template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
