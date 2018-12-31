package users

import(
    "net/http"
    "github.com/press-play/relation/database"
    "github.com/press-play/relation/models"
    "github.com/gin-gonic/gin"
    "gopkg.in/mgo.v2"
)

func Find(c *gin.Context) {
    // TODO: Omit password and api token.
    // .Select(bson.M{"password": 0})
}

func CreateAccount(c *gin.Context) {
    db := c.MustGet(database.Param).(*mgo.Database)
    user := models.User{}
    err := c.Bind(&user)
    if err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    result, err := models.UserInsert(user, db)
    if err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    c.JSON(http.StatusOK, result)
}

func Update(c *gin.Context) {}
func Delete(c *gin.Context) {}
