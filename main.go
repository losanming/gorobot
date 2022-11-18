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
	"runtime"
	"strings"
	"time"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		logrus.Errorln("request method is not post")
		return
	}
	resp, err := ioutil.ReadAll(r.Body)
	if !ErrHandlers(err) {
		return
	}
	find := gojsonq.New().FromString(string(resp)).Find("message_type")
	if find == nil {
		return
	}
	find_group := gojsonq.New().FromString(string(resp)).Find("group_id")
	if find_group == nil {
		return
	}
	if find.(string) != "group" || int64(find_group.(float64)) != global.DAIBIAODAHUI {
		logrus.Println("find_type: ", find.(string), " group: ", find_group.(float64))
		return
	}
	// 获取消息
	find_msg := gojsonq.New().FromString(string(resp)).Find("raw_message")
	if find_msg == nil {
		logrus.Errorln("find_msg result is nil")
		return
	}
	split := strings.Split(find_msg.(string), "=")
	if len(split) != 2 {
		logrus.Errorln("split len is wrong", split)
		return
	}
	//通道和协程处理后面写,后面要拆开实现
	go GotoSendMsg(split[0], split[1])
}

func main() {
	go BeginTask()
	logrus.Println("定时任务开始")
	go Gc()
	logrus.Info("GC定时任务开始")
	http.HandleFunc("/", ServeHTTP)
	logrus.Println("监听开始")
	//监听httpserver
	server := &http.Server{Addr: ":5701"}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err: ", err)
		} else {
			defer server.Close()
		}
	}()
	err := server.ListenAndServe()
	if !ErrHandlers(err) {
		return
	}
}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	writer2 := os.Stdout
	logrus.SetLevel(logrus.InfoLevel)
	_, err := os.Stat("./logs")
	if err != nil {
		err := os.MkdirAll(fmt.Sprintf("./logs"), 0766)
		if err != nil {
			panic("mkdir is failer ")
		}
	}
	writer3, err := os.OpenFile("./logs/server.log", os.O_WRONLY|os.O_CREATE,
		0777)
	if err != nil {
		return
	}
	logrus.SetOutput(io.MultiWriter(writer2, writer3))
	fmt.Println("log load is ok")
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
	spec := global.CRON
	c.AddFunc(spec, Task)
	c.Start()
	select {}
}

func GotoSendMsg(key, value string) {
	if key == "百科" {
		err, rs := service.GetWikiInfo(value)
		if err != nil {
			logrus.Errorln(err)
			return
		}
		err = module.SendMsgById(int64(global.DAIBIAODAHUI), rs)
		if err != nil {
			logrus.Errorln(err)
		}
	} else if key == "天气" {
		err2, result := service.GetWeather(value)
		if err2 != nil {
			logrus.Errorln("err: ", err2)
			return
		}
		message := fmt.Sprintf("今天是%s,%s,%s今天%s,最高温度%s,最低温度%s,实时温度%s,%s,大风等级是%s,风速%s。空气质量%s,  出门建议:%s", result.Date, result.Week, result.City, result.Wea,
			result.Tem1, result.Tem2, result.Tem, result.Win, result.WinSpeed, result.WinSpeed, result.Aqi.AirLevel, result.Aqi.AirTips)
		err := module.SendMsgById(global.DAIBIAODAHUI, message)
		if err != nil {
			logrus.Errorln(err)
		}
	} else if key == "原神抽卡" {
		if value == "单抽" {
			result := utils.DrawCord(1)
			if result == nil {
				return
			}
			err := module.SendMsgById(global.DAIBIAODAHUI, result[0])
			if err != nil {
				logrus.Errorln(err)
			}
		} else if value == "十连抽" {
			result := utils.DrawCord(2)
			if result == nil {
				return
			}
			rs := strings.Join(result, ",")
			err := module.SendMsgById(global.DAIBIAODAHUI, rs)
			if err != nil {
				logrus.Errorln(err)
			}
		}
	} else if key == "原神角色" {

	}
}

func Gc() {
	for {
		go func() {
			runtime.GC()
			logrus.Info("begin GC", time.Now())
		}()
		time.Sleep(90 * time.Minute)
	}
}

func ErrHandlers(err error) bool {
	if err != nil {
		logrus.Errorln(err)
		return false
	}
	return true
}
