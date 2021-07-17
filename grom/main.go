package main

import (
	"fmt"
	database "goTestProject/init"
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

func main() {
	var users []User
	db := database.GetDb()
	db.Find(&users)
	fmt.Println(users)
}
