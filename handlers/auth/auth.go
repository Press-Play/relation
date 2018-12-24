package auth

import(
    "net/http"
    "fmt"
    "crypto/rand"
    "github.com/press-play/relation/models"
    "github.com/press-play/relation/database"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Login struct {
    Email    string `json:"email" bson:"email" binding:"required"`
    Password string `json:"password" bson:"password" binding:"required"`
}

type Token struct {
    Token string `json:"api_token" bson:"api_token"`
}

func RequestToken(c *gin.Context) {
    // Returns the API authentication token given a matching email and password combination.
    // TODO: Rate limiting.
    result := Token{}
    db := c.MustGet(database.Param).(*mgo.Database)

    params := Login{}
    err := c.Bind(&params)
    if err != nil {
        c.Error(err)
        return
    }

    user := models.User{}
    err = db.C(models.UserCollection).Find(bson.M{"email": params.Email}).One(&user)
    if err != nil {
        c.Error(err)
        return
    }

    // Check that the password matches.
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
    if err != nil {
        c.Error(err)
        return
    }

    // Generate and store API authentication token.
    result.Token = generateToken()
    user.APIToken = result.Token
    err = db.C(models.UserCollection).Update(bson.M{"_id": user.ID}, &user)
    if err != nil {
        c.Error(err)
        return
    }

    c.JSON(http.StatusOK, result)
}

func InvalidateToken(c *gin.Context) {
    // Invalidates the API authentication token for the currently logged in user.

}

func generateToken() string {
    b := make([]byte, 24)
    rand.Read(b)
    return fmt.Sprintf("%x", b)
}
