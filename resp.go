package goRedis

type Resp struct {
	Code    int
	Message string
	Data    interface{}
}
