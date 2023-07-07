package mysql

import (
	"fmt"
	"redisdemo/model"
)

// Mysql查询所有的用户的数据
func QueryAllUsers() (user []*model.User, err error) {
	db.Find(&user)
	fmt.Println(user)
	return user, nil
}

// Mysql根据用户id查询信息
func QueryUserById(userid string) (user *model.User, err error) {
	result := db.Where("user_id = ?", userid).Find(&user)
	if result.Error != nil {
		return nil, err
	}
	return user, nil
}

// Mysql根据id删除用户
func DeleteUserById(userid string) (err error) {
	var user *model.User
	result := db.Where("user_id = ?", userid).Delete(&user)
	if result.Error != nil {
		fmt.Println("删除用户失败", err)
		return err
	}
	return nil
}

func UpdateUserById(userid string, email string) (err error) {
	var user *model.User
	result := db.Where("user_id = ?", userid).Find(&user)
	if result.Error != nil {
		fmt.Println("query user error:", err)
		return err
	}

	result = db.Model(&user).UpdateColumn("email", email)
	if result.Error != nil {
		fmt.Println("update column error:", err)
		return err
	}

	return nil
}
