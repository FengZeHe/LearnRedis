package logic

import (
	"log"
	"redisdemo/model"
	"redisdemo/mysql"
)

func GetAllPosts() (data []*model.Post, err error) {
	data, err = mysql.QueryAllPosts()
	if err != nil {
		log.Println("GetAllPosts error:", err)
		return nil, err
	}
	return data, nil
}
