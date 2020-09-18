package processes

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func ShowMenu()  {
	fmt.Println("----------恭喜xxx登陆成功------------")
	fmt.Println("----------1 显示在线用户列表----------")
	fmt.Println("----------2 发送消息-----------------")
	fmt.Println("----------3 消息列表-----------------")
	fmt.Println("----------4 退出系统-----------------")
	fmt.Println("----------请选择（1-4）------------")
	var key int
	_, _ = fmt.Scanf("%d\n", &key)
	var content string

	smp := &SmsProcess{
	}

	switch key {
		case 1:
			fmt.Println("显示在线列表")
			OutputOnlineUser()
		case 2:
			fmt.Println("发送消息")
			fmt.Println("请输入你想对大家说的")
			fmt.Scanf("%s\n", &content)
			smp.SendGroup(content)
		case 3:
			fmt.Println("消息列表")
		case 4:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("错误输入")
	}
}

func ServerProcessMes(conn net.Conn)  {
	tr := &utils.Transfer{
		Conn: conn,
	}
	for  {

		fmt.Println("客户端等待接收服务器的发送的信息")
		mes, err := tr.ReadPkg()
		if err != nil {
			fmt.Println("tr.ReadPkg() err=", err)
			return
		}
		fmt.Printf("mes=%v\n", mes)

		switch mes.Type {
			case message.NotyfyUserStatusMesType:
				fmt.Println("有人上线")
				var notifyUserStatusMes message.NotyfyUserStatusMes
				json.Unmarshal([]byte(mes.Date), &notifyUserStatusMes)
				UpdateUserStatus(&notifyUserStatusMes)
			case message.SmsMesType:
				fmt.Println("有人群发")
				OutPutGroupSms(&mes)
		default:
			fmt.Println("客户端不能处理这个消息的类型")

		}


	}

}
