package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func login(w http.ResponseWriter, r *http.Request) {
	loginStatus := false
	loginStatus = verifyLoginStatus(r)
	if loginStatus == false {
		if r.Method == "POST" {
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

func requestAllRecords(w http.ResponseWriter, r *http.Request) {
	record := getAllRecords()
	b, _ := json.Marshal(record)
	w.Write(b)
}

func searchRecordsByKeyword(w http.ResponseWriter, r *http.Request) {
	logined := verifyLoginStatus(r)
	if r.Method == "POST" && logined == true {
		r.ParseForm()
		//fmt.Println(r.Form)
		searchtype := r.FormValue("searchtype")
		keyword := r.FormValue("keyword")
		switch searchtype {
		case "recordname":
			records := getRecordsByRecordName(keyword)
			b, _ := json.Marshal(records)
			w.Write(b)
		case "organismname":
			records := getRecordsByOrganismName(keyword)
			b, _ := json.Marshal(records)
			w.Write(b)
		case "tag":
			records := getRecordsByTag(keyword)
			b, _ := json.Marshal(records)
			w.Write(b)
		case "address":
			records := getRecordsByAddress(keyword)
			b, _ := json.Marshal(records)
			w.Write(b)
		case "gps":
			longitude := r.FormValue("longitude")
			latitude := r.FormValue("latitude")
			records := getRecordsByGPS(longitude, latitude)
			b, _ := json.Marshal(records)
			w.Write(b)
		case "season":
			records := getRecordsBySeason(keyword)
			b, _ := json.Marshal(records)
			w.Write(b)
		case "daterange":
			datefrom := r.FormValue("datefrom")
			dateto := r.FormValue("dateto")
			records := getRecordsByDateRange(datefrom, dateto)
			fmt.Println(records)
			b, _ := json.Marshal(records)
			w.Write(b)
		}
	}
}

func searchRecordsByOrganismName(w http.ResponseWriter, r *http.Request) {
	loginStatus := verifyLoginStatus(r)
	if loginStatus == true {
		organismName := r.FormValue("organismname")
		records := getRecordsByOrganismName(organismName)
		b, _ := json.Marshal(records)
		w.Write(b)
	}
}

func searchRecordByRecordID(w http.ResponseWriter, r *http.Request) {
	loginStatus := verifyLoginStatus(r)
	if loginStatus == true {
		if r.Method == "POST" {
			r.ParseForm()
			record := getRecordByRecordID(r.Form.Get("recordid"))
			b, _ := json.Marshal(record)
			w.Write(b)
		}
	}
}

func searchLibrary(w http.ResponseWriter, r *http.Request) {
	records := getLibraryData()
	b, _ := json.Marshal(records)
	w.Write(b)
}

func searchLibraryByLabel(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	label := r.FormValue("label")
	records := getLibraryDataByLabel(label)
	b, _ := json.Marshal(records)
	w.Write(b)
}

/*
func searchLibraryBySpecies(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	species := r.FormValue("species")
	records := getLibraryDataBySpecies(species)
	b, _ := json.Marshal(records)
	w.Write(b)
}
*/

func uploadLibraryWithCSV(w http.ResponseWriter, r *http.Request) {
	loginStatus := verifyLoginStatus(r)
	if loginStatus == true {
		if r.Method == "POST" {
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
	}
}

func modifyRecordByRecordID(w http.ResponseWriter, r *http.Request) {
	loginStatus := verifyLoginStatus(r)
	if loginStatus == true {
		if r.Method == "POST" {
			successUpdate := false

			successUpdate = alterRecordByRecordID(r)

			p := UploadPage{UploadStatus: successUpdate}
			b, _ := json.Marshal(p)
			w.Write(b)
		}
	}
}

func modifyRecordPhotosByRecordID(w http.ResponseWriter, r *http.Request) {
	loginStatus := verifyLoginStatus(r)
	if loginStatus == true {
		if r.Method == "POST" {
			l := LibraryPage{}
			b, _ := json.Marshal(l)
			w.Write(b)
		}
	}
}

func modifyLibraryData(w http.ResponseWriter, r *http.Request) {
	loginStatus := verifyLoginStatus(r)
	if loginStatus == true {
		if r.Method == "POST" {
			fmt.Println("modifyLibraryData")
		}
	}
}

func deleteLibraryData(w http.ResponseWriter, r *http.Request) {
	loginStatus := verifyLoginStatus(r)
	if loginStatus == true {
		if r.Method == "POST" {
			fmt.Println("deleteLibraryData")
		}
	}
}

func deleteRecordByRecordID(w http.ResponseWriter, r *http.Request) {
	loginStatus := verifyLoginStatus(r)
	if loginStatus == true {
		if r.Method == "POST" {
			r.ParseForm()
			removeRecordByRecordID(r.Form.Get("recordid"))
			// need return json tall ok or not, and ajax reload
		}
	}
}

func deleteRecordPhotosByPhotoID() {

}

func parseCoordinateString(val string) float64 {
	chunks := strings.Split(val, ",")
	hours, _ := strconv.ParseFloat(strings.TrimSpace(chunks[0]), 64)
	minutes, _ := strconv.ParseFloat(strings.TrimSpace(chunks[1]), 64)
	seconds, _ := strconv.ParseFloat(strings.TrimSpace(chunks[2]), 64)
	return hours + (minutes / 60) + (seconds / 3600)
}

func parseCoordinate(latitudeValue, latitudePosition, longitudeValue, longitudePosition string) (string, string) {
	lati := parseCoordinateString(latitudeValue)
	long := parseCoordinateString(longitudeValue)

	if latitudePosition == "S" {
		lati *= -1
	}

	if longitudePosition == "W" {
		long *= -1
	}
	la := strconv.FormatFloat(lati, 'f', 6, 64)
	lo := strconv.FormatFloat(long, 'f', 6, 64)
	return la, lo
}
