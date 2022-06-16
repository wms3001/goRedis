package goRedis

import (
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	goRedis := GoRedis{}
	goRedis.Db = 10
	goRedis.Addr = "192.168.4.81:6379"
	goRedis.Password = "glredis"
	goRedis.Connect()
}

func BenchmarkConnect(b *testing.B) {
	goRedis := GoRedis{}
	goRedis.Db = 10
	goRedis.Addr = "192.168.4.81:6379"
	goRedis.Password = "glredis"
	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		goRedis.Connect()
	}
}

func TestGoRedis_Set(t *testing.T) {
	goRedis := GoRedis{}
	goRedis.Db = 10
	goRedis.Addr = "192.168.4.81:6379"
	goRedis.Password = "glredis"
	goRedis.Connect()
	goRedis.Set("glkj", "6666666", 1*time.Minute)
	//resp := goRedis.Set("glkj","6666666",1*time.Minute)
	//fmt.Println(gconv.Map(resp))
}

func TestGoRedis_Get(t *testing.T) {
	goRedis := GoRedis{}
	goRedis.Db = 10
	goRedis.Addr = "192.168.4.81:6379"
	goRedis.Password = "glredis"
	goRedis.Connect()
	goRedis.Get("glkj")
	//resp := goRedis.Get("glkj")
	//fmt.Println(gconv.Map(resp))
}
