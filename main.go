package main

import (
	"net/http"

	"github.com/kvaishak/twitter-server/controller"
)

func main() {
	http.HandleFunc("/users", controller.GetUsers)
	http.HandleFunc("/user", controller.User)

	if err := http.ListenAndServe(":8282", nil); err != nil {
		panic(err)
	}
}
