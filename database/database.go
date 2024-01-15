package database

import (
	"database/sql"
	"github.com/KaynHvH/achievify/utils"
	"log"
	
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "responses.database")
	if err != nil {
		log.Fatal(err)
	}
	CreateResponsesTable(db)
	return db
}

func CreateResponsesTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS responses (
		id TEXT,
		response TEXT,
		PRIMARY KEY (id)
	)`)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertResponse(db *sql.DB, response string) (string, error) {
	id := utils.GenerateUniqueID()
	_, err := db.Exec("INSERT INTO responses (id, response) VALUES (?, ?)", id, response)
	if err != nil {
		log.Println("Error inserting response into database:", err)
		return "", err
	}
	return id, nil
}

func GetResponseByID(db *sql.DB, id string) (string, error) {
	var response string
	err := db.QueryRow("SELECT response FROM responses WHERE id = ?", id).Scan(&response)
	if err != nil {
		log.Println("Error getting response from database:", err)
		return "", err
	}
	return response, nil
}
