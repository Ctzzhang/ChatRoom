package processes

import (
	"chatroom/client/model"
	"chatroom/common/message"
	"fmt"
)

var onlineUser map[int]*model.User = make(map[int]*model.User, 10)
var curUser model.CurUser

func OutputOnlineUser()  {
	for id, user := range onlineUser {
		fmt.Println("用户id：\t", id)
		fmt.Println("yonghua\t", user.UserStatus)
	}
}


func UpdateUserStatus(mes *message.NotyfyUserStatusMes) {

	user, ok := onlineUser[mes.UserId]
	if !ok {
		user =&model.User{
			UserId: mes.UserId,
			UserStatus : mes.Status,
		}
	}
	user.UserStatus = mes.Status
	onlineUser[mes.UserId] = user
}



