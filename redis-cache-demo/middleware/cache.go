package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"redisdemo/controller"
	"redisdemo/gredis"
)

/*
将redis缓存的逻辑写在这里
1. 功能是判断redis中是否存在key,如果存在则取出缓存并返回数据；c.Abort
2. 如果redis中key不存在，则c.Next()继续查询数据库
*/
func CacheMiddleware(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 判断一下是哪种key
		if key == gredis.CACHE_POSTS || key == gredis.CACHE_USERS {
			if isExists := gredis.ExistKey(key); isExists == false {
				c.Next() // 缓存不存在, 查询sql ,写入redis缓存
			} else {
				// 取出缓存
				switch key {
				case gredis.CACHE_POSTS:
					data, _ := gredis.GetCacheAllPosts(key)
					controller.ResponseSuccess(c, data)
				case gredis.CACHE_USERS:
					data, _ := gredis.GetCacheAllUsers(key)
					controller.ResponseSuccess(c, data)
				}
				c.Abort()
			}
		}
		if key == gredis.KeyPostIdSet || key == gredis.KeyUserIdSet {
			postId := c.Param("postid")
			currentKey := fmt.Sprintf("%s%s", key, postId)
			if isExists := gredis.ExistKey(currentKey); isExists == false {
				c.Next() // 缓存不存在, 查询sql ,写入redis缓存
			} else {
				switch key {
				case gredis.KeyPostIdSet:
					data, _ := gredis.GetCachePostById(currentKey)
					controller.ResponseSuccess(c, data)
				}
				c.Abort()
			}
		}
	}
}
