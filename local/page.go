package main

import (
	"net/http"
)

func initDynamic(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		initIndex(w, r)
	} else if r.URL.Path[0:9] == "/library/" {
		initLibraryPage(w, r)
	} else if r.URL.Path[0:7] == "/photo/" {
		initGalleryPage(w, r)
	} else if r.URL.Path[0:8] == "/record/" {
		initRecordPage(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func initIndex(w http.ResponseWriter, r *http.Request) {
	loginStatus := verifyLoginStatus(r)
	if loginStatus == true {
		if r.Method == "GET" {
			l := loadPrivateIndexData(w, r)
			renderTemplate(w, "i.html", l)
		} else if r.Method == "POST" {

		}
	} else if loginStatus == false {
		if r.Method == "GET" {
			renderTemplate(w, "index.html", nil)
		}
	}
}

func initLoginPage(w http.ResponseWriter, r *http.Request) {
	sueecssLogin := verifyLoginStatus(r)
	if sueecssLogin == false {
		renderTemplate(w, "login.html", nil)
	} else {
		http.Redirect(w, r, "/", 303)
	}
}

func initForgotPage(w http.ResponseWriter, r *http.Request) {

}

func initRegisterPage(w http.ResponseWriter, r *http.Request) {
	logined := verifyLoginStatus(r)
	if logined == false {
		renderTemplate(w, "register.html", nil)
	} else {
		http.Redirect(w, r, "/", 303)
	}
}

func initConsolePage(w http.ResponseWriter, r *http.Request) {
}

func initUploadPage(w http.ResponseWriter, r *http.Request) {
	loginStatus := verifyLoginStatus(r)
	if loginStatus == true {
		renderTemplate(w, "upload.html", nil)
	} else {
		http.Redirect(w, r, "/login", 303)
	}
}

func initLibraryPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "library.html", nil)
}

func initGalleryPage(w http.ResponseWriter, r *http.Request) {

}

func initRecordsPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "records.html", nil)
}

func initRecordPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "record.html", nil)
}

func initTestPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "test.html", nil)
}

func test(w http.ResponseWriter, r *http.Request) {

}
