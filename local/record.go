package main

import "github.com/importfmt/logger"

func checkErr(err error, message string) {
	if err != nil {
		logger.Error.Println(message)
		//fmt.Println(err)
	}
}

func checkWarn(err error, message string) {
	if err != nil {
		logger.Warning.Println(message)
		//fmt.Println(err)
	}
}

func checkInfo(err error, message string) {
	if err != nil {
		logger.Info.Println(message)
		//fmt.Println(err)
	}
}

func catchFalse(b bool, message string) {
	if b != true {
		logger.Warning.Println(message)
	}
}

func catchTrue(b bool, message string) {
	if b != false {
		logger.Warning.Println(message)
	}
}
