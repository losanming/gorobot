package db

import (
	"example.com/m/config/global"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id       uint   `json:"id" gorm:"primary_key"`
	UserName string `json:"user_name" gorm:"varchar(64)"`
	PassWord string `json:"pass_word" gorm:"text(256)"`
}

type UserRegister struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

type UserLogin struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

func (u *User) Create() error {
	db := global.GDB
	defer db.Close()
	err := db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&u).Error
	})
	return err
}

func (u User) FindUserInfoByUserName() (result User, err error) {
	if u.UserName == "" {
		return result, err
	}

	db := global.GDB
	defer db.Close()
	err = db.Model(&User{}).Where("user_name = ? ", u.UserName).Scan(&result).Error
	return result, err
}
