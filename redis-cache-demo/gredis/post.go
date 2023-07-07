package gredis

import (
	"encoding/json"
	"log"
	"redisdemo/model"
	"time"
)

// 缓存全部文章
func SetCacheAllPosts(data []*model.Post) (err error) {
	strdata, _ := json.Marshal(data)
	err = rdb.Set(ctx, "CACHE/all-posts", strdata, 10*time.Second).Err()
	if err != nil {
		log.Println("SET redis CACHE/all-posts error", err)
		return err
	}
	return nil
}

// 获取文章缓存
func GetCacheAllPosts(key string) (data []*model.Post, err error) {
	res, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Println("GET redis post error:", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(res), &data)
	return data, nil
}
