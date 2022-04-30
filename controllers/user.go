package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/balajiss36/Mongodb/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session // This is uc defined in main.go which is retuned which is passed and ran with the GetUser, CreateUser
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
} // GetSession function returns mongodb.session as s which goes as input in this function
// This function bascially returns a modb session as struct so that it can be accessed by Golang to perform on objects

// Getuser is a struct method because we need to get the values from the mongdb session
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) { // Method function which is received by the function and returns
	id := p.ByName("id") // ps the parameters we send on the API call
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectIdHex(id)                                                          // Convert to hex to be readable by mongdb
	u := models.User{}                                                                   // Value from the db is stored in the User struct defined in models
	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil { // Create a mongodb db on the mongbdb session and collection named Users and finding this id
		w.WriteHeader(404)
		return
	}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj) // In the response w we print the string from output uj after unmarshalling it
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) { // Params is blank because we create{
	u := models.User{}                                 // Save the output here
	json.NewDecoder(r.Body).Decode(&u)                 // request body will be decoded and saved in u
	u.Id = bson.NewObjectId()                          // Create new user id for user
	uc.session.DB("mongo-golang").C("users").Insert(u) // In the mongodb collection insert this user
	ux, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", ux)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id") // Byname is a function in httprouter to get the id by name provided
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(id)
	if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted User", oid)
}
