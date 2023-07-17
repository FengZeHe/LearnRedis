package gredis

import (
	"context"
	"log"
	"time"

	//"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

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
	SubChannel()
	return nil
}

// 订阅 channel
func SubChannel() {
	sub := rdb.Subscribe(ctx, "post_cache")
	ch := sub.Channel()
	go func() {
		for msg := range ch {
			if err := DeleteKey(msg.Payload); err != nil {
				log.Println("delete key ERROR:", err)
			}
		}
	}()
}

// 查询用户key是否存在
func ExistKey(key string) bool {
	n, err := rdb.Exists(ctx, key).Result()
	if err != nil {
		log.Println("find exist user Key error :", err)
	}
	if n == 0 {
		log.Println(key, "key no exist")
		return false
	}
	log.Println(key, "key exist")
	return true
}

// 设置key过期
func SetKeyExpired(key string) (err error) {
	err = rdb.ExpireAt(ctx, key, time.Now().Add(-10*time.Second)).Err()
	if err != nil {
		log.Println("Set Key Expired ERROR:", err)
		return err
	}
	return nil
}

// 删除某个key
func DeleteKey(key string) (err error) {
	if err = rdb.Del(ctx, key).Err(); err != nil {
		log.Println("Delete Key  ERROR:", err)
		return err
	}
	return nil
}
