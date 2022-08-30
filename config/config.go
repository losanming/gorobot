package config

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"runtime"
)

var goos, exepath, exedir, runpath string

func GetEnv() {
	// 输出环境日志
	goos = runtime.GOOS
	exepath, _ = filepath.Abs(os.Args[0])
	exedir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	runpath, _ = os.Getwd()
	logrus.Info("操作系统：", goos)

	logrus.Info("操作系统：", goos)
	logrus.Info("执行程序路径：", exepath)
	logrus.Info("执行程序文件路径：", exedir)
	logrus.Info("运行目录路径：", runpath)
}
