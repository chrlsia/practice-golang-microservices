package controllers

import (
	"encoding/json"
	"github.com/chrlsia/practice-golang-microservices/mvc/services"
	"github.com/chrlsia/practice-golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	//Sto req, sto URL,Query() kai Get("user_id")
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number not found",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		//just return the bad request to client
		return
	}

	user, apiErr := services.GetUser(userId)
	//apiErr einai enaw pointer
	//ara exei ennoia na einai nil
	//epeomnvw mporv na to sygkrinv sthn epomenh entolh me nill
	//kai aytow einai o logow poy to pernav vw pointer
	//stiw diaforew functions
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))

		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}

/*
 curl localhost:8080/users?user_id=123 -v
{"id":1,"first_name":"Federico","last_name":"Leon","email":"myemail@gmail.com"}

 curl localhost:8080/users?user_id=123a -v
{"message":"user_id must be a number not found","status":400,"code":"bad request"}

 curl localhost:8080/users?user_id=1234 -v
{"message":"user 1234 was not found","status":404,"code":"not found"}
*/
