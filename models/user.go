package models // Project is divided into different packages so we define the name of the package here

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id     bson.ObjectId `json:"id" bson:"_id"` // This is because the Id will be created automatically by bson(mongodb with json)
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"` // bson is for mongodb to be stored in the db
}
