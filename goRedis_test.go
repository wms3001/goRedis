package goRedis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
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

func TestGoRedis_Conn(t *testing.T) {
	goRedis := GoRedis{}
	goRedis.Db = 10
	goRedis.Addr = "192.168.4.81:6379"
	goRedis.Password = "glredis"
	goRedis.Connect()
	goRedis.Conn("myClient")
	//resp := goRedis.Conn("myClient")
	//fmt.Println(gconv.Map(resp))
}

func TestGoRedis_Publish(t *testing.T) {
	goRedis := GoRedis{}
	goRedis.Db = 10
	goRedis.Addr = "192.168.4.81:6379"
	goRedis.Password = "glredis"
	goRedis.Connect()
	goRedis.Publish("testChannel", "213424234")
	goRedis.Publish("testChannel", "213424234")
	goRedis.Publish("testChannel", "213424234")
	goRedis.Publish("testChannel", "213424234")
	goRedis.Publish("testChannel", "213424234")
	goRedis.Publish("testChannel", "213424234")
	goRedis.Publish("testChannel", "213424234")
	goRedis.Publish("testChannel", "213424234")
	goRedis.Publish("testChannel", "213424234")
	goRedis.Publish("testChannel", "213424234")
	goRedis.Publish("testChannel", "213424234")
	goRedis.Publish("testChannel", "213424234")
	goRedis.Publish("testChannel", "213424234")
	goRedis.Publish("testChannel", "213424234")
	goRedis.Publish("testChannel", "213424234")
	goRedis.Publish("testChannel", "213424234")
	goRedis.Publish("testChannel", "213424234")
	goRedis.Publish("testChannel", "213424234")
	goRedis.Publish("testChannel", "213424234")
	goRedis.Publish("testChannel", "213424234")
	goRedis.Publish("testChannel", "213424234")

}

func TestTt(t *testing.T) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.4.81:6379",
		Password: "glredis", // no password set
		DB:       10,        // use default DB
	})
	err := rdb.Publish(ctx, "mychannel1", "payload").Err()
	if err != nil {
		panic(err)
	}
	pubsub := rdb.Subscribe(ctx, "mychannel1")

	// Close the subscription when we are done.
	defer pubsub.Close()
	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}

func TestGoRedis_Subscribe(t *testing.T) {
	goRedis := GoRedis{}
	goRedis.Db = 10
	goRedis.Addr = "192.168.4.81:6379"
	goRedis.Password = "glredis"
	goRedis.Connect()
	pubsub := goRedis.Subscribe("testChannel")
	//defer goRedis.CloseSub(pubsub.PubSub)
	pubsub.PubSub.Receive(context.Background())
	ch := pubsub.PubSub.Channel()
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
	//ctx := context.Background()
	//for {
	//	msg, err := resp.PubSub.ReceiveMessage(*resp.Ctx)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	fmt.Println(msg.Channel, msg.Payload)
	//}
}
