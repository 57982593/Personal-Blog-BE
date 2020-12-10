package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
)

var db *gorm.DB
var NotFoundErr = gorm.ErrRecordNotFound

func InitDatabase() {
	host := viper.Get("database.mysql.host")
	user := viper.Get("database.mysql.user")
	password := viper.GetString("database.mysql.password")
	name := viper.GetString("database.mysql.name")
	charset := viper.GetString("database.mysql.charset")
	isDevelopment := viper.GetBool("isDevelopment")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", user, password, host, name, charset)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect mysql %s", err.Error()))
	} else {
		db.DB().SetMaxIdleConns(viper.GetInt("database.mysql.pool.min"))
		db.DB().SetMaxOpenConns(viper.GetInt("database.mysql.pool.max"))
		if isDevelopment {
			db.LogMode(true)
		}
	}
}

