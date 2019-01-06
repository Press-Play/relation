package models

import (
    "time"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "golang.org/x/crypto/bcrypt"
    "github.com/press-play/relation/utils"
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

func UserFindOne(query bson.M, db *mgo.Database) (record User, err error) {
    user := User{}
    err = db.C(UserCollection).Find(query).One(&user)
    if err != nil {
        return User{}, err
    }

    return user, err
}

func UserInsert(user User, db *mgo.Database) (record User, err error) {
    // Create an _id so we can return it.
    user.ID = bson.NewObjectId()

    // Record the time the record was created.
    user.Created = time.Now()

    // Hash the password.
    pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
    if err != nil {
        return User{}, err
    }
    user.Password = string(pass)

    // Generate API authentication token for use immediately.
    user.APIToken = utils.GenerateToken()

    // Insert the reocrd into database collection.
    err = db.C(UserCollection).Insert(&user)
    if err != nil {
        return User{}, err
    }

    // Remove password from the response.
    user.Password = ""

    // Return the record.
    return user, err
}

// func UserUpdate(user User, db *mgo.Database) (record User, err error) {}
// func UserDelete(user User, db *mgo.Database) (record User, err error) {}
