package processes

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
)

type  SmsProcess struct {
	
}

func (this *SmsProcess)SendGroup(content string) (err error) {
	var  mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.User.UserId = curUser.UserId
	smsMes.User.UserStatus = curUser.UserStatus

	data, err:= json.Marshal(smsMes)
	if err != nil {
		fmt.Println("sendgroupMes json.Marshal(smsMes) err= ", err)
	}
	mes.Date = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("data, err = json.Marshal(mes) err=", err)
	}
	tr := &utils.Transfer{
		Conn: curUser.Conn,
	}
	err = tr.WritePkg(data)
	if err != nil {
		fmt.Println("data, err = json.Marshal(mes) err=", err)
	}
	return
}
