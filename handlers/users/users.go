package users

import(
    "time"
    "net/http"
    "github.com/press-play/relation/database"
    "github.com/press-play/relation/models"
    "github.com/gin-gonic/gin"
    "gopkg.in/mgo.v2"
)

func Find(c *gin.Context) {}

func Insert(c *gin.Context) {
    db := c.MustGet(database.Param).(*mgo.Database)
    result := models.User{}
    err := c.Bind(&result)
    if err != nil {
        c.Error(err)
        return
    }

    result.Created = time.Now()
    err = db.C(models.UserCollection).Insert(&result)
    if err != nil {
        c.Error(err)
        return
    }

    // TODO: Return _id in the result
    c.JSON(http.StatusOK, result)
}

func Update(c *gin.Context) {}
func Delete(c *gin.Context) {}
