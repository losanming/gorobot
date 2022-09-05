package response

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
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

// Cors 开放所有接口的OPTIONS方法
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		//c.Header("Access-Control-Allow-Origin", "*")
		//c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		//c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		//c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		//c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
	}
}
