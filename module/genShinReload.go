package module

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

// 年龄行为
type Age struct {
	AgeEvents []Events `json:"age_events"`
}

type Events struct {
	Event []int `json:"event"`
}

// 日常行为
type DayEvent struct {
	Id    int    `json:"id"`
	Event string `json:"event"`
	Props []int  `json:"props"`
}

type DayEvents struct {
	Event []DayEvent `json:"day_events"`
}

// 死亡线
type DieEvennt struct {
	Element []int `json:"element"`
	QqMan   []int `json:"qq_man"`
}

var GenShinAge Age
var GenShinDayEvents DayEvents
var GenShinDieEvents DieEvennt

func init() {
	//加载事件线
	//加载死亡线、日常线和年龄线
	age_file, err := os.Open("./genShinReload/age.json")
	if err != nil {
		fmt.Println("open age file is fail ")
		return
	}
	age_tmp, err := ioutil.ReadAll(age_file)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	json.Unmarshal(age_tmp, &GenShinAge)
	//___
	day_event_file, err := os.Open("./genShinReload/dayEvent.json")
	if err != nil {
		fmt.Println("open day file is fail ")
		return
	}
	dayEvent_tmp, err := ioutil.ReadAll(day_event_file)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	json.Unmarshal(dayEvent_tmp, &GenShinDayEvents)

	//___
	die_event_file, err := os.Open("./genShinReload/dieEvent.json")
	if err != nil {
		fmt.Println("open die file is fail ")
		return
	}
	dieEvent_tmp, err := ioutil.ReadAll(die_event_file)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	json.Unmarshal(dieEvent_tmp, &GenShinDieEvents)

	if GenShinDayEvents.Event == nil || GenShinAge.AgeEvents == nil || GenShinDieEvents.Element == nil {
		logrus.Errorln("load genshinreload is fail")
		return
	}
	logrus.Info("load genshin is ok")
}
