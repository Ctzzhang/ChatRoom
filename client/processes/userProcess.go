package processes

import (
	"chatroom/client/model"
	"chatroom/client/utils"
	"chatroom/common/message"
"encoding/binary"
"encoding/json"
"fmt"
"net"
	"os"
)

type  UserProcessor struct {

}

func (this *UserProcessor)Login(userId int, userPwd string) (error error)  {
	fmt.Println("登陆成功用户名%d， 密码%s \n" ,userId, userPwd)
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("链接报错")
	}
	defer conn.Close()


	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	data, err :=json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
	}
	mes.Date = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
	}
	// 长度和内容
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf []byte
	buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	n, err := conn.Write(buf)
	if n !=4 || err != nil {
		fmt.Println("onn.Write(bytes) err=", err)
		return
	}
	fmt.Printf("消息的长度已经发送=%d, %v\n", len(data), string(data))

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("onn.Write(data) err=", err)
		return
	}

	ut :=&utils.Transfer{
		Conn: conn,
	}

	mes, err = ut.ReadPkg()
	if err != nil {
		fmt.Println("接受消息失败err =", err )
	}
	fmt.Println("登陆返回数据 mes = ", mes)
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Date), &loginResMes)
	if loginResMes.Code == 200 {

		curUser.Conn = conn
		curUser.UserId = userId
		curUser.UserStatus = message.UserOnline



		fmt.Println("在线用户列表")
		for _, v := range  loginResMes.UserIds {
			fmt.Println("在线用户id=", v)
			
			user:=&model.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUser[v] = user
		}


		go ServerProcessMes(conn)
		for  {
			ShowMenu()
		}
		fmt.Println(loginResMes)
	} else  {
		fmt.Println(loginResMes)
	}

	return
}

func (this *UserProcessor) Register(userId int, userPwd string, userName string) (err error) {

	fmt.Println("客户端注册开始了用户名%d， 密码%s \n" ,userId, userPwd)
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("链接报错")
	}
	defer conn.Close()


	var mes message.Message
	mes.Type = message.RegisterMesType

	var registerMes message.RegisterMes
	registerMes.UserId = userId
	registerMes.UserPwd = userPwd
	registerMes.UserName = userName

	data, err :=json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
	}
	mes.Date = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
	}
	// 长度和内容
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf []byte
	buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	n, err := conn.Write(buf)
	if n !=4 || err != nil {
		fmt.Println("onn.Write(bytes) err=", err)
		return
	}
	fmt.Printf("消息的长度已经发送=%d, %v\n", len(data), string(data))

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("onn.Write(data) err=", err)
		return
	}

	ut :=&utils.Transfer{
		Conn: conn,
	}

	mes, err = ut.ReadPkg()
	if err != nil {
		fmt.Println("接受消息失败err =", err )
	}
	fmt.Println("登陆返回数据 mes = ", mes)
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Date), &registerResMes)
	if registerResMes.Code == 200 {

		fmt.Println("注册成功", registerResMes)
		os.Exit(0)
	} else  {
		fmt.Println("注册失败",registerResMes)
		os.Exit(0)
	}

	return
}

















