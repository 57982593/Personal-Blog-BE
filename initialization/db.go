package initialization

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
)

var Db *gorm.DB

func Mysql() {
	var err error
	host := viper.Get("database.mysql.host")
	fmt.Println(host)
	user := viper.Get("database.mysql.user")
	password := viper.GetString("database.mysql.password")
	name := viper.GetString("database.mysql.name")
	charset := viper.GetString("database.mysql.charset")
	isDevelopment := viper.GetBool("isDevelopment")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", user, password, host, name, charset)
	Db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect mysql %s", err.Error()))
	} else {
		Db.DB().SetMaxIdleConns(viper.GetInt("database.mysql.pool.min"))
		Db.DB().SetMaxOpenConns(viper.GetInt("database.mysql.pool.max"))
		if isDevelopment {
			Db.LogMode(true)
		}
	}
}

func GetDb() *gorm.DB {
	return Db
}

