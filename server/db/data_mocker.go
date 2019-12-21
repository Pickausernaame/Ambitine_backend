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
		Deposit:     5300000,
		Accepted:    false,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:      "gelo",
		Receiver:    "evv",
		ImgUrl:      "https://d.newsweek.com/en/full/1176971/obesity-meme.png",
		Description: "Lets go in mumu!",
		Pastdue:     1818417777,
		Deposit:     300232100,
		Accepted:    false,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:      "tim",
		Receiver:    "evv",
		ImgUrl:      "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description: "Я, бл***, в своём познании настолько преисполнился, что я как будто бы уже 100 триллионов миллиардов лет, бл***, проживаю на триллионах и триллионах таких же планет, понимаешь? Как эта Земля. Мне уже этот мир абсолютно понятен, и я здесь ищу только одного: покоя, умиротворения и вот этой гармонии от слияния с бесконечно вечным.",
		Pastdue:     1828417777,
		Deposit:     300231000,
		Accepted:    false,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:      "ntn",
		Receiver:    "evv",
		ImgUrl:      "https://bostonglobe-prod.cdn.arcpublishing.com/resizer/RHDkOXAijJlc7rv8Owlk19kcyt4=/1440x0/arc-anglerfish-arc2-prod-bostonglobe.s3.amazonaws.com/public/Y5GUIDYVWJGVFP5MWNYJR5375I.png",
		Description: "CREATE NEW MEMES PLS",
		Pastdue:     1928417777,
		Deposit:     30230,
		Accepted:    false,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:      "ntn",
		Receiver:    "evv",
		ImgUrl:      "https://memepedia.ru/wp-content/uploads/2017/08/%D1%81%D0%BC%D0%B5%D1%85-%D0%B4%D0%B6%D0%B5%D0%B9%D0%BC%D1%81%D0%BE%D0%BD%D0%B0.jpg",
		Description: "CREATE NEW MEMES PLS",
		Pastdue:     1428417777,
		Deposit:     30034000,
		Accepted:    false,
	}

	m.createPromise(p)
}
