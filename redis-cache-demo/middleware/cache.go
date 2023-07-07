package middleware

import (
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
		isExists := gredis.ExistUserKey(key)
		if isExists == false { // 缓存不存在, 查询sql ,写入redis缓存
			c.Next()
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
}
