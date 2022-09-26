package service

import (
	"encoding/json"
	"example.com/m/global"
	"example.com/m/module"
	"example.com/m/utils"
	"fmt"
	"github.com/sirupsen/logrus"
)

type Wiki struct {
	Data WikiData `json:"data"`
}
type WikiData struct {
	Text string `json:"text"`
}

func BeginChannel(c chan []string) {
	var c_close bool
	for {
		if c_close {
			return
		}
		select {
		case split, ok := <-c:
			if !ok {
				c_close = true
			} else {
				if split[0] == "百科" {
					err, rs := GetWikiInfo(split[1])
					if err != nil {
						logrus.Errorln(err)
						return
					}
					err = module.SendMsgById(413944516, rs)
					if err != nil {
						logrus.Errorln(err)
					}
				} else if split[0] == "天气" {
					err := module.SendMsgById(413944516, "功能开发中")
					if err != nil {
						logrus.Errorln(err)
					}
				}
			}
		default:
			fmt.Println("input is wrong")
			return
		}
	}
}

func GetWikiInfo(find string) (err error, rs string) {
	var info Wiki
	url := global.WIKIURL + find
	resp, err := utils.SendRequest(url, nil, nil, "GET")
	if err != nil {
		return err, rs
	}
	_ = json.Unmarshal(resp, &info)
	rs = info.Data.Text
	return err, rs
}

func GetWeather(city string) {

}
