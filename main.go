package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"mytest/master/config"
	"mytest/master/module"
	"mytest/master/routers"
	"os"
)

func main() {
	//获取环境输出
	config.GetEnv()
	//加载数据库

	err := module.InitDB()
	if err != nil {
		logrus.Panicf("panic is : %s ", err)
	}

	routers.Run()

}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	writer2 := os.Stdout
	logrus.SetLevel(logrus.InfoLevel)
	writer3, err := os.OpenFile("./logs/server.log", os.O_WRONLY|os.O_CREATE,
		0777)
	if err != nil {
		return
	}
	logrus.SetOutput(io.MultiWriter(writer2, writer3))
}
