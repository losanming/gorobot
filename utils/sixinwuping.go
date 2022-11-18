package utils //作为主函数里四星物品的包

import (
	"math/rand"
	"time"
)

// 全局变量
var gaiup int

// 通用卡池
func tongyong(a int) string {
	return FourPermannet[a]
}

// up卡池
func up(a int) string {
	return FourStat_up[a]
}

func Sixin() string {
	rand.Seed(time.Now().UnixNano()) //生成随机数种子
	var a int
	var b int = rand.Intn(2)
	if gaiup == 1 || b == 0 {
		a = rand.Intn(3)
		result := up(a)
		gaiup = 0
		return result
	} else if gaiup == 0 {
		a = rand.Intn(len(FourPermannet))
		result := tongyong(a)
		gaiup = 1
		if a == 20 || a == 21 || a == 27 {
			gaiup = 0
		}
		return result
	}
	return ""
}
