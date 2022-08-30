package module

import (
	"github.com/gin-gonic/gin"
	"mytest/master/module/user"
)

func RegisterRoutes(r *gin.RouterGroup) {
	new(user.Controller).RegisterRoute(r)
}
