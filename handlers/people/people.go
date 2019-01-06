package people

import(
    "net/http"
    "github.com/press-play/relation/database"
    "github.com/press-play/relation/models"
    "github.com/gin-gonic/gin"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

func Find(c *gin.Context) {
    db := c.MustGet(database.Param).(*mgo.Database)
    id := bson.ObjectIdHex(c.Param("_id"))

    result, err := models.PersonFindId(id, db)
    if err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    c.JSON(http.StatusOK, result)
}

func Insert(c *gin.Context) {
    db := c.MustGet(database.Param).(*mgo.Database)
    person := models.Person{}
    err := c.Bind(&person)
    if err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    result, err := models.PersonInsert(person, db)
    if err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    // TODO: Return _id in the result
    c.JSON(http.StatusOK, result)
}

func Update(c *gin.Context) {}
func Delete(c *gin.Context) {}
