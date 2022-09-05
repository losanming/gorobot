package module

import (
	"example.com/m/config/global"
	db2 "example.com/m/module/db"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
	"time"
)

var (
	DB *gorm.DB
)

func InitDB() error {

	DB, err := gorm.Open("mysql", global.DBURL)
	if err != nil {
		return err
	}

	global.GDB = DB
	DB.LogMode(true)
	db := global.GDB.DB()
	if err != nil {
		return err
	}
	db.SetMaxIdleConns(1024)
	db.SetMaxOpenConns(256)
	db.SetConnMaxLifetime(time.Hour)

	//@@TODO 生成表格
	DB.CreateTable(&db2.User{})
	return nil
}
