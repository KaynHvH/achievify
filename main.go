package main

import (
	"database/sql"
	"github.com/KaynHvH/achievify/database"
	"github.com/KaynHvH/achievify/routers"
	"log"
	"net/http"
)

func main() {
	db := database.InitDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("Error closing database:", err)
			return
		}
	}(db)

	router := routers.NewRouter(db)

	log.Println("Server works properly")
	if err := http.ListenAndServe(":3030", router); err != nil {
		log.Fatal(err)
	}
}
