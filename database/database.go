package database

import (
    "log"
    "time"
    "github.com/press-play/relation/models"
    "gopkg.in/mgo.v2"
    // "gopkg.in/mgo.v2/bson"
)

var (
    Session *mgo.Session
    Mongo *mgo.DialInfo
)

const (
    Param = "db"
)

func Connect() {
    uri := "mongodb://localhost:27017/relation"
    m, err := mgo.ParseURL(uri)
    s, err := mgo.Dial(uri)
    if err != nil {
        panic(err)
    }

    // https://godoc.org/labix.org/v2/mgo#Session.SetMode
    s.SetMode(mgo.Monotonic, true)

    Mongo = m
    Session = s
}

func PopulateTestData() {
    m := Mongo
    s := Session

    err := s.DB("relation").DropDatabase()
    if err != nil {
        panic(err)
    }

    c := s.DB(m.Database).C("people")
    err = c.Insert(&models.Person{Name: "Khanh Nguyen", Created: time.Now()},
                    &models.Person{Name: "Test User", Created: time.Now()})
    if err != nil {
        log.Fatal(err)
    }
}