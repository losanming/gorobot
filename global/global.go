package global

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"os"
	"strconv"
	"time"
)

var (
	LOCALHOSTPORT = ""
	DAIBIAODAHUI  int64
	CRON          = ""
	// 处理协议
	BAIKE  = 1 // 维基百科  百科:
	TIANQI = 2 // 天气   天气:
	HELP   = 3 // 帮助信息
)

//API

var (
	MOYU   = "https://api.vvhan.com/api/moyu?type=json"
	DAY60S = "https://api.vvhan.com/api/60s?type=json"
)

// APIKEYS
const (
	FUTUREWEATHERKEY = "251518e073ef6c3c9504dd286c3f6a86"
)

func init() {
	_, err := os.Stat("./config/config.ini")
	if err != nil {
		panic("can't found config")
	}
	load, err := ini.Load("./config/config.ini")
	url := load.Section("common").Key("url").String()
	group_id := load.Section("common").Key("group_id").String()
	cron := load.Section("common").Key("cron").String()
	if url == "" || group_id == "" {
		logrus.Errorln("config.ini params is wrong")
		Exit()
	}
	LOCALHOSTPORT = url
	CRON = cron
	atoi, err := strconv.Atoi(group_id)
	if err != nil {
		fmt.Println("group_id atoi is failed")
		Exit()
	}

	DAIBIAODAHUI = int64(atoi)
	fmt.Println("config load is ok")
}

func Exit() {
	fmt.Println("5s 后关闭程序")
	time.Sleep(5 * time.Second)
	os.Exit(1)
}
