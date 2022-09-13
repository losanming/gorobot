package db

import (
	"example.com/m/config/global"
	"time"
)

type Message struct {
	Id         int       `json:"id" gorm:"primary_key"`
	UserId     int       `json:"user_id"`
	Content    string    `json:"content" gorm:"type:varchar(128)"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (m Message) GetMessageList(page, pagesize int) (result []Message, err error) {
	db := global.GDB
	err = db.Model(&Message{}).Limit(pagesize).Offset(page - 1).Scan(&result).Error
	if err != nil {
		return result, err
	}
	return result, err
}

func (m *Message) CreateMessage() (err error) {
	db := global.GDB
	return db.Create(&m).Error
}

func (m *Message) DeleteMessageById(err error) {
	db := global.GDB
	err = db.Delete(&m).Error
	if err != nil {
		return
	}
	return
}
