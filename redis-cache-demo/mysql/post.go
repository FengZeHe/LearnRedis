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

// 根据post id查询数据
func QueryPostByPostId(postid string) (post *model.Post, err error) {
	if err = db.Table("post").Where("post_id = ?", postid).Find(&post).Error; err != nil {
		return nil, err
	}
	return post, err
}

// 根据id查询数据
func QueryPostById(id string) (post *model.Post, err error) {
	if err = db.Table("post").Where("id = ?", id).Find(&post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func UpdatePostById(postModel *model.Post, post model.Post) (err error) {
	if err = db.Table("post").Model(&postModel).Updates(post).Error; err != nil {
		return err
	}
	return nil
}

func DeletePostById(id string) (err error) {
	var post model.Post
	err = db.Table("post").Delete(&post, id).Error
	if err != nil {
		log.Println("mysql Delete Post By Id ERROR:", err)
		return err
	}
	return nil
}
