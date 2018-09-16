package people

import(
    "log"
    "net/http"
    "github.com/press-play/relation/database"
    "github.com/press-play/relation/models"
    "github.com/gin-gonic/gin"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

func PeopleFind(c *gin.Context) {
    db := c.MustGet(database.Param).(*mgo.Database)
    id := bson.ObjectIdHex(c.Param("_id"))

    result := models.Person{}
    err := db.C(models.PersonCollection).FindId(id).One(&result)
    if err != nil {
        log.Fatal(err)
    }

    c.JSON(http.StatusOK, result)
}

func PeopleInsert(c *gin.Context) {}
func PeopleUpdate(c *gin.Context) {}
func PeopleDelete(c *gin.Context) {}