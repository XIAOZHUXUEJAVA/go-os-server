package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type User struct {
	Id       int    `gorm:"type:int(11) auto_increment; primary_key; comment: '用户编号'" json:"user_id"  `
	Username string `gorm:"type:varchar(255); comment: '用户名';" json:"user_name"  `
	Password string `gorm:"type:varchar(255); comment: '密码';" json:"user_password"  `
}

// CreateUser 创建用户并保存到数据库
func CreateUser(user *User) error {
	err := db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUserByUsername 根据用户名查询用户
func GetUserByUsername(username string) (*User, error) {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	fmt.Println(user.Username)
	return &user, nil
}

// CheckUserPassword 检查用户密码是否匹配
func CheckUserPassword(user *User, password string) bool {
	return user.Password == password
}
