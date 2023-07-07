package gredis

import (
	"context"
	"log"

	//"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
)

var ctx = context.Background()

func InitRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

// 查询用户key是否存在
func ExistUserKey(key string) bool {
	n, err := rdb.Exists(ctx, key).Result()
	if err != nil {
		log.Println("find exist user Key error :", err)
	}
	if n == 0 {
		log.Println(key, "key不存在")
		return false
	}
	log.Println(key, "key 存在")
	return true
}
