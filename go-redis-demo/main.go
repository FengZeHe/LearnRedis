package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "your_password",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successful connection")
	}

	err = rdb.Set(ctx, "Jim", "man", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "Jim").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

}
