package processes

import (
	"chatroom/common/message"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {

}

func (this *SmsProcess)SendGroupMes(mes *message.Message) (err error) {

	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Date), &smsMes)
	if err != nil {
		fmt.Println("err = json.Unmarshal([]byte(mes.Date), smsMes) err=", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("err = json.Unmarshal([]byte(mes.Date), smsMes) err=", err)
		return
	}

	for id, up := range userMgr.OnlineUser {
		if id == smsMes.User.UserId {
			continue
		}
		this.SendMesToOther(data, up.Conn)
	}
	return
}

func (this *SmsProcess)SendMesToOther(data []byte, conn net.Conn)  {
	tr :=&utils.Transfer{
		Conn: conn,
	}

	err := tr.WritePkg(data)
	if err != nil {
		fmt.Println("SendMesToOther err=", err)
	}
	fmt.Println("群发消息结束")
}



