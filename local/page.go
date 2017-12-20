package main

import (
	"net/http"
)

func pushPublicIndexPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

func pushPrivateIndexPage(w http.ResponseWriter, r *http.Request) {
	l := loadPrivateIndexData(w, r)
	renderTemplate(w, "i.html", l)
}

func pushLoginPage(w http.ResponseWriter, r *http.Request) {
	sueecssLogin := verifyLoginStatus(r)
	if sueecssLogin == false {
		renderTemplate(w, "login.html", nil)
	} else {
		http.Redirect(w, r, "/", 303)
	}
}

func pushForgotPage(w http.ResponseWriter, r *http.Request) {

}

func pushRegisterPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "register.html", nil)
}

func pushConsolePage(w http.ResponseWriter, r *http.Request) {
}

func pushUploadPage(w http.ResponseWriter, r *http.Request) {
	loginStatus := verifyLoginStatus(r)
	if loginStatus == true {
		renderTemplate(w, "upload.html", nil)
	} else {
		http.Redirect(w, r, "/login", 303)
	}
}

func pushRecordPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "record.html", nil)
}

func pushRecordsPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "records.html", nil)
}

func pushLibraryPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "library.html", nil)
}

func pushGalleryPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "gallery.html", nil)
}
