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

func (goRedis *GoRedis) Conn(clientName string) *Resp {
	ctx := context.Background()
	cn := goRedis.Client.Conn(ctx)
	defer cn.Close()
	cn.ClientSetName(ctx, clientName)
	name, err := cn.ClientGetName(ctx).Result()
	var resp = &Resp{}
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "ok"
		resp.Data = name
	}
	return resp
}

func (goRedis *GoRedis) Publish(channel string, message string) *Resp {
	ctx := context.Background()
	rel := goRedis.Client.Publish(ctx, channel, message)
	var resp = &Resp{}
	if rel.Err() != nil {
		resp.Code = -1
		resp.Message = rel.Err().Error()
	} else {
		resp.Code = 1
		resp.Message = "ok"
		resp.Data, _ = rel.Result()
	}
	return resp
}

func (goRedis *GoRedis) Subscribe(channel string) *Resp {
	ctx := context.Background()
	goRedis.Client.Subscribe(ctx, channel)
	var resp = &Resp{}
	resp.Code = 1
	resp.Message = "ok"
	return resp
}

func (goRedis *GoRedis) CloseSub(pubSub *redis.PubSub) {
	pubSub.Close()
}

func (goRedis *GoRedis) Do(param map[string]interface{}) *Resp {
	ctx := context.Background()
	cmd := param["cmd"]
	rel := goRedis.Client.Do(ctx, cmd)
	var resp = &Resp{}
	if rel.Err() != nil {
		resp.Code = -1
		resp.Message = rel.Err().Error()
	} else {
		resp.Code = 1
		resp.Message = "ok"
		resp.Data = rel
	}
	return resp
}
