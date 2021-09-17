package data

import (
	"fmt"
	"goTestProject/Personal-Blog/initialization"
	"time"
)

type User struct {
	Password string
	Account  string
	UserName string
	UpdateAt int64
	CreateAt int64
}

func (User) TableName() string {
	return "user"
}

func Create(u *User) {
	db := initialization.Db
	nowTime := time.Now().Unix()
	info := &User{
		UserName: u.UserName,
		Password: u.Password,
		Account: u.Account,
		UpdateAt: nowTime,
		CreateAt: nowTime,
	}
	if err := db.Create(info).Error; err != nil {
		fmt.Println("err:", err)
	}
}

func GetUser(u *User) *User {
	user := &User{}
	db := initialization.Db
	// TODO 报错为解决
	if err := db.Where("UserName = ?", u.UserName).First(user); err != nil {
		fmt.Println("err:", *err)
	}
	return user
}
