package db

import (
	"os"

	"github.com/Pickausernaame/Ambitine_backend/server/kanzler"

	"github.com/Pickausernaame/Ambitine_backend/server/models"
)

const (
	REAL_PRIVATE1 = "d1d3c8e1e89e38553d745c5e3eaca431a70bcf368f3b41697850cf26c6acb199"
	REAL_ADRESS1  = "0x0d43bB78be45cD511f110C40371d0826f8494aF9"

	REAL_PRIVATE2 = "459edb08fde808b9da7b3df79974e59544e7f072f997ab4637eba3d763cd3619"
	REAL_ADRESS2  = "0x1c2f77a9678A72613ad61c9867ec1E0fA4b61999"

	REAL_PRIVATE3 = "6ec8542e3e6226ca6b2256a7c8d076616415aae6f78086e32911fbc59060ab88"
	REAL_ADRESS3  = "0x5738E0746B515CA27Ad9a24c60F72661660d33Bc"

	REAL_PRIVATE4 = "4fd9d12e5e85a4351559d08d45060a527a03de613de5e419bfe23b8e3f4dbdf2"
	REAL_ADRESS4  = "0xBbFf721bE0d602A84414a22F4eBB886b361BA3C3"

	REAL_PRIVATE5 = "3d6ae8b13b180abe5f524a8b7b0d31e3c41cb6f06b4d07bab3989058ea7c8700"
	REAL_ADRESS5  = "0x3358b8E5152BF2Bcb11D82ccd61c01455BB00c32"
)

type Mocker struct {
	DB *DBHandler
	WM *kanzler.WalletManager
}

func (m *Mocker) createUser(u models.SignUpUserStruct, privateKey string, address string) {
	m.DB.InsertNewUser(u, privateKey, address)
	//fmt.Println("user: ", err)
}

func (m *Mocker) createPromise(p models.Promise) {
	m.DB.SetNewPromise(p)
	//fmt.Println("promise: ", )
}

func (m *Mocker) setUserAvatar(n string, u string) {
	m.DB.SetUserImgUrl(n, u)
	//fmt.Println("imgurl: ", err)
}

