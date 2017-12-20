package main

import (
	"net/http"
)

func indexController(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		loginStatus := verifyLoginStatus(r)
		if loginStatus == true {
			if r.Method == "GET" {
				pushPrivateIndexPage(w, r)
			}
		} else {
			if r.Method == "GET" {
				pushPublicIndexPage(w, r)
			}
		}
	}
}

func loginController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		pushLoginPage(w, r)
	case "POST":
		login(w, r)
	}
}

func logoutController(w http.ResponseWriter, r *http.Request) {
	logout(w, r)
}

func forgotController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		pushForgotPage(w, r)
	case "POST":
		loginStatus := verifyLoginStatus(r)
		if loginStatus == false {
			forgot(w, r)
		}
	}
}

func registerController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		loginStatus := verifyLoginStatus(r)
		if loginStatus == false {
			pushRegisterPage(w, r)
		} else {
			http.Redirect(w, r, "/", 303)
		}
	case "POST":
		loginStatus := verifyLoginStatus(r)
		if loginStatus == false {
			register(w, r)
		}
	}
}

func consoleController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		loginStatus := verifyLoginStatus(r)
		if loginStatus == true {
			pushConsolePage(w, r)
		} else {
			http.Redirect(w, r, "/", 303)
		}
	case "POST":
	}
}

func uploadController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		loginStatus := verifyLoginStatus(r)
		if loginStatus == true {
			pushUploadPage(w, r)
		} else {
			http.Redirect(w, r, "/login", 303)
		}
	}
}

func recordController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		pushRecordPage(w, r)
	case "POST":
		loginStatus := verifyLoginStatus(r)
		if loginStatus == true {
			uploadRecord(w, r)
		}
	}
}

func recordsController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		pushRecordsPage(w, r)
	case "POST":
		loginStatus := verifyLoginStatus(r)
		if loginStatus == true {

		}
	}
}

func libraryController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		pushLibraryPage(w, r)
	case "POST":
		searchLibraryByLabel(w, r)
	}
}

func galleryController(w http.ResponseWriter, r *http.Request) {

}
