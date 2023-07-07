package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"redisdemo/gredis"
	"redisdemo/logic"
	"redisdemo/model"
)

/*
获取全部的post
1. 如果redis中key不存在，缓存不存在查询sql -> 返回json -> 将返回结果存到redis(kv存储)
2. 如果redis中key存在，则取出redis缓存,返回json数据
*/
func HandleGetAllPost(c *gin.Context) {
	var data []*model.Post
	var err error
	data, err = logic.GetAllPosts()
	if err != nil {
		log.Println("logic GETAllPosts error:", err)
		ResponseError(c, data)
	}
	ResponseSuccess(c, data)
	if err = gredis.SetCacheAllPosts(data); err != nil {
		log.Println("gredis SET Cache All Posts ERROR:", err)
	}
}
