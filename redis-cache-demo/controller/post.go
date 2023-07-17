package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"redisdemo/gredis"
	"redisdemo/logic"
	"redisdemo/model"
	"strconv"
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

/*
根据post id获取post全部信息
*/
func HandleGetPostByPostId(c *gin.Context) {
	postid := c.Param("postid")
	data, err := logic.GetPostByPostId(postid)
	if err != nil {
		log.Println("logic GetPostByPostId error", err)
		ResponseError(c, data)
	}
	ResponseSuccess(c, data)
	// 设置redis缓存
	if err = gredis.SetCachePostById(data, postid); err != nil {
		log.Println("Set Cache Post By Post Id ERROR:", err)
	}
}

/*
根据id获取post的信息
*/
func HandleGetPostById(c *gin.Context) {
	id := c.Param("id")
	data, err := logic.GetPostById(id)
	if err != nil {
		log.Println("logic GetPostById error", err)
		ResponseError(c, data)
	}
	ResponseSuccess(c, data)
}

/*
修改文章内容
*/
func HandleUpdatePostById(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("ID"))
	title := c.PostForm("Title")
	content := c.PostForm("Content")
	currentpost := model.Post{Id: id, Title: title, Content: content}
	if err := logic.UpdatePostById(&currentpost, currentpost); err != nil {
		log.Println("logic Update Post By Id ERROR", err)
		ResponseError(c, "update error")
	}
	ResponseSuccess(c, "success")

	/*
		修改完成后删掉key
	*/
	data, err := logic.GetPostById(c.PostForm("ID"))
	if err != nil {
		fmt.Println("logic Query PostById ERROR:", err)
	}
	currentKey := fmt.Sprintf("%s%s", gredis.KeyPostIdSet, data.PostId)
	gredis.UpdatePost(currentKey)

}

// 删除文章内容
func HandleDeletePostById(c *gin.Context) {
	id := c.Param("id")
	// 先查Mysql中 post_id
	data, err := logic.GetPostById(id)
	if err != nil {
		fmt.Println("logic Query PostById ERROR:", err)
	}

	if err = logic.DeletePostById(id); err != nil {
		fmt.Println("logic Delete PostById ERROR:", err)
		ResponseError(c, "error")
	}
	ResponseSuccess(c, "success")

	postid := data.PostId
	currentKey := fmt.Sprintf("%s%s", gredis.KeyPostIdSet, postid)
	if err := gredis.SetKeyExpired(currentKey); err != nil {
		log.Println("gredis SetKeyExpired ERROR:", err)
	}
}
