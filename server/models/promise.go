package models

type Promise struct {
	Author         string  `json:"author_username"`
	Receiver       string  `json:"receiver_username"`
	AuthorImgUrl   string  `json:"author_img_url"`
	ReceiverImgUrl string  `json:"receiver_img_url"`
	Description    string  `json:"promise_description"`
	Pastdue        int64   `json:"pastdue"`
	Deposit        float64 `json:"deposit"`
	Accepted       int     `json:"accepted"`
}
