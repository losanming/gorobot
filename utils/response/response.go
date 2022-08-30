package response

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"regexp"
	"strings"
)

type CB func(c *gin.Context) (data map[string]interface{}, err error)

func API(cb CB) gin.HandlerFunc {

	return func(c *gin.Context) {
		data, err := cb(c)
		if data == nil {
			data = make(map[string]interface{}, 0)
		}
		if err != nil {
			msg := err.Error()
			code := "400"

			reg := regexp.MustCompile(`^__([\d]+?)__`)
			regList := reg.FindStringSubmatch(msg)
			if len(regList) == 2 {
				code = regList[1]
				msg = strings.Replace(msg, regList[0], "", 1)
			}

			if gorm.IsRecordNotFoundError(err) {
				msg = "对象不存在"
				code = "200"
			}

			c.JSON(200, gin.H{
				"code": code,
				"msg":  msg,
				"data": data,
			})
			return
		}

		c.JSON(200, gin.H{
			"code": "200",
			"msg":  "success",
			"data": data,
		})
	}
}
