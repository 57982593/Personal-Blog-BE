package data

import (
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

func InsertUser(account string, password string) error  {
	db := initialization.Db
	nowTime := time.Now().Unix()
	info := &User{
		UserName: account,
		Password: password,
		Account: "",
		UpdateAt: nowTime,
		CreateAt: nowTime,
	}
	err := db.Create(info).Error
	return err
}
