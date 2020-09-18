package main

import (
	"chatroom/common/message"
	"chatroom/server/processes"
	"chatroom/server/utils"
	"fmt"
	"net"
)

type Process struct {
	Conn net.Conn
}

func (this *Process)ServerProcessMes(mes *message.Message) (err error)  {
	switch mes.Type {
	case message.LoginMesType:
		us :=&processes.UserProcess{
			Conn : this.Conn,
		}
		err = us.ServerProcessLogin(mes)
	case message.RegisterMesType:
		us :=&processes.UserProcess{
			Conn : this.Conn,
		}
		err = us.ServerProcessRegister(mes)
	case message.SmsMesType:
		fmt.Println("群发信息")
		smp := &processes.SmsProcess{}
		smp.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在 无法处理")

	}
	return
}

func (this *Process)processAll()  {

	for  {
		ut := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := ut.ReadPkg()
		if err != nil {
			fmt.Println("接受buf err=", err)
			return
		}
		fmt.Println("mes=", mes)
		err = this.ServerProcessMes(&mes)
		if err != nil {
			fmt.Println("返回消息失败")
		}
	}
	
}