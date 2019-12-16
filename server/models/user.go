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

func (u *SignUpUserStruct) Validation() bool {
	if u.Email == "" || u.Nickname == "" || u.Password == "" {
		return false
	}

	return true
}
