package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func PrintFileLog(filename string, v ...interface{}) {
	file, err := os.OpenFile("./logs/"+filename+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		log.Println("log open filename:", filename, ", error:", err.Error())
		return
	}
	logger := log.New(file, "", log.LstdFlags)
	log.Println(fmt.Sprint(v...))
	logger.Println(fmt.Sprint(v...))
	file.Close()
}

func HandleErr(err error) {
	if err != nil {
		logrus.Errorf("err is : %s", err)
	}
}
