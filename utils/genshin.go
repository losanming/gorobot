package utils

import (
	"math/rand"
	"time"
)

var dian int = 0       //出紫点数10
var dia *int = &dian   //指针 出紫点数
var DIAN int = 0       //出金进度77+
var DIA *int = &DIAN   //指针 出金点数
var chance int         //抽卡随机数
var cha *int = &chance //指针 抽卡随机数
var count int = 0      //目前抽卡数
var cou *int = &count  //指针 目前抽卡数
var sum int = 0        //总抽卡数
var su *int = &sum     //指针 总抽卡数
var baodi int = 0      //出金小保底 全局变量
func xiaobaodi() string {
	if baodi == 0 {
		var a int = rand.Intn(10)
		var b int = rand.Intn(len(Five_card))
		if a <= 4 {
			baodi = 0
			return Five_up[0]
		} else {
			baodi = 1
			return Five_card[b]
		}
	} else {
		baodi = 0
		return Five_up[0]
	}
}

// 概率提升函数
func gailvup(DIA *int, cha *int) {
	var a int = rand.Intn(10)
	switch *DIA {
	case 77, 78:
		if a == 1 {
			*DIA = 0
			*cha = 0
		} else {
			*cha = rand.Intn(1001)
			*DIA++
		}
	case 79:
		if a <= 2 {
			*DIA = 0
			*cha = 0
		} else {
			*cha = rand.Intn(1001)
			*DIA++
		}
	case 80:
		if a <= 3 {
			*DIA = 0
			*cha = 0
		} else {
			*cha = rand.Intn(1001)
			*DIA++
		}
	case 81:
		if a <= 4 {
			*DIA = 0
			*cha = 0
		} else {
			*cha = rand.Intn(1001)
			*DIA++
		}
	case 82:
		if a <= 5 {
			*DIA = 0
			*cha = 0
		} else {
			*cha = rand.Intn(1001)
			*DIA++
		}
	case 83:
		if a <= 6 {
			*DIA = 0
			*cha = 0
		} else {
			*cha = rand.Intn(1001)
			*DIA++
		}
	case 84:
		if a <= 7 {
			*DIA = 0
			*cha = 0
		} else {
			*cha = rand.Intn(1001)
			*DIA++
		}
	case 85:
		if a <= 8 {
			*DIA = 0
			*cha = 0
		} else {
			*cha = rand.Intn(1001)
			*DIA++
		}
	case 86:
		if a <= 9 {
			*DIA = 0
			*cha = 0
		} else {
			*cha = rand.Intn(1001)
			*DIA++
		}
	case 87:
		*DIA = 0
		*cha = 0
	}
}

// 单抽函数
func danchou(dia *int, DIA *int, cha *int) {
	if *dia == 9 {
		*dia = 0
		*DIA++
		*cha = 16 //十连必紫
	} else if *DIA >= 77 { //第77抽概率概率提升
		gailvup(DIA, cha)
		*dia++
	} else {
		*cha = rand.Intn(1001) //生成一千个随机数 0~159；160~289；290~1000
		*dia++
		*DIA++
	}
}

// 判定函数
func panding(cha *int, cou *int, su *int, DIA *int, dia *int) string {
	if *cha >= 0 && *cha < 6 {
		result := xiaobaodi()
		*cou = 0
		*DIA = 0
		return result
	} else if *cha >= 6 && *cha <= 57 {
		result := Sixin()
		*dia = 0
		return result
	} else if *cha > 58 && *cha <= 1000 {
		result := Sanxin()
		return result
	}
	return ""
}

func DrawCord(t int) []string {
	// 五星概率：1.60% RAIDENSHOGUN DILUC MONA QIQI JEAN KEQING
	// 四星概率：13.00%
	// 三星概率：85.40%
	rand.Seed(time.Now().Unix()) //作为随机数种子
	//

	var final_card []string
	if t == 1 {
		count++
		sum++
		danchou(dia, DIA, cha)
		result := panding(cha, cou, su, DIA, dia)
		final_card = append(final_card, result)
		return final_card
	} else if t == 2 {
		for a := 1; a <= 10; a++ {
			time.Sleep(time.Millisecond * 400)
			count++
			sum++
			danchou(dia, DIA, cha)
			result := panding(cha, cou, su, DIA, dia)
			final_card = append(final_card, result)
		}
		return final_card
	}
	return final_card
}
