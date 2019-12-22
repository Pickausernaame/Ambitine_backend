package db

import (
	"fmt"

	"github.com/Pickausernaame/Ambitine_backend/server/kanzler"

	"github.com/Pickausernaame/Ambitine_backend/server/models"
)

type Mocker struct {
	DB *DBHandler
	WM *kanzler.WalletManager
}

func (m *Mocker) createUser(u models.SignUpUserStruct, privateKey string, address string) {
	//privateKey, address, err := m.WM.CreateWallet()
	//if err != nil {
	//	fmt.Println("Unable to create new wallet:", err)
	//	return
	//}
	fmt.Println("user: ", m.DB.InsertNewUser(u, privateKey, address))
}

func (m *Mocker) createPromise(p models.Promise) {
	fmt.Println("promise: ", m.DB.SetNewPromise(p))
}

func (m *Mocker) setUserAvatar(n string, u string) {
	fmt.Println("imgurl: ", m.DB.SetUserImgUrl(n, u))
}

func (m *Mocker) Mock() {

	evv := models.SignUpUserStruct{
		Nickname: "evv",
		Email:    "ya_eb@gmail.com",
		Password: "bmstu",
		Token:    "cbgsAcfQr7U:APA91bG7cjJtpEJtogtrEza0uferz6qwOnC2PHZi0Sg6d9J7qCH-jJ5kWbS59p8hJ2fXTLh2FzBmcE3tVmOY-ArdmG1HohD9NMXB-qtujQlGdzuvZqclks51IhHSKAaNanFO7N3UWP0c",
	}
	privateKey := "2ebf2ba43c108f76b387a522ab18f7dfd5e3c1daacedefa89e0b83bfc2db5015"
	address := "0x67057856B8527Af81Ef3802e64eFEB1a97C14D30"
	m.createUser(evv, privateKey, address)

	ntn := models.SignUpUserStruct{
		Nickname: "ntn",
		Email:    "ebaboba@gmail.com",
		Password: "1488",
	}
	privateKey = "d14cade251eeddee89f7bd24a56f5fc2d58ad791b456eac599bcb5798cdd5fce"
	address = "0xDeA087aFdd4aE37902f626EAd264eE982D78Dc6a"
	m.createUser(ntn, privateKey, address)

	tim := models.SignUpUserStruct{
		Nickname: "tim",
		Email:    "kazik@mail.ru",
		Password: "urus",
	}
	privateKey = "f38057c879e9aedc33d5823c15ab2640496afed059f3781ef708df4577e945b2"
	address = "0x7F51DCbdBdb4BB0A8a10387B5Fc3A9405F47a03f"
	m.createUser(tim, privateKey, address)

	oleg := models.SignUpUserStruct{
		Nickname: "gel0",
		Email:    "lolo@ya.ru",
		Password: "qwerty",
		Token:    "cFptjzyMPD4:APA91bHQktFeKjbnX7Se0pKt5Mdf94vIarRY02ctbune2kj59Tfe1OqdbUPfcnVUGvl0iof2KcSKtDfy2l0ad8Pj4FZIGTq-RQ3MXrWzwjyy8anuXtrW2Z3QPp6-RJExs1gb4lJf2zgx",
	}
	privateKey = "630e0cee6c70243c85655ff39bf8ba2822356df78f223db374589a5a53f265eb"
	address = "0x34f2361235dFa60d20571cC059Ecf53ed02AA05e"
	m.createUser(oleg, privateKey, address)

	oleg_full := models.SignUpUserStruct{
		Nickname: "OLEG_KRUTO_OZVUCHIVAET_TORGOVYU_FEDERACIU",
		Email:    "lolo@ya.ru",
		Password: "DROIDEK",
	}
	privateKey = "f7ea4e300c90742decc44f65b1d7a5a7308ed799f08b6140d69e0bcd9d448b29"
	address = "0x668cbE895A9A2C24421530d919d2B34B337272fB"
	m.createUser(oleg_full, privateKey, address)

	m.setUserAvatar("OLEG_KRUTO_OZVUCHIVAET_TORGOVYU_FEDERACIU", "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg")
	m.setUserAvatar("evv", "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg")
	m.setUserAvatar("ntn", "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg")
	m.setUserAvatar("tim", "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg")
	m.setUserAvatar("oleg", "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg")

	p := models.Promise{
		Author:         "evv",
		Receiver:       "ntn",
		ReceiverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:   "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:    "Do this feed right now",
		Pastdue:        1576417777,
		Deposit:        12,
		Accepted:       0,
	}

	m.createPromise(p)

	p = models.Promise{
		Author:         "evv",
		Receiver:       "tim",
		ReceiverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:   "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:    "Okay we try do the best",
		Pastdue:        1588417777,
		Deposit:        23,
		Accepted:       0,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:         "gelo",
		Receiver:       "ntn",
		ReceiverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:   "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:    "Lets get my honey, bro",
		Pastdue:        1688417777,
		Deposit:        23,
		Accepted:       0,
	}

	m.createPromise(p)

	p = models.Promise{
		Author:         "tim",
		Receiver:       "ntn",
		ReceiverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:   "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:    "Subscribe for GGG !",
		Pastdue:        1788417777,
		Deposit:        29,
		Accepted:       0,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:         "evv",
		Receiver:       "gel0",
		ReceiverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:   "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:    "Lets do something great!",
		Pastdue:        1888417777,
		Deposit:        5.3,
		Accepted:       0,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:         "gelo",
		Receiver:       "evv",
		ReceiverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:   "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:    "Lets go in mumu!",
		Pastdue:        1818417777,
		Deposit:        30,
		Accepted:       0,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:         "tim",
		Receiver:       "evv",
		ReceiverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:   "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:    "Я, бл***, в своём познании настолько преисполнился, что я как будто бы уже 100 триллионов миллиардов лет, бл***, проживаю на триллионах и триллионах таких же планет, понимаешь? Как эта Земля. Мне уже этот мир абсолютно понятен, и я здесь ищу только одного: покоя, умиротворения и вот этой гармонии от слияния с бесконечно вечным.",
		Pastdue:        1828417777,
		Deposit:        3,
		Accepted:       0,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:         "ntn",
		Receiver:       "evv",
		ReceiverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:   "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:    "CREATE NEW MEMES PLS",
		Pastdue:        1928417777,
		Deposit:        3.2,
		Accepted:       0,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:         "ntn",
		Receiver:       "evv",
		ReceiverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:   "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:    "CREATE NEW MEMES PLS",
		Pastdue:        1428417777,
		Deposit:        3.9,
		Accepted:       0,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:         "evv",
		Receiver:       "tim",
		ReceiverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:   "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:    "FALLING DOWN",
		Pastdue:        1428417777,
		Deposit:        3.1,
		Accepted:       -1,
	}
	m.createPromise(p)
	p = models.Promise{
		Author:         "evv",
		Receiver:       "gelo",
		ReceiverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:   "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:    "GDE MOI GENGIIII",
		Pastdue:        1528417777,
		Deposit:        30034000,
		Accepted:       -1,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:         "evv",
		Receiver:       "ntn",
		ReceiverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:   "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:    "Cool weekends",
		Pastdue:        1438417777,
		Deposit:        30034000,
		Accepted:       1,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:         "evv",
		Receiver:       "tim",
		ReceiverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:   "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:    "GDE MOI SHPAK",
		Pastdue:        1638417777,
		Deposit:        30034000,
		Accepted:       1,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:         "ntn",
		Receiver:       "evv",
		ReceiverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:   "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:    "GO V VORONEZ",
		Pastdue:        1618417777,
		Deposit:        30034000,
		Accepted:       1,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:         "ntn",
		Receiver:       "evv",
		ReceiverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:   "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:    "GO V MINSK",
		Pastdue:        1318417777,
		Deposit:        30034000,
		Accepted:       -1,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:         "tim",
		Receiver:       "evv",
		ReceiverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:   "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:    "POBOREMSYA",
		Pastdue:        1378417777,
		Deposit:        30034000,
		Accepted:       1,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:         "tim",
		Receiver:       "evv",
		ReceiverImgUrl: "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		AuthorImgUrl:   "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg",
		Description:    "POEBEMSYA",
		Pastdue:        1478417777,
		Deposit:        30034000,
		Accepted:       -1,
	}
	m.createPromise(p)

}
