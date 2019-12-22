package server

import (
	"bytes"
	"net/http"

	"github.com/Pickausernaame/Ambitine_backend/server/models"
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
	_, err = http.Post("http://localhost:8088/api/push", "application/json", r)

	return
}

// func (instance *App) CreateNewPromise(c *gin.Context) {
// 	var (
// 		p models.Promise
// 	)

// 	decoder := json.NewDecoder(c.Request.Body)
// 	decoder.DisallowUnknownFields()
// 	err := decoder.Decode(&p)

// 	if err != nil {
// 		fmt.Println("Unable to decode CreateNewPromise request:", err)
// 		c.Status(400)
// 		return
// 	}

// 	err = instance.DB.SetNewPromise(p)

// 	if err != nil {
// 		fmt.Println("Unable to create new promise :", err)
// 		c.Status(400)
// 		return
// 	}

// 	// err = instance.SendNotification(p, token)

// 	if err != nil {
// 		fmt.Println("Unable to send notifications :", err)
// 		c.Status(400)
// 		return
// 	}

// 	c.Status(201)
// }
