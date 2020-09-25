package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kvaishak/twitter-server/services"
)

func GetUsers(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Haha Haha")

	usersArr, err := services.GetUsers()

	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(err.Error()))
	}
	json.NewEncoder(response).Encode(usersArr)
}
