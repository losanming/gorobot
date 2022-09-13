package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 注册和登陆时都需要保存session信息
func SaveAuthSession(c *gin.Context, info interface{}) {
	session := sessions.Default(c)
	session.Set("uid", info)
	// c.SetCookie("user_id",string(info.(map[string]interface{})["b"].(uint)), 1000, "/", "localhost", false, true)
	session.Save()
}

// 退出时清除session
func ClearAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func AuthSessionMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionValue := session.Get("uid")
		if sessionValue == nil {
			c.Redirect(http.StatusFound, "/")
			return
		}

		uidInt, _ := strconv.Atoi(sessionValue.(string))

		if uidInt <= 0 {
			c.Redirect(http.StatusFound, "/")
			return
		}

		// 设置简单的变量
		c.Set("uid", sessionValue)

		c.Next()
		return
	}
}
