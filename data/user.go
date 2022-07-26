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

func Create(u *User)error {
	db := initialization.Db
	nowTime := time.Now().Unix()
	info := &User{
		UserName: u.UserName,
		Password: u.Password,
		Account: u.Account,
		UpdateAt: nowTime,
		CreateAt: nowTime,
	}
	err := db.Create(info).Error
	return err
}

func GetUser(u *User) (*User, error) {
	user := &User{}
	db := initialization.Db
	err := db.Where("account = ?", u.Account).First(user).Error
	return user, err
}
