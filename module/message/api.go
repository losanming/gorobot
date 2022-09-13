package message

import (
	"example.com/m/module/db"
	"example.com/m/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

type Controller struct{}

// 注册路由
func (self Controller) RegisterRoute(r *gin.RouterGroup) {
	r.POST("/message/list", response.API(self.MessageHistory)) // 历史记录
}

func (self *Controller) MessageHistory(c *gin.Context) (data map[string]interface{}, err error) {
	var message db.Message
	p := c.PostForm("page")
	ps := c.PostForm("page_size")
	page, _ := strconv.Atoi(p)
	page_size, _ := strconv.Atoi(ps)
	result, err := message.GetMessageList(page, page_size)
	if err != nil {
		logrus.Error("getmessage is fail err: ", err)
		return data, err
	}
	data = make(map[string]interface{})
	data["result"] = result
	return data, err
}

