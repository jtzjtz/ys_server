package redisdb

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jtzjtz/kit/conn/redis_pool"
	"log"
	"time"
)

var RedisPool *redis.Pool

func Init() {
	options := &redis_pool.Options{
		Host:               "172.20.98.10:6379",
		PassWord:           "",
		Database:           0,
		InitCap:            10,
		MaxCap:             100,
		IsWait:             true,
		IdleTimeout:        5 * time.Second,
		DialConnectTimeout: 5 * time.Second,
		DialReadTimeout:    5 * time.Second,
		DialWriteTimeout:   5 * time.Second,
	}

	var err error

	RedisPool, err = redis_pool.NewRedisPool(options)

	if err != nil {
		log.Printf("%#v\n", err)
		return
	}
}
