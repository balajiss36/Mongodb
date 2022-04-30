package main

import (
	"net/http"

	"github.com/balajiss36/Mongodb/controllers"
	"github.com/julienschmidt/httprouter" // This is similar to gorilla packages which routes the request as per the API call
	"gopkg.in/mgo.v2"
)

func main() { // Entry point to the project
	r := httprouter.New() // New function is present inside the httprouter
	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r) // Listen to port and put it to r.
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27107") // This is to connect with mongodb on port 27107
	//which is the default one and returns the session in the NewUserController function of controller
	if err != nil {
		panic(err)
	}
	return s // Returns the mongodb connection.
}
