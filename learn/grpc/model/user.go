package model

import (
	database "goTestProject/Personal-Blog/initialization"
)

type User struct {
	Id               int64
	Account          string // qq 微信 唯一标识
	Way              string // 登录方式 qq=QQ | wx=微信
	Name             string
	Email            string
	Phone            string
	IsNotify         int64
	AgreementVersion string // 用户协议版本
	CreateAt         int64
	UpdateAt         int64
	LastLoginAt      int64
}

func (User) TableName() string {
	return "user"
}

func (u *User) GetUserInfo(id int64) error {
	db := database.GetDb()
	return db.Where("id = ?", id).First(u).Error
}

func (u User) GetUserList(offset, limit int64) ([]User, int64) {
	db := database.GetDb()
	var total int64
	var users []User
	db.Model(&User{}).Count(&total)
	db.Offset(offset).Limit(limit).Find(&users)
	return users, total
}
func DeleteUserInfo(userId int64) error {
	db := database.GetDb()
	return db.Delete(&User{}, userId).Error
}
