package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Pickausernaame/Ambitine_backend/server/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Hello struct {
	Msg string `json:"msg"`
}

func HelloFunc(c *gin.Context) {
	res := Hello{Msg: "Hi, my dear friend!"}
	c.JSON(200, res)
}

func (instance *App) SignInHand(c *gin.Context) {
	var (
		loginUser models.SignInUserStruct
	)

	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&loginUser)

	if err != nil {
		fmt.Println("Unable to decode SignIn request:", err)
		c.Status(400)
		return
	}

	id, err := instance.DB.GetUserIdByNicknameAndPassword(loginUser)

	if err != nil {
		fmt.Println("Unable to check user existing:", err)
		c.Status(400)
		return
	}

	sessionId := instance.createSessionId(id)
	c.SetCookie("session_id", sessionId, 3600, "/", "", false, false)
	c.Status(201)
}

func (instance *App) SignUpHand(c *gin.Context) {
	var (
		newUser   models.SignUpUserStruct
		loginUser models.SignInUserStruct
	)

	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&newUser)
	if err != nil || !newUser.Validation() {
		fmt.Println("Unable to decode SignUp request: ", err)
		c.Status(400)
		return
	}

	err = instance.DB.InsertNewUser(newUser)

	if err != nil {
		fmt.Println("Unable to insert new user to database:", err)
		c.Status(409)
		return
	}

	loginUser.Nickname = newUser.Nickname
	loginUser.Password = newUser.Password

	err, id := instance.DB.CheckUserExist(newUser.Nickname, newUser.Email)

	if err != nil {
		fmt.Println("Unable to check user existing:", err)
		c.Status(400)
		return
	}

	sessionId := instance.createSessionId(id)
	c.SetCookie("session_id", sessionId, 3600, "/", "", false, false)
	c.Status(201)
}

func (instance *App) createSessionId(id int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})
	// ToDo: Error handle
	spiceSalt, _ := ioutil.ReadFile("secret.conf")
	secretStr, _ := token.SignedString(spiceSalt)
	return secretStr
}
