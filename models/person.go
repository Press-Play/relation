package models

import (
    "time"
    "gopkg.in/mgo.v2/bson"
)

const (
    PersonCollection = "people"
)

type Person struct {
    ID       bson.ObjectId `bson:"_id,omitempty"`
    Name      string
    Email     string
    Relations []bson.ObjectId
    Created   time.Time
}
