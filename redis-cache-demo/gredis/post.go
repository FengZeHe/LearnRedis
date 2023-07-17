package gredis

import (
	"encoding/json"
	"fmt"
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

// 根据文章Id获取缓存
func GetCachePostById(key string) (data *model.Post, err error) {
	res, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Println("GET redis post error:", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(res), &data)
	return data, nil
}

// 根据文章Id设置缓存
func SetCachePostById(data *model.Post, postid string) (err error) {
	strdata, _ := json.Marshal(data)
	key := fmt.Sprintf("%s%s", KeyPostIdSet, postid)
	err = rdb.Set(ctx, key, strdata, 10*time.Second).Err()
	if err != nil {
		log.Println("SET redis ERROR:", err)
		return err
	}
	return nil
}

// 更新文章缓存
func UpdatePost(key string) {
	rdb.Publish(ctx, "post_cache", key)
}
