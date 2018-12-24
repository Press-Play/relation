package models

import (
    "time"
    "gopkg.in/mgo.v2/bson"
)

const (
    UserCollection = "users"
)

type User struct {
    ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
    Password  string `json:"password,omitempty" bson:"password,omitempty" binding:"required"`
    Email     string `json:"email,omitempty" bson:"email,omitempty" binding:"required"`
    APIToken  string `json:"api_token,omitempty" bson:"api_token,omitempty"`
    Created   time.Time `json:"created" bson:"created"`
}
