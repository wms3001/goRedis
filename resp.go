package goRedis

import (
	"context"
	"github.com/go-redis/redis/v9"
)

type Resp struct {
	Code    int
	Message string
	Data    interface{}
	PubSub  *redis.PubSub
	Ctx     *context.Context
}
