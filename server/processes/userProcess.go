package processes

import (
	"chatroom/common/message"
	"chatroom/server/model"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	UserId int
}


func (this *UserProcess)ServerProcessLogin(mes *message.Message) (err error) {
	var loginmes message.LoginMes
	err = json.Unmarshal([]byte(mes.Date), &loginmes)
	if err != nil {
		fmt.Println("json.unmarshal err = ", err)
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes  message.LoginResMes

	user, err := model.MyUserDao.Login(loginmes.UserId, loginmes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			loginResMes.Code = 300
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 400
		} else if err == model.ERROR_USER_NOTXEISTS {
			loginResMes.Code = 500
		}
		loginResMes.Err = err.Error()


	} else {

		loginResMes.Code = 200
		this.UserId = loginmes.UserId
		userMgr.AddOnlineUser(this)

		this.NotifyOthersOnlineUser(loginmes.UserId)
		for id,_ := range userMgr.OnlineUser {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}

		
		fmt.Printf("用户消息user%v", user)
	}
	//

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.marshal faile err=", err)
		return
	}
	resMes.Date = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.marshal err=" , err)
		return
	}
	tr := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tr.WritePkg(data)
	if err != nil {
		fmt.Println("服务器发送数据失败")
	}

	return
}


func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {

	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Date), &registerMes)
	if err != nil {
		fmt.Println("json.unmarshal err = ", err)
	}

	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	var regsiterResMes  message.RegisterResMes

	user, err := model.MyUserDao.Register(registerMes.UserId, registerMes.UserPwd, registerMes.UserName)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			regsiterResMes.Code = 300
		} else if err == model.ERROR_USER_PWD {
			regsiterResMes.Code = 400
		} else if err == model.ERROR_USER_NOTXEISTS {
			regsiterResMes.Code = 500
		}
		regsiterResMes.Err = err.Error()


	} else {

		regsiterResMes.Code = 200
		fmt.Printf("用户消息user%v", user)
	}
	//

	data, err := json.Marshal(regsiterResMes)
	if err != nil {
		fmt.Println("json.marshal faile err=", err)
		return
	}
	resMes.Date = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.marshal err=" , err)
		return
	}
	tr := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tr.WritePkg(data)
	if err != nil {
		fmt.Println("服务器发送数据失败")
	}

	return
}

func (this *UserProcess) NotifyOthersOnlineUser(userId int)  {
	fmt.Println("通知其他在线上线")

	for id, up:=range  userMgr.OnlineUser {
		if id == userId {
			continue
		}
		up.NotityMetoOther(userId)
		
	}
}

func (this *UserProcess) NotityMetoOther(userId int)  {
	var mes message.Message
	mes.Type = message.NotyfyUserStatusMesType
	var notifyUserStatues message.NotyfyUserStatusMes
	notifyUserStatues.UserId = userId
	notifyUserStatues.Status = message.UserOnline
	data, err := json.Marshal(notifyUserStatues)
	if err != nil {
		fmt.Println("ata, err := json.Marshal(notifyUserStatues) err=", err)
		return
	}
	mes.Date = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("ata, err := json.Marshal(notifyUserStatues) err=", err)
		return
	}

	tr := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tr.WritePkg(data)
	if err != nil {
		fmt.Println("err = tr.WritePkg(data) err=", err)
		return
	}
}