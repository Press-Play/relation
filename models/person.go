package models

import (
    "time"
    "gopkg.in/mgo.v2"
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

func PersonFindId(id bson.ObjectId, db *mgo.Database) (record Person, err error) {
    person := Person{}
    err = db.C(PersonCollection).FindId(id).One(&person)
    if err != nil {
        return Person{}, err
    }

    return person, err
}

func PersonInsert(person Person, db *mgo.Database) (record Person, err error) {
    // Create an _id so we can return it.
    person.ID = bson.NewObjectId()

    // Record the time the record was created.
    person.Created = time.Now()

    // Insert the reocrd into database collection.
    err = db.C(PersonCollection).Insert(&person)
    if err != nil {
        return Person{}, err
    }

    // Return the record.
    return person, err
}

// func PersonUpdate(u User, db *mgo.Database) (record User, err error) {}
// func PersonDelete(u User, db *mgo.Database) (record User, err error) {}
