package module

import (
	"example.com/m/module/user"
	"example.com/m/module/ws"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	new(user.Controller).RegisterRoute(r)
	new(ws.Controller).RegisterRoute(r)
}
