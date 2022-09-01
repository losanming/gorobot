package websocket

import "github.com/gin-gonic/gin"

type ServerInterface interface {
	Run(gin *gin.Context)
}
