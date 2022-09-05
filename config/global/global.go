package global

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

const (
	INFO  = "info"
	DBURL = "root:root@tcp(127.0.0.1:3306)/chat?charset=utf8&parseTime=True&loc=Local"
)

var (
	GDB *gorm.DB
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logrus.Panicln(err)
			}
		}()
		c.Next()
	}
}
