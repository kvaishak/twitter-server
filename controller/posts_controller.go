package controller

import (
	"encoding/json"
	"net/http"

	"github.com/kvaishak/twitter-server/errors"
	"github.com/kvaishak/twitter-server/services"
)

func GetFollowersPost(response http.ResponseWriter, request *http.Request) {

	username := request.URL.Query().Get("username")
	lastPublishTime := request.URL.Query().Get("time")
	(response).Header().Set("Access-Control-Allow-Origin", "*")

	if len(username) == 0 {
		apiError := &errors.AppError{
			Message:    "Please send a username to fetch the posts",
			StatusCode: http.StatusNotFound,
			Status:     "not found",
		}
		response = handleError(apiError, response)
		return
	}

	if len(lastPublishTime) > 0 {
		postsData, apiErr := services.GetFollowersTimedPost(username, lastPublishTime)
		if apiErr != nil {
			response = handleError(apiErr, response)
			return
		}
		json.NewEncoder(response).Encode(postsData)
	} else {
		postsData, apiErr := services.GetFollowersPost(username)
		if apiErr != nil {
			response = handleError(apiErr, response)
			return
		}
		json.NewEncoder(response).Encode(postsData)
	}

}

func handleError(apiErr *errors.AppError, response http.ResponseWriter) http.ResponseWriter {
	jsonValue, _ := json.Marshal(apiErr)
	response.WriteHeader(apiErr.StatusCode)
	response.Write([]byte(jsonValue))
	return response
}
