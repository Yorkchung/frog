package main

import (
	"regexp"
)

func isEnglish(str string) bool {
	isEnglish := false
	match, _ := regexp.MatchString(`^[a-zA-Z]+$`, str)
	if match == true {
		isEnglish = true
	}
	return isEnglish
}

func filterEmail(e string) bool {
	formatOK := false
	match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, e)
	if match == true {
		formatOK = true
	}
	return formatOK
}

func filterUsername(un string) bool {
	formatOK := false
	match, _ := regexp.MatchString("^[0-9a-zA-Z-]+$", un)
	if match == true {
		formatOK = true
	}
	return formatOK
}

func filterPassword(pw string) bool {
	formatOK := false
	match, _ := regexp.MatchString("^[0-9a-zA-Z]+$", pw)
	if match == true {
		formatOK = true
	}
	return formatOK
}

func filterHTML(text string) bool {
	formatOK := true
	return formatOK
}
