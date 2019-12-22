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
		Email:    "ya_eb@gmail.com",
		Password: "bmstu",
		Token:    "cbgsAcfQr7U:APA91bG7cjJtpEJtogtrEza0uferz6qwOnC2PHZi0Sg6d9J7qCH-jJ5kWbS59p8hJ2fXTLh2FzBmcE3tVmOY-ArdmG1HohD9NMXB-qtujQlGdzuvZqclks51IhHSKAaNanFO7N3UWP0c",
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
		Nickname: "gel0",
		Email:    "lolo@ya.ru",
		Password: "qwerty",
		Token:    "cFptjzyMPD4:APA91bHQktFeKjbnX7Se0pKt5Mdf94vIarRY02ctbune2kj59Tfe1OqdbUPfcnVUGvl0iof2KcSKtDfy2l0ad8Pj4FZIGTq-RQ3MXrWzwjyy8anuXtrW2Z3QPp6-RJExs1gb4lJf2zgx",
	}
	m.createUser(oleg)

	oleg_full := models.SignUpUserStruct{
		Nickname: "OLEG_KRUTO_OZVUCHIVAET_TORGOVYU_FEDERACIU",
		//Email:    "lolo@ya.ru",
		Password: "DROIDEK",
	}
	m.createUser(oleg_full)

	p := models.Promise{
		Author:        "evv",
		Receiver:      "ntn",
		ReciverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:  "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:   "Do this feed right now",
		Pastdue:       1576417777,
		Deposit:       1000,
		Accepted:      0,
	}

	m.createPromise(p)

	p = models.Promise{
		Author:        "evv",
		Receiver:      "tim",
		ReciverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:  "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:   "Okay we try do the best",
		Pastdue:       1588417777,
		Deposit:       23000,
		Accepted:      0,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:        "gelo",
		Receiver:      "ntn",
		ReciverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:  "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:   "Lets get my honey, bro",
		Pastdue:       1688417777,
		Deposit:       230,
		Accepted:      0,
	}

	m.createPromise(p)

	p = models.Promise{
		Author:        "tim",
		Receiver:      "ntn",
		ReciverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:  "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:   "Subscribe for GGG !",
		Pastdue:       1788417777,
		Deposit:       29000,
		Accepted:      0,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:        "evv",
		Receiver:      "gel0",
		ReciverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:  "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:   "Lets do something great!",
		Pastdue:       1888417777,
		Deposit:       5300000,
		Accepted:      0,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:        "gelo",
		Receiver:      "evv",
		ReciverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:  "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:   "Lets go in mumu!",
		Pastdue:       1818417777,
		Deposit:       300232100,
		Accepted:      0,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:        "tim",
		Receiver:      "evv",
		ReciverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:  "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:   "Я, бл***, в своём познании настолько преисполнился, что я как будто бы уже 100 триллионов миллиардов лет, бл***, проживаю на триллионах и триллионах таких же планет, понимаешь? Как эта Земля. Мне уже этот мир абсолютно понятен, и я здесь ищу только одного: покоя, умиротворения и вот этой гармонии от слияния с бесконечно вечным.",
		Pastdue:       1828417777,
		Deposit:       300231000,
		Accepted:      0,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:        "ntn",
		Receiver:      "evv",
		ReciverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:  "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:   "CREATE NEW MEMES PLS",
		Pastdue:       1928417777,
		Deposit:       30230,
		Accepted:      0,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:        "ntn",
		Receiver:      "evv",
		ReciverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:  "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:   "CREATE NEW MEMES PLS",
		Pastdue:       1428417777,
		Deposit:       30034000,
		Accepted:      0,
	}

	m.createPromise(p)
}
