package main

import (
	"chatroom/client/processes"
	"fmt"
	"os"
)

var userId int
var userPwd string
var userName string


func main()  {

	var  key int
	for {
			fmt.Println("-----------欢迎登陆多人聊天系统--------")
			fmt.Println("\t\t\t 1 登陆聊天室")
			fmt.Println("\t\t\t 2 注册用户")
			fmt.Println("\t\t\t 3 退出系统")
			fmt.Println("\t\t\t 请选择（1-3）：")

			fmt.Scanf("%d\n",&key)

		switch key {
			case 1:
				fmt.Println("登陆聊天室")

				fmt.Println("请输入用户的id")
				fmt.Scanf("%d\n",&userId)
				fmt.Println("请输入密码")
				fmt.Scanf("%s\n", &userPwd)

				us :=&processes.UserProcessor{}
				err := us.Login(userId,userPwd)
				if err != nil {
					fmt.Println("登陆成功")
				} else {
					fmt.Println("登陆失敗")
				}

			case 2:
				fmt.Println("注册用户")

				fmt.Println("请输入用户的id")
				fmt.Scanf("%d\n",&userId)
				fmt.Println("请输入密码")
				fmt.Scanf("%s\n", &userPwd)
				fmt.Println("请输入昵称")
				fmt.Scanf("%s\n", &userName)
				us :=&processes.UserProcessor{}
				err := us.Register(userId, userPwd, userName)
				if err != nil {
					fmt.Println("注册成功")
				} else {
					fmt.Println("注册失败")
				}

			case 3:
				fmt.Println("退出系统")
				os.Exit(0)
			default:
				fmt.Println("输入有误")
		
		}

	}

	
}
















