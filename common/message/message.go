package message

import "chatroom/client/model"

const (
	LoginMesType ="LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
	NotyfyUserStatusMesType = "NotyfyUserStatusMes"
	SmsMesType = "SmsMes"
)

const (
	UserOnline = iota
	UserOffline
	UserBusySatus

)




type  Message struct {
	Type string `json:"type"`
	Date string `json:"date"`
}

type  LoginMes struct {
	UserId int `json:"user_id"`
	UserPwd string `json:"user_pwd"`
	UserName string `json:"user_name"`
}

type LoginResMes struct {
	Code int `json:"code"`
	UserIds []int `json:"user_ids"`
	Err string `json:"err"`
}

type RegisterMes struct {
	UserId int `json:"user_id"`
	UserPwd string `json:"user_pwd"`
	UserName string `json:"user_name"`
	UserStatus int `json:"user_status"`
}
type RegisterResMes struct {
	Code int `json:"code"`
	Err string `json:"error"`
}





type NotyfyUserStatusMes struct {
	UserId int `json:"user_id"`
	Status int `json:"status"`
}

type SmsMes struct {
	Content string `json:"content"`
	User model.User
}