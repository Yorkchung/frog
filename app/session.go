package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/garyburd/redigo/redis"
)

func newHashPassword(password string) string {
	hash := sha512.New()
	hash.Write([]byte("galagala"))
	bytedPW := hash.Sum([]byte(password))
	stringPW := hex.EncodeToString(bytedPW)
	return stringPW
}

func newHashSession(un string) string {
	hashSess := ""
	username := sha256.New()
	username.Write([]byte("ungogo2017"))
	byteHashSess := username.Sum([]byte(un))
	hashSess = hex.EncodeToString(byteHashSess)
	return hashSess
}

func setCookie(cookieName string, cookieValue string, y int, m int, d int, w http.ResponseWriter, r *http.Request) bool {
	successCreate := false
	expiration := time.Now()
	expiration = expiration.AddDate(y, m, d)
	c := http.Cookie{Name: cookieName, Value: cookieValue, Path: "/", Domain: "importfmt.blog", Expires: expiration}
	http.SetCookie(w, &c)
	successCreate = true
	return successCreate
}

func storeAccount(e, u, p string) {
	pw := newHashPassword(p)
	_, storeAccountDataErr := db.Exec("INSERT INTO userinfo SET email=?, username=?, password=?, createtime=CURRENT_TIMESTAMP", e, u, pw)
	checkErr(storeAccountDataErr, "storeAccountDataErr")
}

func storeSession(w http.ResponseWriter, r *http.Request, un string, pw string) bool {
	success := false
	hashSess := ""
	hashSess = newHashSession(un)
	rdb.Do("HSET", un, "session", hashSess)
	sess, getSessionErr := redis.String(rdb.Do("HGET", un, "session"))
	checkErr(getSessionErr, "get session from redis err")
	if sess == hashSess {
		setCookie("username", un, 0, 0, 1, w, r)
		setCookie("sess", hashSess, 0, 0, 1, w, r)
		success = true
	}
	return success
}

func getSession(un string) (bool, string) {
	exist := false
	sess, getSessionErr := redis.String(rdb.Do("HGET", un, "session"))
	checkErr(getSessionErr, "get session from redis err")
	if getSessionErr == nil {
		exist = true
	}
	return exist, sess
}

func clearSession(w http.ResponseWriter) {
	log.Println("logout()")
	cs := &http.Cookie{
		Name:   "sess",
		Value:  "",
		Path:   "/",
		Domain: "importfmt.blog",
		MaxAge: -1,
	}
	http.SetCookie(w, cs)
	cl := &http.Cookie{
		Name:   "loginstatus",
		Value:  "",
		Path:   "/",
		Domain: "importfmt.blog",
		MaxAge: -1,
	}
	http.SetCookie(w, cl)
	cuid := &http.Cookie{
		Name:   "userid",
		Value:  "",
		Path:   "/",
		Domain: "importfmt.blog",
		MaxAge: -1,
	}
	http.SetCookie(w, cuid)
	cu := &http.Cookie{
		Name:   "username",
		Value:  "",
		Path:   "/",
		Domain: "importfmt.blog",
		MaxAge: -1,
	}
	http.SetCookie(w, cu)
}

func classifyLoginField(loginfield string) bool {
	isMail := false
	m, classifyLoginFieldErr := regexp.MatchString("@", loginfield)
	checkInfo(classifyLoginFieldErr, "classify login field err")
	if m == true {
		isMail = true
	}
	return isMail
}

func verifyLoginStatus(r *http.Request) bool {
	loginStatus := false
	cu, getCookieValueErr := r.Cookie("username")
	checkInfo(getCookieValueErr, "get username from cookie err")
	if getCookieValueErr == nil {
		clientSess := newHashSession(cu.Value)
		serverSessExist, serverSess := getSession(cu.Value)
		if serverSessExist == true {
			if clientSess == serverSess {
				loginStatus = true
			}
		}
	}
	return loginStatus
}

func verifyAdminStatus() bool {
	adminStatus := false
	return adminStatus
}

func verifyPasswordByUsername(username, password string) bool {
	truePassword := false
	serverHashPW, exist := getPasswordByUsername(username)
	if exist == true {
		clientHashPW := newHashPassword(password)
		if clientHashPW == serverHashPW {
			truePassword = true
		}
	}
	return truePassword
}

func verifyPasswordByMail(email, password string) bool {
	truePassword := false

	return truePassword
}

func verifyAccount() {

}

func verifyUsername() {

}

func verifyEmail() {

}
