package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Pickausernaame/Ambitine_backend/server/kanzler"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Pickausernaame/Ambitine_backend/server/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Hello struct {
	Msg string `json:"msg"`
}

func (instance *App) SendNotify(c *gin.Context) {
	var n models.Notify

	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&n)

	if err != nil {
		fmt.Println("Unable to decode Notify request:", err)
		c.Status(400)
		return
	}

	token, err := instance.DB.GetUserToken(n.Nickname)

	if err != nil {
		fmt.Println("Unable to get user tocken:", err)
		c.Status(400)
		return
	}

	notifyBody := `{
		"notifications": [
			{
				"tokens": ["` + token + `"],
				"platform": 2,
				"title": "` + n.Title + `",
				"message": "` + n.Messege + `"
			}
		]
	}`

	data := []byte(notifyBody)

	r := bytes.NewReader(data)
	_, err = http.Post("http://35.228.98.103:8088/api/push", "application/json", r)
}

func HelloFunc(c *gin.Context) {
	res := Hello{Msg: "Hi, my dear  friend!!!"}
	c.JSON(200, res)
}

func (instance *App) UploadImg(c *gin.Context) {
	err := c.Request.ParseMultipartForm(32 << 20)

	if err != nil {
		fmt.Println("Upload error: ", err)
		c.Status(400)
		return
	}

	file, _, err := c.Request.FormFile("avatar")
	if err != nil {
		fmt.Println("Upload error: ", err)
		c.Status(400)
		return
	}

	defer file.Close()
	id, _ := c.Get("id")

	imgFileName := strconv.Itoa(int(id.(float64))) + strconv.Itoa(int(time.Now().Unix()))

	picpath := "./static/avatars/img" + imgFileName + ".jpeg"
	f, err := os.OpenFile(picpath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Upload error: ", err)
		c.Status(400)
		return
	}

	ImgUrl := "http://35.228.98.103:9090/avatars/img" + imgFileName + ".jpeg"

	err = instance.DB.UpdateUserImgUrl(int(id.(float64)), ImgUrl)
	if err != nil {
		fmt.Println("Upload error: ", err)
		c.Status(400)
		return
	}

	c.Status(200)
	defer f.Close()
	io.Copy(f, file)
}

func (instance *App) GetUserBalance(c *gin.Context) {
	id, _ := c.Get("id")
	addr, err := instance.DB.GetAddressById(int(id.(float64)))

	if err != nil {
		fmt.Println("Unable to get balance by id:", err)
		c.Status(400)
		return
	}

	balance, _, _ := instance.WM.CheckBalance(addr)

	ethBalance, _ := balance.Float64()
	usdBalance := kanzler.EtherPerUsd() * ethBalance

	c.JSON(200, gin.H{"balance": usdBalance})
}

func (instance *App) SignInHand(c *gin.Context) {
	var (
		loginUser models.SignInUserStruct
	)

	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&loginUser)

	fmt.Println("Login:\n", loginUser)

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
	id, _ := c.Get("id")

	_, err := instance.DB.GetNicknameById(int(id.(float64)))

	if err != nil {
		fmt.Println("User does not exist. Id: ", id)
		c.Status(400)
		return
	}

	err = instance.DB.RemoveTockenById(int(id.(float64)))

	if err != nil {
		fmt.Println("Unable to remove device token form db: ", err)
		c.Status(400)
		return
	}

	c.Status(200)
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
		fmt.Println("Can't find any promises: ", err)
		c.Status(404)
		return
	}
	fmt.Println(ps)
	dept := 0.
	for _, p := range ps {
		if p.Accepted == 0 {
			dept = dept + p.Deposit
		}
	}
	fmt.Println(dept)
	err = instance.DB.UpdateDeptById(int(id.(float64)), dept)
	if err != nil {
		fmt.Println("Updating debt by id err: ", err)
		c.Status(405)
		return
	}
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
		fmt.Println("Can't find any promise in db: ", err)
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

	fmt.Println("Signup:\n", newUser)

	privateKey, address, err := instance.WM.CreateWallet()
	if err != nil {
		fmt.Println("Unable to create new wallet:", err)
		c.Status(409)
		return
	}

	err = instance.DB.InsertNewUser(newUser, privateKey, address)

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
	c.SetCookie("session_id", sessionId, 3600, "/", "", false, false)
	c.Status(201)
}

func (instance *App) UserInfo(c *gin.Context) {
	id, _ := c.Get("id")
	u, err := instance.DB.GetUserInfo(int(id.(float64)))
	if err != nil {
		fmt.Println("Getting user's info error: ", err)
		c.Status(409)
		return
	}

	balance, _, err := instance.WM.CheckBalance(u.Wallet)
	if err != nil {
		fmt.Println("Getting user's balance error: ", err)
		c.Status(409)
		return
	}

	u.Balance, _ = balance.Float64()
	u.Balance = u.Balance*kanzler.EtherPerUsd() - u.Debt
	fmt.Println(u)
	c.JSON(200, u)
	return
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
