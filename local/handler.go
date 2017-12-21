package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	loginStatus := false
	r.ParseForm()
	loginField := r.FormValue("account")
	pw := r.FormValue("password")
	isMail := classifyLoginField(loginField)
	//fmt.Println(loginField, pw)
	if isMail == true {

	} else {
		ok := verifyPasswordByUsername(loginField, pw)
		if ok == true {
			storeSession(w, r, loginField, pw)
			loginStatus = true
		}
	}
	p := LoginPage{LoginStatus: loginStatus}
	b, _ := json.Marshal(p)
	w.Write(b)
}

func logout(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/", 303)
}

func forgot(w http.ResponseWriter, r *http.Request) {

}

func register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	em := r.FormValue("email")
	un := r.FormValue("username")
	pw := r.FormValue("password")
	emailFormatOK := filterEmail(em)
	usernameFormatOK := filterUsername(un)
	passwordFormatOK := filterPassword(pw)
	catchFalse(emailFormatOK, "register email format err")
	catchFalse(usernameFormatOK, "register userinfo format err")
	catchFalse(passwordFormatOK, "register password format err")
	storeAccount(em, un, pw)
	success := storeSession(w, r, un, pw)
	if success == true {
		http.Redirect(w, r, "/", 303)
	}
}

func uploadRecord(w http.ResponseWriter, r *http.Request) {
	cu, _ := r.Cookie("username")
	unFormatOK := filterUsername(cu.Value)
	if unFormatOK == true {
		UID, getUIDOK := getUIDByUsername(cu.Value)
		catchFalse(getUIDOK, "get uid by username err")
		storeRecord(w, r, UID)
	}
}

func getRecord(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	recordID := r.FormValue("recordid")
	record := searchRecordByRecordID(recordID)
	b, _ := json.Marshal(record)
	w.Write(b)
}

func getAllRecords(w http.ResponseWriter, r *http.Request) {
	record := searchAllRecords()
	b, _ := json.Marshal(record)
	w.Write(b)
}

func getRecordsByTag(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	tag := r.FormValue("tag")
	records := searchRecordsByTag(tag)
	b, _ := json.Marshal(records)
	w.Write(b)
}

func getRecordsByKeyword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	searchtype := r.FormValue("searchtype")
	keyword := r.FormValue("keyword")
	switch searchtype {
	case "tag":
		records := searchRecordsByTag(keyword)
		b, _ := json.Marshal(records)
		w.Write(b)
	case "organismname":
		records := searchRecordsByOrganismName(keyword)
		b, _ := json.Marshal(records)
		w.Write(b)
	case "address":
		records := searchRecordsByAddress(keyword)
		b, _ := json.Marshal(records)
		w.Write(b)
	case "gps":
		longitude := r.FormValue("longitude")
		latitude := r.FormValue("latitude")
		records := searchRecordsByGPS(longitude, latitude)
		b, _ := json.Marshal(records)
		w.Write(b)
	case "season":
		records := searchRecordsBySeason(keyword)
		b, _ := json.Marshal(records)
		w.Write(b)
	case "daterange":
		datefrom := r.FormValue("datefrom")
		dateto := r.FormValue("dateto")
		records := searchRecordsByDateRange(datefrom, dateto)
		fmt.Println(records)
		b, _ := json.Marshal(records)
		w.Write(b)
	}
}

func getRecordByRecordID(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	record := searchRecordByRecordID(r.Form.Get("recordid"))
	b, _ := json.Marshal(record)
	w.Write(b)
}

func modifyRecordByRecordID(w http.ResponseWriter, r *http.Request) {
	successUpdate := false
	successUpdate = alterRecordByRecordID(r)
	p := UploadPage{UploadStatus: successUpdate}
	b, _ := json.Marshal(p)
	w.Write(b)
}

func modifyRecordPhotosByRecordID(w http.ResponseWriter, r *http.Request) {
	l := LibraryPage{}
	b, _ := json.Marshal(l)
	w.Write(b)
}

func deleteRecordByRecordID(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	removeRecordByRecordID(r.FormValue("recordid"))
	// need return json tall ok or not, and ajax reload
}

func deleteRecordPhotosByPhotoID() {

}

func getLibraryData(w http.ResponseWriter, r *http.Request) {
	records := searchLibraryData()
	b, _ := json.Marshal(records)
	w.Write(b)
}

func getLibraryByLabel(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	label := r.FormValue("label")
	records := searchLibraryDataByLabel(label)
	b, _ := json.Marshal(records)
	w.Write(b)
}

func uploadLibraryWithCSV(w http.ResponseWriter, r *http.Request) {
	successUpload := false
	cu, _ := r.Cookie("username")
	unFormatOK := filterUsername(cu.Value)
	if unFormatOK == true {
		UID, getUIDOK := getUIDByUsername(cu.Value)
		catchFalse(getUIDOK, "get uid by username err")
		successUpload = storeLibraryWithCSV(r, UID)
	}
	p := UploadPage{UploadStatus: successUpload}
	b, _ := json.Marshal(p)
	w.Write(b)
}

func modifyLibraryData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("modifyLibraryData")
}

func deleteLibraryData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleteLibraryData")
}

func getGalleryByKeyword(w http.ResponseWriter, r *http.Request) {

}

func getPhoto(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	photoID := r.FormValue("photoid")
	photo := searchPhotoByPhotoID(photoID)
	b, _ := json.Marshal(photo)
	w.Write(b)
}
