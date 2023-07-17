package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"redisdemo/controller"
	"redisdemo/gredis"
	"redisdemo/middleware"
	"redisdemo/mysql"
)

func main() {
	err := gredis.InitRedis()
	if err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		panic(err)
	}
	fmt.Println("redis connect success!")

	err = mysql.InitMysql()
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		panic(err)
	}

	r := gin.Default()
	/*
		Get Users
	*/
	r.GET("/getalluser", middleware.CacheMiddleware(gredis.CACHE_USERS), controller.HandleGetAllUsers)
	r.GET("/getuserbyid/:id", controller.HandleQueryUserById)
	r.DELETE("/deluserbyid/:id", controller.HandleDeleteUserById)
	r.POST("/updateuserbyid", controller.HandeUpdateUserById)

	/*
		Get Posts
	*/
	r.GET("/getallpost", middleware.CacheMiddleware(gredis.CACHE_POSTS), controller.HandleGetAllPost)
	r.GET("/getpostbyid/:id", controller.HandleGetPostById)
	r.GET("/getpostbypostid/:postid", middleware.CacheMiddleware(gredis.KeyPostIdSet), controller.HandleGetPostByPostId)
	r.POST("/updatepostbyid", controller.HandleUpdatePostById)
	r.DELETE("/deletepostbyid/:id", controller.HandleDeletePostById)

	r.Run(":9001")
}
