package models

type SignupReq struct {
	User       string `json:"user"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}
