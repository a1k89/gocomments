package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
	"log"
	"net/http"
	_ "github.com/google/uuid"
	"time"
)

type User struct {
	Uid uuid.UUID `json:"uid"`
	Name string
}

type Comment struct {
	Owner User
	CreatedAt time.Time `json:"createdAt"`
	Body string `json:"body"`
	Uid uuid.UUID `json:"uid"`
}

var db *gorm.DB
var err error

var (
	comments = []Comment {
		{Body: "Hello comment1", Uid:uuid.New()},
		{Body: "Hello comment2", Uid:uuid.New()},
	}
)

func main() {
	router := mux.NewRouter()

	db, err = gorm.Open( "postgres", "host=localhost port=5432 user=postgres dbname=gocomments sslmode=disable password=postgres")
	if err != nil {
		panic("failed to connect DB")
	}
	defer db.Close()

	db.AutoMigrate(&Comment{})
	db.Delete(&Comment{})

	for index := range comments {
		db.Create(&comments[index])
	}

	router.HandleFunc("/comments/", GetComments).Methods("GET")
	router.HandleFunc("/comments/", CreateComment).Methods("POST")
	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}

func GetComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var comments []Comment
	db.Find(&comments)
	json.NewEncoder(w).Encode(&comments)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var c Comment
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}

	db.Save(&c)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&c)
}