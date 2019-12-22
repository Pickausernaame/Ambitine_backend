package models

type UserStruct struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	ImgUrl   string `json:"imgurl"`
	Region   string `json:"region"`
	About    string `json:"about"`
	Carma    string `json:"carma"`
	Token    string `json:"token"`
}

type SignUpUserStruct struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type SignInUserStruct struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type AutoComplete struct {
	Nickname string `json:"username"`
	Img      string `json:"img_url"`
}

type Promise struct {
	Author        string `json:"author_username"`
	Receiver      string `json:"reciever_username"`
	ReciverImgUrl string `json:"reciver_img_url"`
	AuthorImgUrl  string `json:"author_img_url"`
	Description   string `json:"promise_description"`
	Pastdue       int64  `json:"pastdue"`
	Deposit       int    `json:"deposit"`
	Accepted      int    `json:"accepted"`
	Expires       string `json:"expires"`
}

// type Notification struct {
// 	Author       string `json:"author_username"`
// 	Receiver     string `json:"reciever_username"`
// 	AuthorImgUrl string `json:"author_img_url"`
// 	Token        string `json:"token"`
// 	Messege      string `json:"messege"`
// 	Title        string `json:"title"`
// 	Expires      string `json:"expires"`
// }

type FeedPromise []Promise

//func (u *SignUpUserStruct) Validation() bool {
//	if u.Email == "" || u.Nickname == "" || u.Password == "" {
//		return false
//	}
//
//	return true
//}
