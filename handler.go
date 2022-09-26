package main

import "github.com/sirupsen/logrus"

func ErrHandler(err error) bool {
	if err != nil {
		logrus.Errorln(err)
		return false
	}
	return true
}
