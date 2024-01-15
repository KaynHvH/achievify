package ai

import (
	"database/sql"
	"encoding/json"
	"fmt"
	database "github.com/KaynHvH/achievify/database"
	"github.com/KaynHvH/achievify/utils"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const (
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)

func GenerateResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	client := resty.New()
	userContent := mux.Vars(r)["content"]

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	response, err := client.R().
		SetAuthToken(os.Getenv("TOKEN")).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model": "gpt-3.5-turbo",
			"messages": []interface{}{
				map[string]interface{}{"role": "user", "content": "Make the points that I need to accomplish to achieve the set goal that I will give at the end. Write your answer in bullet points, without additional comments. For example, if my goal is to write a book, tell me: 1. Choose the topic you would like to write about 2. Choose what should be in the book, etc. My goal is: " + userContent},
			},
			"max_tokens": 150,
		}).
		Post(apiEndpoint)

	if err != nil {
		log.Fatalf("Error while sending the request: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	body := response.Body()

	var data map[string]interface{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error while decoding JSON response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	choices, ok := data["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		fmt.Println("No choices found in the response.")
		http.Error(w, "No choices found in the response", http.StatusInternalServerError)
		fmt.Println("DATA:", data)
		fmt.Println("OK:")
		return
	}

	content := choices[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	content = utils.RemoveNewlines(content)

	db := database.InitDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("Error closing database:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}(db)

	id, err := database.InsertResponse(db, content)
	if err != nil {
		log.Println("Error inserting response into database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("Response successfully inserted into the database with ID:", id)

	_, err = w.Write([]byte(id))

	//<-- OLD IDEA WITH BUTTON TO MOVE USER TO PAGE WITH RESPONSE -->
	//htmlTemplate := `<a href="{{.URL}}">Generate</a>`
	//
	//tmpl, err := template.New("link").Parse(htmlTemplate)
	//if err != nil {
	//	log.Println("Error parsing HTML template:", err)
	//	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	//	return
	//}
	//
	//err = tmpl.Execute(w, struct {
	//	URL string
	//}{
	//	URL: responseData.URL,
	//})
	//
	//if err != nil {
	//	log.Println("Error executing HTML template:", err)
	//	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	//	return
	//}
}
