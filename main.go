package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"mytest/cqhttpServer/module"
	"os"
)

func main() {
	list, err := module.GetGroupList()
	if err != nil {
		logrus.Errorln(err)
	}
	fmt.Println(list)
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
	//server := &http.Server{
	//	Addr:         ":5700",
	//	ReadTimeout:  5 * time.Second,
	//	WriteTimeout: 5 * time.Second,
	//}
	//err := server.ListenAndServe()
	//if err != nil {
	//	fmt.Println("err : ", err)
	//	return
	//}
}
