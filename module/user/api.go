package user

import (
	"github.com/gin-gonic/gin"
	"mytest/master/utils/response"
)

type Controller struct{}

// 注册路由
func (self Controller) RegisterRoute(r *gin.RouterGroup) {
	r.POST("/user/WebLogin", response.API(self.WebLogin)) // 后台 - 账号密码登录
}

func (self *Controller) WebLogin(c *gin.Context) (data map[string]interface{}, err error) {
	return data, nil
}
