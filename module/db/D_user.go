package db

type User struct {
	Id       uint   `json:"id" gorm:"primary_key"`
	UserName string `json:"user_name" gorm:"varchar(64)"`
	PassWord string `json:"pass_word" gorm:"varchar(64)"`
}
