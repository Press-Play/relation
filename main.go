package main

import (
    "github.com/press-play/relation/database"
    "github.com/press-play/relation/handlers/people"
    "github.com/press-play/relation/handlers/users"
    "github.com/gin-gonic/gin"
)

func main() {
    database.Connect()

    r := gin.Default()

    r.Use(Database)

    r.GET("/ping", ping)

    v1 := r.Group("/api/v1")
    {
        v1.GET("/people/:_id", people.Find)
        v1.POST("/people/new", people.Insert)

        v1.POST("/users", users.Insert)
    }
    r.Run(":8000")
}

func Database(c *gin.Context) {
    s := database.Session.Clone()
    defer s.Close()

    c.Set(database.Param, database.Session.DB(database.Mongo.Database))
    c.Next()
}

func ping(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "pong",
    })
}