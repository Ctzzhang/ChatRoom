package processes

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
)

func OutPutGroupSms(mes *message.Message)  {

	fmt.Println()
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Date), &smsMes)
	if err != nil {
		fmt.Println("\terr := json.Unmarshal([]byte(mes.Date), &smsMes) err=", err)
		return
	}

	info, _ := fmt.Printf("用户id: %d 对大家说：%s\n", smsMes.User.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()





}














