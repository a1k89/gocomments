package main

import (
	"gocomments/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/goapi/comments/{content_type}/", api.CreateCommentHandler).Methods("POST")
	router.HandleFunc("/goapi/comments/{content_type}/", api.GetCommentsHandler).Methods("GET")
	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
