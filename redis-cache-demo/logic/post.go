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

func GetPostById(id string) (data *model.Post, err error) {
	data, err = mysql.QueryPostById(id)
	if err != nil {
		log.Println("Mysql Query Post By Id ERROR:", err)
		return nil, err
	}
	return data, nil
}

func GetPostByPostId(postid string) (data *model.Post, err error) {
	data, err = mysql.QueryPostByPostId(postid)
	if err != nil {
		log.Println("mysql Query Post By Id", err)
		return nil, err
	}
	return data, nil
}

func UpdatePostById(postModel *model.Post, post model.Post) (err error) {
	err = mysql.UpdatePostById(postModel, post)
	return err
}

func DeletePostById(postid string) (err error) {
	err = mysql.DeletePostById(postid)
	if err != nil {
		log.Println("mysql Delete Post By Id", err)
		return err
	}
	return nil
}
