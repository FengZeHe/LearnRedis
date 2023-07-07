package logic

import (
	"fmt"
	"redisdemo/model"
	"redisdemo/mysql"
)

// 获取全部的用户
func GetAllUsers() (data []*model.User, err error) {
	data, err = mysql.QueryAllUsers()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 通过用户id查找用户信息
func GetUserById(userid string) (data *model.User, err error) {
	data, err = mysql.QueryUserById(userid)
	if err != nil {
		return nil, err
	}

	return data, err
}

// 通过用户id删除信息
func DeleteUserById(userid string) (err error) {
	err = mysql.DeleteUserById(userid)
	if err != nil {
		fmt.Println("Mysql delete user  error", err)
	}
	return err
}

// 通过用户id修改邮箱
func UpdateUserById(userid string, email string) (err error) {
	err = mysql.UpdateUserById(userid, email)
	if err != nil {
		return err
	}
	return err
}
