package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kvaishak/twitter-server/errors"
	"github.com/kvaishak/twitter-server/services"
)

func GetUserTweets(response http.ResponseWriter, request *http.Request) {

	username := request.URL.Query().Get("username")
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

	postsData, apiErr := services.GetUsersPost(username)
	if apiErr != nil {
		response = handleError(apiErr, response)
		return
	}
	json.NewEncoder(response).Encode(postsData)

}

func GetAllPost(response http.ResponseWriter, request *http.Request) {

	(response).Header().Set("Access-Control-Allow-Origin", "*")

	postsData, apiErr := services.GetAllUsersPost()
	if apiErr != nil {
		response = handleError(apiErr, response)
		return
	}
	json.NewEncoder(response).Encode(postsData)

}

func GetFollowersPost(response http.ResponseWriter, request *http.Request) {

	username := request.URL.Query().Get("username")
	lastTweetId := request.URL.Query().Get("cursor")
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

	if len(lastTweetId) > 0 {
		postsData, apiErr := services.GetFollowersTimedPost(username, lastTweetId)
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

func NewPost(response http.ResponseWriter, request *http.Request) {

	(response).Header().Set("Access-Control-Allow-Origin", "*")

	if request.Method == "POST" {
		reqBody, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Fatal(err)
		}
		isPosted, apiErr := services.NewPost(reqBody)

		if apiErr != nil {
			response = handleError(apiErr, response)
			return
		}

		json.NewEncoder(response).Encode(isPosted)
		return
	}

	apiError := &errors.AppError{
		Message:    "Only POST request Authorized for this URL",
		StatusCode: http.StatusNotFound,
		Status:     "not found",
	}
	response = handleError(apiError, response)

}

func handleError(apiErr *errors.AppError, response http.ResponseWriter) http.ResponseWriter {
	jsonValue, _ := json.Marshal(apiErr)
	response.WriteHeader(apiErr.StatusCode)
	response.Write([]byte(jsonValue))
	return response
}
