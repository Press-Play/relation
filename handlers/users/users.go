package users

import(
    "time"
    "net/http"
    "github.com/press-play/relation/database"
    "github.com/press-play/relation/models"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "gopkg.in/mgo.v2"
)

func Find(c *gin.Context) {
    // TODO: Omit password and api token.
    // .Select(bson.M{"password": 0})
}

func Insert(c *gin.Context) {
    db := c.MustGet(database.Param).(*mgo.Database)
    result := models.User{}
    err := c.Bind(&result)
    if err != nil {
        c.Error(err)
        return
    }

    // TODO: Check for duplicate email before inserting.
    result.Created = time.Now()
    pass, err := bcrypt.GenerateFromPassword([]byte(result.Password), bcrypt.MinCost)
    if err != nil {
        c.Error(err)
        return
    }
    result.Password = string(pass)

    err = db.C(models.UserCollection).Insert(&result)
    if err != nil {
        c.Error(err)
        return
    }

    // Remove password from the response.
    result.Password = ""

    // TODO: Return _id in the result
    c.JSON(http.StatusOK, result)
}

func Update(c *gin.Context) {}
func Delete(c *gin.Context) {}
