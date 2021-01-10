package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kvaishak/twitter-server/errors"
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

func UserAuth(response http.ResponseWriter, request *http.Request) {

	uid := request.URL.Query().Get("uid")

	userJWT, apiErr := services.GetUserJWT(uid)

	if apiErr != nil {

		jsonValue, _ := json.Marshal(apiErr)
		response.Write([]byte(jsonValue))
		return
	}

	(response).Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(response).Encode(userJWT)
}

func FollowUser(response http.ResponseWriter, request *http.Request, uid string) {

	followerName := request.URL.Query().Get("followerName")

	isFollowed, apiErr := services.FollowUser(uid, followerName)

	if apiErr != nil {

		jsonValue, _ := json.Marshal(apiErr)
		response.Write([]byte(jsonValue))
		return
	}

	json.NewEncoder(response).Encode(isFollowed)
}

func CheckIfFollowing(response http.ResponseWriter, request *http.Request, uid string) {

	followerName := request.URL.Query().Get("followerName")

	isFollowed, apiErr := services.IsFollowing(uid, followerName)

	if apiErr != nil {

		jsonValue, _ := json.Marshal(apiErr)
		response.Write([]byte(jsonValue))
		return
	}

	json.NewEncoder(response).Encode(isFollowed)
}

func NewUser(response http.ResponseWriter, request *http.Request) {

	// Setting Cors and preflight to responses
	(response).Header().Set("Access-Control-Allow-Origin", "*")
	(response).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(response).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(response).Header().Set("Content-Type", "text/html; charset=ascii")

	if request.Method == "POST" {
		reqBody, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Fatal(err)
		}
		isCreated, apiErr := services.CreateUser(reqBody)

		if apiErr != nil {
			jsonValue, _ := json.Marshal(apiErr)
			response.WriteHeader(apiErr.StatusCode)
			response.Write([]byte(jsonValue))
			return
		}

		json.NewEncoder(response).Encode(isCreated)
		return
	}

	apiError := errors.AppError{
		Message:    "Only POST request Authorized for this URL",
		StatusCode: http.StatusNotFound,
		Status:     "not found",
	}
	jsonValue, _ := json.Marshal(apiError)
	response.WriteHeader(apiError.StatusCode)
	response.Write([]byte(jsonValue))

}
