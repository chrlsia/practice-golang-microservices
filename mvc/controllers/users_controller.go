package controllers

import (
	"encoding/json"
	"github.com/chrlsia/practice-golang-microservices/mvc/services"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)

	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("user_id must be a number"))
		//just return the bad request to client
		return
	}

	user, err := services.GetUser(userId)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		//Handle the err and return to the client
		return
	}
	// return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
