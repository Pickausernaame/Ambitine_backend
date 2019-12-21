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
}

type SignUpUserStruct struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInUserStruct struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type Promise struct {
	Author      string `json:"author_username"`
	Receiver    string `json:"reciever_username"`
	ImgUrl      string `json:"img_url"`
	Description string `json:"promise_description"`
	Pastdue     int64  `json:"pastdue"`
	Deposit     int    `json:"deposit"`
	Accepted    bool   `json:"accepted"`
}

type FeedPromise []Promise

//func (u *SignUpUserStruct) Validation() bool {
//	if u.Email == "" || u.Nickname == "" || u.Password == "" {
//		return false
//	}
//
//	return true
//}
