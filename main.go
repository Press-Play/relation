package main

import (
    "log"
    "github.com/press-play/relation/database"
    "github.com/press-play/relation/models"
    "github.com/press-play/relation/handlers/auth"
    "github.com/press-play/relation/handlers/people"
    "github.com/press-play/relation/handlers/users"
    "github.com/gin-gonic/gin"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

func main() {
    database.Connect()

    r := gin.Default()

    r.Use(Database)

    r.GET("/ping", ping)

    public := r.Group("/api/v1")
    {
        public.POST("/auth/login", auth.RequestToken)
        public.POST("/users", users.CreateAccount)
    }

    v1 := r.Group("/api/v1")
    v1.Use(Authentication)
    {
        v1.GET("/people/:_id", people.Find)
        v1.POST("/people/new", people.Insert)

        v1.POST("/logout", auth.InvalidateToken)
    }

    r.Run(":8000")
}

func Database(c *gin.Context) {
    s := database.Session.Clone()
    defer s.Close()

    c.Set(database.Param, database.Session.DB(database.Mongo.Database))
    c.Next()
}

func Authentication(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    userID := bson.ObjectIdHex(c.Request.Header.Get("User-ID"))
    db := c.MustGet(database.Param).(*mgo.Database)

    user, err := models.UserFindId(userID, db)
    if err != nil {
        c.AbortWithError(403, err)
        return
    }

    // Return 403 if the token does not match.
    log.Print("user.APIToken: ", user.APIToken)
    if user.APIToken != token {
        c.AbortWithStatus(403)
    }

    c.Next()
}

func ping(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "pong",
    })
}