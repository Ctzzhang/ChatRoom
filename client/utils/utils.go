package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf [8090]byte
}


func (this *Transfer)WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))


	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)

	_, err = this.Conn.Write(this.Buf[:4])
	if err != nil {
		fmt.Println("conn.write() err=", err)
		return
	}

	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("conn.write()err=", err)
		return
	}


	return
}

func (this *Transfer)ReadPkg() (mes message.Message, err error)  {

	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		fmt.Println("接受buf err=", err)
		return
	}
	fmt.Println("接受buf =", this.Buf[:4])

	var dataLen uint32
	dataLen = binary.BigEndian.Uint32(this.Buf[:4])
	n, err := this.Conn.Read(this.Buf[:dataLen])
	if n != int(dataLen) || err !=nil {
		fmt.Println("接受buf err=", err)
		return
	}
	err = json.Unmarshal(this.Buf[:dataLen], &mes)
	if err != nil {
		fmt.Println("接受buf err=", err)
		return
	}
	return
}
