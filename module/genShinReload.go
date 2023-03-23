package module

import (
	"encoding/json"
	"example.com/m/utils"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

var (
	age           = 0
	init_physique = 0
	init_lucky    = 0
	init_wisdom   = 0
	real_physique = 0
	real_lucky    = 0
	real_wisdom   = 0
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

// Physique体质 lucky幸运 wisdom智慧
func Reload(Physique, lucky, wisdom int) (err error) {
	init_wisdom = wisdom
	init_lucky = lucky
	init_physique = Physique

	return err
}

func CheckEvent(event_id int) {
	switch event_id {
	case 114918:
		// 生成 start<= x <end 的随机数
		var num = utils.GenShinGetRandomIndex(50)
		age += int(num)
		break
	case 117418:
		age = 1
		break
	case 120718:
		age++
		break
	case 124918:
		var num = utils.GenShinGetRandomIndex(50)
		age += int(num)
		break
	default:
		return
	}
}

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

// 校验种族
func CheckRace(events_id int) string {
	if events_id%10000 == 12 {
		return "yuansu"
	} else if events_id%10000 == 11 {
		return "qiuqiu"
	}
	return ""
}

func BeginReload() {
	//获取种族

}

func LoadOneEvent() {
	if age < 1 || age > 100 {
		return
	}
	fmt.Printf("第 %v 月", age)
	//获取该月份随机事件

}
