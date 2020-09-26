package controller

import (
	"encoding/json"
	"net/http"

	"github.com/kvaishak/twitter-server/services"
)

func GetUsers(response http.ResponseWriter, request *http.Request) {

	usersArr, apiErr := services.GetUsers()

	if apiErr != nil {

		jsonValue, _ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		response.Write([]byte(jsonValue))
		return
	}

	(response).Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(response).Encode(usersArr)
}

func User(response http.ResponseWriter, request *http.Request) {

	username := request.URL.Query().Get("username")
	userData, apiErr := services.GetUserData(username)

	if apiErr != nil {

		jsonValue, _ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		response.Write([]byte(jsonValue))
		return
	}

	(response).Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(response).Encode(userData)
}
