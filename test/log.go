package main

import (
	"blog/pkg/util"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"

	//"log"
	"os"
)

var Log = logrus.New()

//const Log =log

func test() {
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: util.TimeLayout,
	})
	path := `storage/logs/app.log`
	//fmt.Println(global.AppSetting.LogSavePath )
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil || os.IsNotExist(err) {
		file, err = os.Create(path)
		//return
	}
	fmt.Println(123)
	log.SetOutput(file)
}

func main() {
	//test()
	//logrus 测试~
	//log.WithFields(logrus.Fields{
	//	"animal": "walrus",
	//	"size":   10,
	//}).Info("A group of walrus emerges from the ocean")
	//Log.Error("error!!!")
}
