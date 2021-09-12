package data

import (
	"fmt"
	database "goTestProject/initialization"
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

var Test = 1

func (User) TableName() string {
	return "user"
}

func Get(id int64) *User {
	fmt.Print(Test)
	user := &User{}
	db := database.GetDb()
	if err := db.Where("id = ?", id).First(user).Error; err != nil {
		fmt.Println("查询出错了")
	}
	return user
}
func GetUserList() *[]User {
	db := database.Db
	users := []User{}
	if err := db.Find(&users).Error; err != nil {
		_ = fmt.Errorf("get user list fail")
	}
	return &users
}

// func (u *User) GetUserInfo(id int64) error {
// 	db := database.GetDb()
// 	return db.Where("id = ?", id).First(u).Error
// }
