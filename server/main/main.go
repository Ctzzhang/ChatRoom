package main

import (
	"chatroom/server/db"
	"chatroom/server/model"
	"fmt"
	"net"
	"time"
)

func main()  {

	db.InitPool("localhost:6379", 16, 0, 300*time.Second)
	model.InitUserDao()
	fmt.Println("服务器在8889监听")
	listen, err :=net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err!= nil {
		fmt.Println("监听失败 err = ", err)
		return
	}
	for  {
		fmt.Println("等待客户端还链接服务器 。。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept 链接失败 err = ", err)
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	//读取信息
	fmt.Println("链接成功")

/*	for  {
		mes, err := readPkg(conn)
		if err != nil {
			fmt.Println("接受buf err=", err)
			return
		}
		fmt.Println("mes=", mes)
		err = serverProcessMes(conn, &mes)
		if err != nil {
			fmt.Println("返回消息失败")
		}
	}*/

	po := &Process{
		Conn: conn,
	}
	po.processAll()
}
/*
func readPkg(conn net.Conn) (mes message.Message, err error)  {
	buf := make([]byte, 8096)
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("接受buf err=", err)
		return
	}
	fmt.Println("接受buf =", buf[:4])

	var dataLen uint32
	dataLen = binary.BigEndian.Uint32(buf[:4])
	n, err :=conn.Read(buf[:dataLen])
	if n != int(dataLen) || err !=nil {
		fmt.Println("接受buf err=", err)
		return
	}
	err = json.Unmarshal(buf[:dataLen], &mes)
	if err != nil {
		fmt.Println("接受buf err=", err)
		return
	}
	return
}*/

/*func serverProcessMes(conn net.Conn, mes *message.Message) (err error)  {
	switch mes.Type {
		case message.LoginMesType:
			err = serverProcessLogin(conn, mes)
		case message.RegisterMesType:

		default:
			fmt.Println("消息类型不存在 无法处理")
	
	}
	return
}*/
/*
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	var loginmes message.LoginMes
	err = json.Unmarshal([]byte(mes.Date), &loginmes)
	if err != nil {
		fmt.Println("json.unmarshal err = ", err)
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes  message.LoginResMes

	if loginmes.UserId == 123 && loginmes.UserPwd == "123456" {
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 500
		loginResMes.Err = "用户不存在"
	}

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
	err = writePkg(conn, data)
	if err != nil {
		fmt.Println("服务器发送数据失败")
	}

	return
}*/

/*func writePkg(conn net.Conn, data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf []byte
	buf = make([]byte,4)

	binary.BigEndian.PutUint32(buf[:4], pkgLen)

	_, err = conn.Write(buf[:4])
	if err != nil {
		fmt.Println("conn.write() err=", err)
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.write()err=", err)
		return
	}


	return
}*/
















