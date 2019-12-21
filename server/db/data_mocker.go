package db

import (
	"fmt"
	"github.com/Pickausernaame/Ambitine_backend/server/models"
)

type Mocker struct {
	DB *DBHandler
}

func (m *Mocker) createUser(u models.SignUpUserStruct) {
	fmt.Println(m.DB.InsertNewUser(u))
}

func (m *Mocker) createPromise(p models.Promise) {
	fmt.Println(m.DB.SetNewPromise(p))
}

func (m *Mocker) Mock() {
	evv := models.SignUpUserStruct{
		Nickname: "evv",
		//Email:    "ya_eb@gmail.com",
		Password: "bmstu",
	}
	m.createUser(evv)

	ntn := models.SignUpUserStruct{
		Nickname: "ntn",
		//Email:    "ebaboba@gmail.com",
		Password: "1488",
	}
	m.createUser(ntn)

	tim := models.SignUpUserStruct{
		Nickname: "tim",
		//Email:    "kazik@mail.ru",
		Password: "urus",
	}
	m.createUser(tim)

	oleg := models.SignUpUserStruct{
		Nickname: "gelo",
		//Email:    "lolo@ya.ru",
		Password: "qwerty",
	}
	m.createUser(oleg)

	p := models.Promise{
		Author:      "evv",
		Receiver:    "ntn",
		ImgUrl:      "https://cdn5.vectorstock.com/i/1000x1000/51/99/icon-of-user-avatar-for-web-site-or-mobile-app-vector-3125199.jpg",
		Description: "Do this feed right now",
		Pastdue:     1576417777,
		Deposit:     1000,
		Accepted:    false,
	}

	m.createPromise(p)

	p = models.Promise{
		Author:      "evv",
		Receiver:    "tim",
		ImgUrl:      "https://cdn5.vectorstock.com/i/1000x1000/51/99/icon-of-user-avatar-for-web-site-or-mobile-app-vector-3125199.jpg",
		Description: "Okay we try do the best",
		Pastdue:     1588417777,
		Deposit:     23000,
		Accepted:    false,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:      "gelo",
		Receiver:    "ntn",
		ImgUrl:      "https://www.freelogodesign.org/Content/img/logo-samples/sophia.png",
		Description: "Lets get my honey, bro",
		Pastdue:     1688417777,
		Deposit:     230,
		Accepted:    false,
	}

	m.createPromise(p)

	p = models.Promise{
		Author:      "tim",
		Receiver:    "ntn",
		ImgUrl:      "https://media.kidozi.com/unsafe/600x600/img.kidozi.com/design/600/600/0a0909/74062/c6ba04a076242d3249b240eae6461f09.png.jpg",
		Description: "Subscribe for GGG !",
		Pastdue:     1788417777,
		Deposit:     29000,
		Accepted:    false,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:      "evv",
		Receiver:    "gelo",
		ImgUrl:      "https://cdn-prod.medicalnewstoday.com/content/images/articles/234/234239/cocaine-drug.jpg",
		Description: "Lets do something great!",
		Pastdue:     1888417777,
		Deposit:     300000,
		Accepted:    false,
	}
	m.createPromise(p)
}
