package main

import (
	"database/sql"
	"net/http"

	"github.com/importfmt/config"
	"github.com/importfmt/logger"

	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
)

var (
	frogConfig   *config.Config
	db           *sql.DB
	rdb          redis.Conn
	initMySQLErr error
	initRedisErr error
)

func main() {
	frogConfig = config.NewConfig()
	frogConfig.Init("/config/config.json")

	logger.Init(frogConfig.LogPath)
	logger.Info.Println("logger init")

	initTemplate(frogConfig.TemplatePath)

	db, initMySQLErr = sql.Open("mysql", frogConfig.MySQLUsername+":"+frogConfig.MySQLPassword+"@/"+frogConfig.MySQLDatabase)
	checkErr(initMySQLErr, "connectMySQLErr")
	defer db.Close()

	rdb, initRedisErr = redis.Dial("tcp", "localhost:6379")
	checkErr(initRedisErr, "connectRedisErr")
	defer rdb.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/", initDynamic)
	mux.HandleFunc("/login", initLoginPage)
	mux.HandleFunc("/forgot", initForgotPage)
	mux.HandleFunc("/register", initRegisterPage)
	mux.HandleFunc("/console", initConsolePage)
	mux.HandleFunc("/records", initRecordsPage)
	mux.HandleFunc("/upload", initUploadPage)

	mux.HandleFunc("/test", initTestPage)
	mux.HandleFunc("/posttest", test)

	mux.HandleFunc("/requestregister", register)
	mux.HandleFunc("/requestforgot", forgot)
	mux.HandleFunc("/requestlogin", login)
	mux.HandleFunc("/logout", logout)

	mux.HandleFunc("/upload-record", uploadRecord)
	mux.HandleFunc("/upload-library-with-csv", uploadLibraryWithCSV)

	mux.HandleFunc("/modify-library-data", modifyLibraryData)
	mux.HandleFunc("/delete-library-data", deleteLibraryData)

	mux.HandleFunc("/request-all-records", requestAllRecords)
	mux.HandleFunc("/search-records-by-keyword", searchRecordsByKeyword)
	mux.HandleFunc("/search-record-by-record-id", searchRecordByRecordID)
	mux.HandleFunc("/search-records-specify-organismname", searchRecordsByOrganismName)

	mux.HandleFunc("/search-library", searchLibrary)
	mux.HandleFunc("/search-library-by-label", searchLibraryByLabel)

	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(frogConfig.PublicPath))))
	mux.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir(frogConfig.ResourcePath))))
	mux.Handle("/storage/", http.StripPrefix("/storage/", http.FileServer(http.Dir(frogConfig.StoragePath))))

	/*
		if _, err := os.Stat(frogConfig.StoragePath + "photo/"); os.IsNotExist(err) {
			os.Mkdir(frogConfig.StoragePath+"photo/", 777)
		}
	*/

	err := http.ListenAndServe(":80", mux)
	checkErr(err, "ListenAndServe err")
}
