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
	res := Hello{Msg: "Hi, my dear  friend!!!"}
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

	err = instance.DB.UpdateUserToken(loginUser.Nickname, loginUser.Token)

	if err != nil {
		fmt.Println("Unable to update token:", err)
		c.Status(400)
		return
	}

	fmt.Print("SINGIN ", loginUser.Nickname, " token: \n", loginUser.Token, "\n\n")

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

func (instance *App) Logout(c *gin.Context) {
	c.Status(666)
}

func (instance *App) GetAuthorPromises(c *gin.Context) {
	id, _ := c.Get("id")

	nickname, err := instance.DB.GetNicknameById(int(id.(float64)))
	if err != nil {
		fmt.Println("Getting nickname error: ", nickname)
		c.Status(404)
		return
	}
	fmt.Println(nickname)
	ps, err := instance.DB.GetPromisesByAuthor(nickname)
	if err != nil {
		fmt.Println("Getting feed error: ", err)
		c.Status(404)
		return
	}
	if len(ps) == 0 {
		c.Status(404)
		return
	}
	fmt.Println(ps)
	c.JSON(200, ps)
}

func (instance *App) GetReceiverPromises(c *gin.Context) {
	id, _ := c.Get("id")

	nickname, err := instance.DB.GetNicknameById(int(id.(float64)))
	if err != nil {
		fmt.Println("Getting nickname error: ", nickname)
		c.Status(404)
		return
	}
	fmt.Println(nickname)
	ps, err := instance.DB.GetPromisesByReceiver(nickname)
	if err != nil {
		fmt.Println("Getting feed error: ", err)
		c.Status(404)
		return
	}
	fmt.Println(ps)
	if len(ps) == 0 {
		c.Status(404)
		return
	}
	fmt.Println(ps)
	c.JSON(200, ps)
}

func (instance *App) SignUpHand(c *gin.Context) {

	newUser := models.SignUpUserStruct{}
	decoder := json.NewDecoder(c.Request.Body)
	//decoder.DisallowUnknownFields()
	err := decoder.Decode(&newUser)

	if err != nil {
		fmt.Println("Unable to decode SignUp request: ", err)
		c.Status(400)
		return
	}

	fmt.Print(newUser.Nickname, " token: \n", newUser.Token, "\n\n")

	err = instance.DB.InsertNewUser(newUser)

	if err != nil {
		fmt.Println("Unable to insert new user to database:", err)
		c.Status(409)
		return
	}

	err, id := instance.DB.CheckUserExist(newUser.Nickname)

	if err != nil {
		fmt.Println("Unable to check user existing:", err)
		c.Status(400)
		return
	}

	sessionId := instance.createSessionId(id)
	c.SetCookie("session_id", sessionId, 3600, "/p", "", false, false)
	c.Status(201)
}

func (instance *App) GetAllUsers(c *gin.Context) {
	id, _ := c.Get("id")
	users, err := instance.DB.GetUsers(int(id.(float64)), "-")
	fmt.Println(users)
	if err != nil {
		c.Status(404)
	}

	var resp []string
	for _, u := range users {
		resp = append(resp, u.Nickname)
	}
	fmt.Println(resp)
	c.JSON(200, resp)
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
