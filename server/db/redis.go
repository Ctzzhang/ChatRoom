package db

import (
	"github.com/garyburd/redigo/redis"
	"time"
)


var Pool *redis.Pool

func InitPool(address string, mxIdle, maxActivie int, idleTimeout time.Duration)  {
	Pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
		MaxIdle:         mxIdle,
		MaxActive:       maxActivie,
		IdleTimeout:     idleTimeout,
	}
	
}
