package processes

import "fmt"

var (
	userMgr * UserMgr
)

type UserMgr struct {
	OnlineUser map[int]*UserProcess
}

func init()  {
	userMgr = &UserMgr{
		OnlineUser : make(map[int]*UserProcess, 1024),
	}
}

func (this *UserMgr) AddOnlineUser(up * UserProcess)  {
	this.OnlineUser[up.UserId] = up
}

func (this *UserMgr) DelOnlineUser(userId int)  {
	delete(this.OnlineUser, userId)
}

func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.OnlineUser;
}

func (this *UserMgr) GetOnlineUserById(userId int)(up *UserProcess, err error) {
	up, ok := this.OnlineUser[userId]
	if !ok {
		fmt.Println()
		err = fmt.Errorf("用户不存在 %d \n", userId)
		return
	}
	return
}


