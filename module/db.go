package module

import (
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
	"mytest/master/config/global"
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

	return nil
}
