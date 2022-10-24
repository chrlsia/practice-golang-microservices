package controllers

import (
	"log"
	"net/http"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId := req.URL.Query().Get("user_id")
	log.Printf("about to process user id %v", userId)
}
