package models

import (
    "time"
    "gopkg.in/mgo.v2/bson"
)

const (
    PersonCollection = "people"
)

type Person struct {
    ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
    Name      string `json:"name,omitempty" bson:"name,omitempty" binding:"required"`
    Relations []bson.ObjectId `json:"relations" bson:"relations"`
    UserID    bson.ObjectId `json:"user_id,omitempty" bson:"user_id,omitempty"`
    Created   time.Time `json:"created" bson:"created"`
}
