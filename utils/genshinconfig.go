package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var FourStat_up = []string{}

// 常驻
var FourPermannet = []string{}

var ThreePermanent = []string{}

var Five_up = []string{}

var Five_card = []string{}

type Card struct {
	FourStatUp     []string
	FourPermannet  []string
	ThreePermanent []string
	Five_up        []string
	Five_card      []string
}

func init() {
	var c Card
	open, err := os.Open("./config/genshin.json")
	if err != nil {
		fmt.Println("read file is failed")
		Exit()
	}
	b, err := ioutil.ReadAll(open)
	if err != nil {
		fmt.Println("err :", err)
		Exit()
	}
	json.Unmarshal(b, &c)

	FourStat_up = c.FourStatUp
	FourPermannet = c.FourPermannet
	ThreePermanent = c.ThreePermanent
	Five_up = c.Five_up
	Five_card = c.Five_card
}
