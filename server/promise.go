package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Pickausernaame/Ambitine_backend/server/models"
	"github.com/gin-gonic/gin"
)

func (instance *App) SendNotification(p models.Promise, token string) (err error) {
	data := []byte(`
		{
			"notifications": [
				{
					"tokens": ["` + token + `"],
					"platform": 2,
					"title": "` + p.Author + " promesed you that:" + `",
					"message": "` + p.Description + `"
				}
			]
		}`)

	r := bytes.NewReader(data)
	_, err = http.Post("http://35.228.98.103:8088/api/push", "application/json", r)

	return
}

func (instance *App) CreateNewPromise(c *gin.Context) {
	var (
		p models.Promise
	)

	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&p)

	if err != nil {
		fmt.Println("Unable to decode CreateNewPromise request:", err)
		c.Status(400)
		return
	}

	id, _ := c.Get("id")

	p.Author, err = instance.DB.GetNicknameById(int(id.(float64)))

	if err != nil {
		fmt.Println("Unable to get promise author nickname by id :", err)
		c.Status(400)
		return
	}

	p.AuthorImgUrl, err = instance.DB.GetImgUrlByNickname(p.Author)

	if err != nil {
		fmt.Println("Unable to get authorImgUrl by nickname :", err)
		c.Status(400)
		return
	}

	p.ReceiverImgUrl, err = instance.DB.GetImgUrlByNickname(p.Receiver)

	if err != nil {
		fmt.Println("Unable to get authorImgUrl by nickname :", err)
		c.Status(400)
		return
	}

	fmt.Println(p)

	err = instance.DB.SetNewPromise(p)

	if err != nil {
		fmt.Println("Unable to create new promise :", err)
		c.Status(400)
		return
	}

	token, err := instance.DB.GetUserToken(p.Receiver)

	fmt.Print(p.Receiver, " token: \n", token, "\n\n")

	if err != nil {
		fmt.Println("Unable to sget notifications :", err)
		c.Status(400)
		return
	}

	err = instance.SendNotification(p, token)

	if err != nil {
		fmt.Println("Unable to send notifications :", err)
		c.Status(400)
		return
	}

	c.Status(201)
}
