package mysql

import (
	"log"
	"redisdemo/model"
)

// 查询所有的Post数据
func QueryAllPosts() (post []*model.Post, err error) {
	err = db.Table("post").Find(&post).Error
	if err != nil {
		log.Println("QueryAllPosts error:", err)
		return nil, err
	}
	return post, nil
}
