package model

import (
	"chatroom/server/db"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func InitUserDao()  {
	MyUserDao = NewUserDao(db.Pool)
}


var (
	MyUserDao *UserDao
)


type  UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (this *UserDao)getUserById(conn redis.Conn, userId int) (user *User, err error)  {
	res, err := redis.String(conn.Do("HGet", "user", userId))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTXEISTS
		}
		return
	}
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Marshal([]byte(res) srr = ", err)
		return 
	}

	return
}

func (this * UserDao)Login(userId int, userPwd string) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}

	return
}

func (this *UserDao) Register(userId int, userPwd string, userName string) (user User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, userId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	user.UserId = userId
	user.UserPwd = userPwd
	user.UserName = userName

	data, err := json.Marshal(user)
	if err!=nil {
		fmt.Scanf("json.Marshal(user) err = ", err)
	}

	fmt.Printf("json.Marshal(user) %v\n", string(data))

	_,err = conn.Do("Hset", "user", userId, string(data))
	if err != nil {
		fmt.Println("报错报错")
		return
	}
	return
}