func (m *Mocker) Mock() {
	state, exist := os.LookupEnv("STATE")
	if !exist {
		state = "debug"
	}
	Vladimir := models.SignUpUserStruct{
		Nickname: "Vladimir",
		Email:    "ya_eb@gmail.com",
		Password: "bmstu",
		Token:    "cbgsAcfQr7U:APA91bG7cjJtpEJtogtrEza0uferz6qwOnC2PHZi0Sg6d9J7qCH-jJ5kWbS59p8hJ2fXTLh2FzBmcE3tVmOY-ArdmG1HohD9NMXB-qtujQlGdzuvZqclks51IhHSKAaNanFO7N3UWP0c",
	}
	privateKey := ""
	address := ""
	if state == "debug" {
		privateKey = "2ebf2ba43c108f76b387a522ab18f7dfd5e3c1daacedefa89e0b83bfc2db5015"
		address = "0x67057856B8527Af81Ef3802e64eFEB1a97C14D30"
	} else if state == "prod" {
		privateKey = REAL_PRIVATE1
		address = REAL_ADRESS1
	}

	m.createUser(Vladimir, privateKey, address)

	Antony := models.SignUpUserStruct{
		Nickname: "Antony",
		Email:    "ebaboba@gmail.com",
		Password: "1488",
	}
	if state == "debug" {
		privateKey = "d14cade251eeddee89f7bd24a56f5fc2d58ad791b456eac599bcb5798cdd5fce"
		address = "0xDeA087aFdd4aE37902f626EAd264eE982D78Dc6a"
	} else if state == "prod" {
		privateKey = REAL_PRIVATE2
		address = REAL_ADRESS2
	}

	m.createUser(Antony, privateKey, address)

	Temirlan := models.SignUpUserStruct{
		Nickname: "Temirlan",
		Email:    "kazik@mail.ru",
		Password: "urus",
	}
	if state == "debug" {
		privateKey = "f38057c879e9aedc33d5823c15ab2640496afed059f3781ef708df4577e945b2"
		address = "0x7F51DCbdBdb4BB0A8a10387B5Fc3A9405F47a03f"
	} else if state == "prod" {
		privateKey = REAL_PRIVATE3
		address = REAL_ADRESS3
	}
	m.createUser(Temirlan, privateKey, address)

	Oleg := models.SignUpUserStruct{
		Nickname: "Oleg",
		Email:    "lolo@ya.ru",
		Password: "1234",
		Token:    "cFptjzyMPD4:APA91bHQktFeKjbnX7Se0pKt5Mdf94vIarRY02ctbune2kj59Tfe1OqdbUPfcnVUGvl0iof2KcSKtDfy2l0ad8Pj4FZIGTq-RQ3MXrWzwjyy8anuXtrW2Z3QPp6-RJExs1gb4lJf2zgx",
	}
	if state == "debug" {
		privateKey = "630e0cee6c70243c85655ff39bf8ba2822356df78f223db374589a5a53f265eb"
		address = "0x34f2361235dFa60d20571cC059Ecf53ed02AA05e"
	} else if state == "prod" {
		privateKey = REAL_PRIVATE4
		address = REAL_ADRESS4
	}
	m.createUser(Oleg, privateKey, address)

	oleg_full := models.SignUpUserStruct{
		Nickname: "OLEG_KRUTO_OZVUCHIVAET_TORGOVYU_FEDERACIU",
		Email:    "lolo@ya.ru",
		Password: "DROIDEK",
	}
	if state == "debug" {
		privateKey = "f7ea4e300c90742decc44f65b1d7a5a7308ed799f08b6140d69e0bcd9d448b29"
		address = "0x668cbE895A9A2C24421530d919d2B34B337272fB"
	} else if state == "prod" {
		privateKey = REAL_PRIVATE5
		address = REAL_ADRESS5
	}
	m.createUser(oleg_full, privateKey, address)

	m.setUserAvatar("OLEG_KRUTO_OZVUCHIVAET_TORGOVYU_FEDERACIU", "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg")
	m.setUserAvatar("Vladimir", "http://35.228.98.103:9090/avatars/evv.png")
	m.setUserAvatar("Antony", "http://35.228.98.103:9090/avatars/ntn.png")
	m.setUserAvatar("Temirlan", "http://35.228.98.103:9090/avatars/tim.png")
	m.setUserAvatar("Oleg", "http://35.228.98.103:9090/avatars/olg.png")

	// m.setUserAvatar("OLEG_KRUTO_OZVUCHIVAET_TORGOVYU_FEDERACIU", "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg")
	// m.setUserAvatar("Vladimir", "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg")
	// m.setUserAvatar("Antony", "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg")
	// m.setUserAvatar("Temirlan", "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg")
	// m.setUserAvatar("Oleg", "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg")
	// m.setUserAvatar("Oleg", "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg")

	p := models.Promise{
		Author:      "Vladimir",
		Receiver:    "Antony",
		Description: "Сделать фид обещаний прямо сейчас",
		Pastdue:     1766517792000,
		Deposit:     12,
		Accepted:    0,
	}

	m.createPromise(p)

	p = models.Promise{
		Author:      "Vladimir",
		Receiver:    "Temirlan",
		Description: "Настроить апи кошелька, сделать хэндлер на адрес",
		Pastdue:     1766517792000,
		Deposit:     23,
		Accepted:    0,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:      "Oleg",
		Receiver:    "Antony",
		Description: "Настроить работу уведомлений в фоне",
		Pastdue:     1766517792000,
		Deposit:     23,
		Accepted:    0,
	}

	m.createPromise(p)

	p = models.Promise{
		Author:      "Temirlan",
		Receiver:    "Antony",
		Description: "Показать танец живота, дать двоечку мамбетам",
		Pastdue:     1766517792000,
		Deposit:     29,
		Accepted:    0,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:      "Vladimir",
		Receiver:    "Oleg",
		Description: "Исправить орфографические ошибки в JSON запросах на бэк",
		Pastdue:     1766517792000,
		Deposit:     5.3,
		Accepted:    0,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:      "Oleg",
		Receiver:    "Vladimir",
		Description: "Исправить орфографическте ошибки, изменить поля модели, написать хэндлер нотификаций",
		Pastdue:     1766517792000,
		Deposit:     30,
		Accepted:    0,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:      "Vladimir",
		Receiver:    "Oleg",
		Description: "Посмотреть список любимых фильмов и чекнуть их оценки на кинопоиске",
		Pastdue:     1766517792000,
		Deposit:     3,
		Accepted:    0,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:      "Antony",
		Receiver:    "Vladimir",
		Description: "Выбрать, что заказать им и что сказать нам",
		Pastdue:     1766517792000,
		Deposit:     3.2,
		Accepted:    0,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:      "Antony",
		Receiver:    "Vladimir",
		Description: "Стать самым свежим в школе, у",
		Pastdue:     1766517792000,
		Deposit:     3.9,
		Accepted:    0,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:      "Vladimir",
		Receiver:    "Temirlan",
		Description: "Выучить наизусть все песни моргенштерна",
		Pastdue:     1766517792000,
		Deposit:     3.1,
		Accepted:    -1,
	}
	m.createPromise(p)
	p = models.Promise{
		Author:      "Vladimir",
		Receiver:    "Oleg",
		Description: "Не приставать больше ночью и не снимать сторис в инсту",
		Pastdue:     1766517792000,
		Deposit:     2,
		Accepted:    -1,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:      "Vladimir",
		Receiver:    "Antony",
		Description: "Хорошо провести эти выходные",
		Pastdue:     1766517792000,
		Deposit:     17,
		Accepted:    1,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:      "Vladimir",
		Receiver:    "Temirlan",
		Description: "Сконнектиться на счет сета кук",
		Pastdue:     1766517792000,
		Deposit:     25,
		Accepted:    1,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:      "Antony",
		Receiver:    "Oleg",
		Description: "Сдать курсач до нового года",
		Pastdue:     1766517792000,
		Deposit:     44,
		Accepted:    1,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:      "Antony",
		Receiver:    "Vladimir",
		Description: "Провести бой с тенью или игру с самим собой",
		Pastdue:     1766517792000,
		Deposit:     150,
		Accepted:    -1,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:      "Temirlan",
		Receiver:    "Vladimir",
		Description: "Больше не кусать во время борьбы за руку (кусать животом можно)",
		Pastdue:     1766517792000,
		Deposit:     170,
		Accepted:    1,
	}
	m.createPromise(p)

	p = models.Promise{
		Author:      "Oleg",
		Receiver:    "Vladimir",
		Description: "Написать нормальные моки для юзеров и промисовы",
		Pastdue:     1766517792000,
		Deposit:     26,
		Accepted:    1,
	}
	m.createPromise(p)

}
