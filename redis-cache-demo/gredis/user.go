package gredis

import (
	"encoding/json"
	"log"
	"redisdemo/model"
	"time"
)

// 定义一个获取用户缓存 key的方法
func GetCacaheKey(id string) string {
	return KeyUserIdSet + ":" + id
}

// 缓存全部用户
func SetCacheAllUsers(data []*model.User) (err error) {
	strdata, _ := json.Marshal(data)
	err = rdb.Set(ctx, CACHE_USERS, strdata, 5*time.Second).Err()
	if err != nil {
		log.Println("redis set error", err)
		return err
	}
	return nil
}

// GET 获取全部用户缓存
func GetCacheAllUsers(key string) (data []*model.User, err error) {
	res, err := rdb.Get(ctx, key).Result()
	err = json.Unmarshal([]byte(res), &data)
	return data, nil
}

// Set单个用户信息的缓存
func SetCacheUserById(key string, data *model.User) (err error) {
	strdata, _ := json.Marshal(data)
	err = rdb.Set(ctx, key, strdata, 10*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

// Get单个用户信息的缓存
func GetCacheUserById(key string) (data *model.User, err error) {
	res, err := rdb.Get(ctx, key).Result()
	err = json.Unmarshal([]byte(res), &data)
	return data, nil
}

// 删除单个用户的缓存信息
func DelCacheUserById(key string) (err error) {
	err = rdb.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
