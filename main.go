package main

import (
	"example.com/m/global"
	"example.com/m/module"
	"example.com/m/service"
	"example.com/m/utils"
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/thedevsaddam/gojsonq"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := global.C
	c = make(chan []string)
	if r.Method != "POST" {
		logrus.Errorln("request method is not post")
		return
	}
	resp, err := ioutil.ReadAll(r.Body)
	if !ErrHandler(err) {
		return
	}
	find := gojsonq.New().FromString(string(resp)).Find("message_type")
	if find == nil {
		logrus.Errorln("find result is nil")
		return
	}
	find_group := gojsonq.New().FromString(string(resp)).Find("group_id")
	if find_group == nil {
		logrus.Errorln("find_group result is nil")
		return
	}
	if find.(string) != "group" || find_group.(int64) != 413944516 {
		logrus.Println("find_type: ", find.(string), " group: ", find_group.(int64))
		return
	}
	// 获取消息
	find_msg := gojsonq.New().FromString(string(resp)).Find("raw_message")
	if find_msg == nil {
		logrus.Errorln("find_msg result is nil")
		return
	}
	split := strings.Split(find_msg.(string), ":")
	if len(split) != 2 {
		logrus.Errorln("split len is wrong")
		return
	}
	c <- split
	service.BeginChannel(c)
	fmt.Println("input channel channel")
	close(c)
}

func main() {
	http.HandleFunc("/", ServeHTTP)

	//监听httpserver
	server := &http.Server{Addr: ":5701"}
	defer server.Close()
	err := server.ListenAndServe()
	if !ErrHandler(err) {
		return
	}
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

func Task() {
	lists, err := utils.GetMenusListByFile()
	if err != nil {
		logrus.Errorln("get menus is fail err : ", err)
		return
	}
	menu_info := lists[utils.GetRandomIndex()]
	msg := fmt.Sprintf("今天吃什么呢？就是它了： %s,我们要怎么做呢？ %s", menu_info.Name, menu_info.Info)
	err = module.SendMsgById(413944516, msg)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
}

func BeginTask() {
	c := cron.New()
	spec := "10 0/3 * * *"
	c.AddFunc(spec, Task)
	c.Start()
	select {}
}
