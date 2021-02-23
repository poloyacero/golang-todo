package controllers

import (
	"encoding/json"
	"fmt"
	"go-todo/db"
	"go-todo/middlewares"
	"go-todo/models"
	response "go-todo/utility"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var users []models.User

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	fmt.Println("Create User")

	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	password := []byte(user.Password)

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	newUser := models.User{
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	fmt.Println(user.Email, hashedPassword)

	db.InsertUser(newUser)
	response.JSON(w, newUser)
}

func SignInUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("USER", user)
	foundUser := db.GetUser(user)

	pwdMatch := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if pwdMatch == nil {
		token, err := middlewares.GenerateJWT(foundUser)
		if err != nil {
			log.Fatal(err)
		}
		response.JSON(w, token)
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	// array of split the url
	url := strings.Split(r.URL.String(), "/")

	paramId := url[2]

	fmt.Println("PARAMS", paramId)

	if len(url) != 3 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("404 Not Found")
		return
	}

	switch r.Method {
	case "POST":
		fmt.Println("here")
		if paramId == "signup" {
			CreateUser(w, r)
		} else {
			SignInUser(w, r)
		}

		return
	case "GET":

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode("Method not allowed")
		return
	}
}
