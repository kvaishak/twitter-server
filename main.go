package main

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kvaishak/twitter-server/controller"
)

var mySigningKey = []byte("mysupersecretsigningkey")

func main() {
	http.HandleFunc("/users", controller.GetUsers)
	http.HandleFunc("/user", controller.User)
	http.HandleFunc("/user/auth", controller.UserAuth)
	http.HandleFunc("/user/new", controller.NewUser)

	http.HandleFunc("/user/tweets", controller.GetUserTweets)
	http.HandleFunc("/tweets/new", controller.NewPost)

	http.HandleFunc("/alltweets", controller.GetAllPost)

	http.Handle("/user/follow", isAuthorized(controller.FollowUser))
	http.Handle("/user/unfollow", isAuthorized(controller.UnFollowUser))
	http.Handle("/user/following", isAuthorized(controller.CheckIfFollowing))
	http.Handle("/tweets", isAuthorized(controller.GetFollowersPost))

	if err := http.ListenAndServe(":8282", nil); err != nil {
		panic(err)
	}
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request, string)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Add("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {

			w.Header().Add("Connection", "keep-alive")
			w.Header().Add("Access-Control-Allow-Methods", "GET")
			w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
			// w.Header().Add("Access-Control-Max-Age", "86400")
			// w.Header().Set("Access-Control-Expose-Headers", "token")

			return

		} else {

			if r.Header["Token"] != nil {

				token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Error in parsing the Token")
					}
					return mySigningKey, nil
				})

				tokenString := r.Header["Token"][0]
				claims := jwt.MapClaims{}
				parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
					return mySigningKey, nil
				})

				if err != nil {
					fmt.Fprintf(w, err.Error())
				}

				fmt.Println(parsedToken)
				uid := fmt.Sprintf("%v", claims["uid"])

				// for key, val := range claims {
				// 	fmt.Printf("Key: %v, value: %v\n", key, val)
				// }

				if token.Valid {
					endpoint(w, r, uid)
				}
			} else {
				fmt.Fprintf(w, "Not Authorized")
			}

		}

	})
}
