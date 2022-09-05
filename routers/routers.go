package routers

import (
	"example.com/m/config/global"
	"example.com/m/module"
	"example.com/m/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Run() {
	//初始化路由
	r := initRouters()
	//启动服务
	server := &http.Server{
		Addr:           ":8199",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   100 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logrus.Panicf("panic er is : %s", err)
	}

}

func initRouters() *gin.Engine {
	r := gin.New()
	//打印和异常处理
	r.Use(gin.Logger())
	r.Use(global.Recovery())
	r.Use(response.Cors())
	//load file
	r.Static("../static", "static")
	r.LoadHTMLGlob("view/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	//路由分配
	//先不做拦截器
	routerVersion := r.Group("")
	module.RegisterRoutes(routerVersion)
	routerVersion.GET("/home", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})
	return r
}
