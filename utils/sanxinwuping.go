package utils

import (
	"math/rand"
	"time"
)

func Sanxin() string {
	rand.Seed(time.Now().UnixNano()) //生成随机数种子
	var a int = rand.Intn(22)
	return ThreePermanent[a]
}
