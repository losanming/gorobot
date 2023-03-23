package service

import (
	"encoding/json"
	"example.com/m/global"
	"example.com/m/module"
	"example.com/m/utils"
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type Day60SData struct {
	ImgUrl string `json:"imgUrl"`
}

type MoYuData struct {
	ImgUrl string `json:"url"`
}

func Day60S() {
	var d Day60SData
	resp, err := utils.SendRequest(global.DAY60S, nil, nil, "GET")
	if err != nil {
		return
	}
	_ = json.Unmarshal(resp, &d)
	if d.ImgUrl == "" {
		return
	}
	f := fmt.Sprintf("[CQ:image,file=%s]", d.ImgUrl)
	err = module.SendMsgById(global.DAIBIAODAHUI, f, false)
	fmt.Println("Err : ", err)
}

func MoYu() {
	var d MoYuData
	resp, err := utils.SendRequest(global.MOYU, nil, nil, "GET")
	if err != nil {
		return
	}
	_ = json.Unmarshal(resp, &d)
	if d.ImgUrl == "" {
		return
	}
	f := fmt.Sprintf("[CQ:image,file=%s]", d.ImgUrl)
	err = module.SendMsgById(global.DAIBIAODAHUI, f, false)
	fmt.Println("Err : ", err)
}

func Meum() {
	lists, err := utils.GetMenusListByFile()
	if err != nil {
		logrus.Errorln("get menus is fail err : ", err)
		return
	}
	menu_info := lists[utils.GetRandomIndex()]
	msg := fmt.Sprintf("今天吃什么呢？就是它了： %s,我们要怎么做呢？ %s", menu_info.Name, menu_info.Info)
	err = module.SendMsgById(global.DAIBIAODAHUI, msg, false)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
}

func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

func BeginTask() {
	c := newWithSeconds()
	spec := "0 */3 * * * ?"
	c.AddFunc(spec, Day60S)
	//c.AddFunc(spec, MoYu)
	c.Start()
}
