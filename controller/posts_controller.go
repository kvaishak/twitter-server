package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kvaishak/twitter-server/services"
)

func GetFollowersPost(response http.ResponseWriter, request *http.Request) {

	username := request.URL.Query().Get("username")
	lastPublishTime := request.URL.Query().Get("time")
	(response).Header().Set("Access-Control-Allow-Origin", "*")

	if len(lastPublishTime) > 0 {
		fmt.Println("Publishtime present")
		userData, apiErr := services.GetFollowersTimedPost(username, lastPublishTime)
		if apiErr != nil {

			jsonValue, _ := json.Marshal(apiErr)
			response.WriteHeader(apiErr.StatusCode)
			response.Write([]byte(jsonValue))
			return
		}
		json.NewEncoder(response).Encode(userData)
	} else {
		fmt.Println("No Publish Time")
		userData, apiErr := services.GetFollowersPost(username)
		if apiErr != nil {

			jsonValue, _ := json.Marshal(apiErr)
			response.WriteHeader(apiErr.StatusCode)
			response.Write([]byte(jsonValue))
			return
		}
		json.NewEncoder(response).Encode(userData)
	}

}
