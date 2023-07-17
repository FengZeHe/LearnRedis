package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"redisdemo/gredis"
	"redisdemo/logic"
	"redisdemo/model"
)

// 获取全部用户
func HandleGetAllUsers(c *gin.Context) {
	var data []*model.User
	var err error
	data, err = logic.GetAllUsers()
	if err != nil {
		log.Println("logic GetAllUsers error")
	}
	if err = gredis.SetCacheAllUsers(data); err != nil {
		log.Println("Set Cache error:", err)
	}

	ResponseSuccess(c, data)
}

// 根据id查询用户
func HandleQueryUserById(c *gin.Context) {
	/*
		1.生成key name,查询redis是否存在这个key

		2.如果存在则直接返回

		3.如不存在，则查询Mysql数据库

		4.查询失败-> 返回；查询成功-> 返回，并将信息放进redis缓存中，并设置15s的过期时间

	*/
	var data *model.User
	id := c.Param("id")
	key := gredis.GetCacaheKey(id)
	isExist := gredis.ExistKey(key)
	var err error
	if isExist == false {
		data, err = logic.GetUserById(id)
		if err != nil {
			fmt.Println("query failed")
		}
		err = gredis.SetCacheUserById(key, data)
		if err != nil {
			fmt.Println("cache user by id error:", err)
		}
		ResponseSuccess(c, data)
	} else {
		data, err = gredis.GetCacheUserById(key)
		ResponseSuccess(c, data)
	}

}

// 根据用户id删除
func HandleDeleteUserById(c *gin.Context) {
	/*
		1. 查询redis中key是否存在，存在的话删掉，不存在就跳过。
		2. 在Mysql中根据用户Id删除该用户
	*/
	id := c.Param("id")
	fmt.Println(id)
	key := gredis.GetCacaheKey(id)
	isExist := gredis.ExistKey(key)
	if isExist == true {
		err := gredis.DelCacheUserById(key)
		if err != nil {
			fmt.Println("redis del key error:", err)
		}
	}
	err := logic.DeleteUserById(id)
	if err != nil {
		fmt.Println("Mysql delete user  error", err)
	}
	ResponseSuccess(c, "success")
}

// 根据用户id更新
func HandeUpdateUserById(c *gin.Context) {
	/*
		1. 直接更新Mysql数据库
		2. 更改完后删除对应redis的key
		3. 重新查询，并写入redis缓存中
	*/
	id := c.PostForm("user_id")
	email := c.PostForm("email")
	err := logic.UpdateUserById(id, email)
	if err != nil {
		fmt.Println("logic update user error:", err)
	}
	key := gredis.GetCacaheKey(id)
	isExist := gredis.ExistKey(key)
	if isExist == true {
		gredis.DelCacheUserById(key)
	}

	ResponseSuccess(c, "success")

}
