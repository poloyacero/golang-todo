package main

import (
	"fmt"
	"go-todo/controllers"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("opensesame")

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Headers", r.Header)
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]
		if r.Header["Authorization"] != nil {
			token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			fmt.Println("TOKEN", token)

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func main() {
	fmt.Println("Hello World")
	http.Handle("/tasks", isAuthorized(controllers.TaskHandler))
	http.Handle("/tasks/", isAuthorized(controllers.TaskSpecificHandler))
	http.HandleFunc("/auth/", controllers.UserHandler)
	http.Handle("/test", isAuthorized(test))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
