package service

import (
	"encoding/json"
	"example.com/m/utils"
	"fmt"
	url2 "net/url"
)

type Wiki struct {
	Data WikiData `json:"data"`
}
type WikiData struct {
	Text string `json:"text"`
}

type WeatherRequest struct {
	City string `json:"city"`
	Ip   string `json:"ip"`
}
type WeatherResponse struct {
	Date       string `json:"date"`        //日期
	Week       string `json:"week"`        //星期
	UpdateTime string `json:"update_time"` //气象台更新时间
	City       string `json:"city"`        //城市
	Country    string `json:"country"`     //国家
	Wea        string `json:"wea"`         //天气
	Tem        string `json:"tem"`         //实时温度
	Tem1       string `json:"tem1"`        //高温
	Tem2       string `json:"tem2"`        //低温
	Win        string `json:"win"`         //风向
	WinSpeed   string `json:"win_speed"`   //风力
	WinMeter   string `json:"win_meter"`   //风速
	Humidity   string `json:"humidity"`    //湿度
	Visibility string `json:"visibility"`  //能见度
	Pressure   string `json:"pressure"`    //气压
	Aqi        struct {
		UpdateTime string `json:"update_time"`
		Air        string `json:"air"`       //空气质量
		AirLevel   string `json:"air_level"` // 质量等级
		AirTips    string `json:"air_tips"`  //空气质量描述
		Kouzhao    string `json:"kouzhao"`
		Yundong    string `json:"yundong"`
		Waichu     string `json:"waichu"`
		Kaichuang  string `json:"kaichuang"`
		Jinghuaqi  string `json:"jinghuaqi"`
	} `json:"aqi"`
	Code int `json:"code"`
}

func GetWikiInfo(find string) (err error, rs string) {
	var info Wiki
	escape := url2.QueryEscape(find)
	url := fmt.Sprintf("https://api.wer.plus/api/dub?t=%s", escape)
	resp, err := utils.SendRequest(url, nil, nil, "GET")
	if err != nil {
		return err, rs
	}
	_ = json.Unmarshal(resp, &info)
	rs = info.Data.Text
	return err, rs
}

func GetWeather(city string) (err error, result WeatherResponse) {
	escape := url2.QueryEscape(city)
	url := fmt.Sprintf("https://api.wer.plus/api/tian?city=%s", escape)
	resp, err := utils.SendRequest(url, nil, nil, "GET")
	if err != nil {
		return err, result
	}
	_ = json.Unmarshal(resp, &result)
	return err, result
}
