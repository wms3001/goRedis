package goRedis

import (
	"context"
	"github.com/go-redis/redis/v9"
	"time"
)

type GoRedis struct {
	Addr     string
	Password string
	Db       int
	IsTSL    bool
	Client   *redis.Client
}

func (goRedis *GoRedis) Connect() {
	option := &redis.Options{}
	option.Addr = goRedis.Addr
	option.Password = goRedis.Password
	option.DB = goRedis.Db
	rdb := redis.NewClient(option)
	goRedis.Client = rdb
}

func (goRedis *GoRedis) Set(key string, val interface{}, expire time.Duration) *Resp {
	ctx := context.Background()
	rel := goRedis.Client.Set(ctx, key, val, expire)
	var resp = &Resp{}
	if rel.Err() != nil {
		resp.Code = -1
		resp.Message = rel.Err().Error()
	} else {
		resp.Code = 1
		resp.Message = rel.String()
		resp.Data = rel.Val()
	}
	return resp
}

func (goRedis *GoRedis) Get(key string) *Resp {
	ctx := context.Background()
	rel := goRedis.Client.Get(ctx, key)
	var resp = &Resp{}
	if rel.Err() != nil {
		resp.Code = -1
		resp.Message = rel.Err().Error()
	} else {
		resp.Code = 1
		resp.Message = rel.String()
		resp.Data, _ = rel.Result()
	}
	return resp
}
