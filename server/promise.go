package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Pickausernaame/Ambitine_backend/server/kanzler"
	"github.com/Pickausernaame/Ambitine_backend/server/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// func (instance *App) SendNotify(text string) (err error) {
// 	data := []byte(text)
// 	r := bytes.NewReader(data)
// 	_, err = http.Post("http://35.228.98.103:8088/api/push", "application/json", r)

// 	return
// }

func (instance *App) SendNotification(notification string, p models.Promise, token string) (err error) {
	data := []byte(notification)

	fmt.Println("Notification debug:\n", p, "\n", token)

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

	if p.Author == p.Receiver {
		fmt.Println("Similar promise author and receiver")
		c.Status(409)
		return
	}

	if p.Receiver == "" {
		fmt.Println("Promise receiver is empty")
		c.Status(409)
		return
	}

	id, _ := c.Get("id")

	p.Author, err = instance.DB.GetNicknameById(int(id.(float64)))

	if err != nil {
		fmt.Println("Unable to get promise author nickname by id :", err)
		c.Status(400)
		return
	}

	err, _ = instance.DB.CheckUserExist(p.Receiver)

	if err != nil {
		fmt.Println("Unable to find receive by nickname :", err)
		c.Status(400)
		return
	}

	addr, err := instance.DB.GetAddressById(int(id.(float64)))

	if err != nil {
		fmt.Println("Unable to get wallet adress by id :", err)
		c.Status(400)
		return
	}

	_, floatBalance, err := instance.WM.CheckBalance(addr)

	balance, _ := floatBalance.Float64()
	balance = balance * kanzler.EtherPerUsd()
	dept, err := instance.DB.GetDebtById(int(id.(float64)))

	balance = balance - dept

	fmt.Println("Promise deposite: ", p.Deposit)
	fmt.Println("user balance: ", balance)

	if err != nil || p.Deposit > balance || p.Deposit <= 0.0 {
		fmt.Println("Deposit is: ", p.Deposit, "\nBalance is: ", balance, "\n\n")
		fmt.Println("Unable to get balance by wallet addres, or user set wrong balance:", err)
		c.Status(401)
		return
	}

	fmt.Println(p)

	err = instance.DB.SetNewPromise(p)

	if err != nil {
		fmt.Println("Unable to create new promise :", err)
		c.Status(400)
		return
	}
	dept += p.Deposit + dept
	err = instance.DB.UpdateDeptById(int(id.(float64)), dept)

	token, err := instance.DB.GetUserToken(p.Receiver)
	fmt.Print(p.Receiver, " token: \n", token, "\n\n")

	if err != nil {
		fmt.Println("Unable to sget notifications :", err)
		c.Status(400)
		return
	}

	err = instance.SendNotification(
		`{
			"notifications": [
				{
					"tokens": ["`+token+`"],
					"platform": 2,
					"title": "`+p.Author+" promesed you that:"+`",
					"message": "`+p.Description+`"
				}
			]
		}`,
		p, token)

	if err != nil {
		fmt.Println("Unable to send notifications :", err)
		c.Status(400)
		return
	}

	c.Status(201)
}

func (instance *App) Solution(c *gin.Context) {
	sol := models.Solution{}

	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&sol)

	if err != nil {
		fmt.Println("Unable to decode Solution request:", err)
		c.Status(400)
		return
	}

	//todo Проверить что ты автор
	id, _ := c.Get("id")

	nickname, err := instance.DB.GetNicknameById(int(id.(float64)))
	if err != nil {
		fmt.Println("Unable to get nickname by id from cookie:", err)
		c.Status(400)
		return
	}

	exist, err := instance.DB.IsUserReceiverOfPromise(nickname, sol.Promise_id)
	if !exist || err != nil {
		if err != nil {
			fmt.Println("Unable to check exist:", err)
		} else {
			fmt.Println("This user is not receiver of promise")
		}
		c.Status(400)
		return
	}

	isAccepted, err := instance.DB.IsPromiseAccepted(sol.Promise_id)
	if isAccepted {
		fmt.Println("Promise is already accepted:", err)
		c.Status(400)
		return
	}

	if sol.Accepted == 1 {
		p, err := instance.DB.UpdatePromiseStatus(sol)
		if err != nil {
			fmt.Println("Unable to change promise status:", err)
			c.Status(400)
			return
		}

		token, err := instance.DB.GetUserToken(p.Author)

		if err != nil {
			fmt.Println("Unable send get user token:", err)
			c.Status(400)
			return
		}

		err = instance.SendNotification(`
		{
			"notifications": [
				{
					"tokens": ["`+token+`"],
					"platform": 2,
					"title": "Your promis was accepted!",
					"message": "`+`Congratulations!`+p.Receiver+` was accepted your promise `+p.Description+`.\n"
				}
			]
		}`,
			p, token)

		if err != nil {
			fmt.Println("Unable send notofication:", err)
			c.Status(400)
			return
		}

		c.Status(200)
		return
	} else if sol.Accepted == -1 {

		p, err := instance.DB.GetPromisesById(sol.Promise_id)
		if err != nil {
			fmt.Println("Unable to get promise by ID:", err)
			c.Status(400)
			return
		}
		from, err := instance.DB.GetPrivateByNickname(p.Author)
		if err != nil {
			fmt.Println("Unable to get private key by nickname:", err)
			c.Status(400)
			return
		}
		fromWallet, err := instance.DB.GetAddressByNickname(p.Author)
		if err != nil {
			fmt.Println("Unable to get address of wallet by nickname:", err)
			c.Status(400)
			return
		}
		to, err := instance.DB.GetAddressByNickname(p.Receiver)
		if err != nil {
			fmt.Println("Unable to get address of wallet by nickname:", err)
			c.Status(400)
			return
		}

		_, balance, err := instance.WM.CheckBalance(fromWallet)
		if err != nil {
			fmt.Println("Unable to get Balance:", err)
			c.Status(400)
			return
		}
		bal, _ := balance.Float64()
		bal = bal * kanzler.EtherPerUsd()

		if bal < p.Deposit {
			c.Status(408)
			return
		} else {
			err = instance.WM.MakeTransaction(from, to, p.Deposit/kanzler.EtherPerUsd())
			if err != nil {
				fmt.Println("Unable to make transaction:", err)
				c.Status(400)
				return
			}
		}
		_, err = instance.DB.UpdatePromiseStatus(sol)
		if err != nil {
			fmt.Println("Unable to change promise status:", err)
			c.Status(400)
			return
		}

		token, err := instance.DB.GetUserToken(p.Author)

		if err != nil {
			fmt.Println("Unable send get user token:", err)
			c.Status(400)
			return
		}

		err = instance.SendNotification(`
		{
			"notifications": [
				{
					"tokens": ["`+token+`"],
					"platform": 2,
					"title": "Your promis was declined.",
					"message": "`+p.Receiver+` was declined your promise `+p.Description+`.\n"
				}
			]
		}`,
			p, token)

		c.Status(200)
		return
	}

}
