package main

import (
	"example.com/m/utils"
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	c := cron.New()
	spec := "10 0/3 * * *"
	c.AddFunc(spec, Task)
	c.Start()
	select {}
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

func Task() {
	lists, err := utils.GetMenusListByFile()
	if err != nil {
		logrus.Errorln("get menus is fail err : ", err)
		return
	}
	menu_info := lists[utils.GetRandomIndex()]
	msg := fmt.Sprintf("今天吃什么呢？就是它了： %s,我们要怎么做呢？ %s", menu_info.Name, menu_info.Info)
	err = utils.SendMsgById(413944516, msg)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
}
