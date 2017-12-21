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
	mux.HandleFunc("/", indexController)
	mux.HandleFunc("/login", loginController)
	mux.HandleFunc("/logout", logoutController)
	mux.HandleFunc("/forgot", forgotController)
	mux.HandleFunc("/register", registerController)
	mux.HandleFunc("/console", consoleController)
	mux.HandleFunc("/upload", uploadController)
	mux.HandleFunc("/record", recordController)
	mux.HandleFunc("/record/data", recordDataController)
	mux.HandleFunc("/records", recordsController)
	mux.HandleFunc("/records/data", recordsDataController)
	mux.HandleFunc("/library", libraryController)
	mux.HandleFunc("/library/data", libraryDataController)
	mux.HandleFunc("/gallery", galleryController)
	mux.HandleFunc("/gallery/data", galleryDataController)

	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(frogConfig.PublicPath))))
	mux.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir(frogConfig.ResourcePath))))
	mux.Handle("/storage/", http.StripPrefix("/storage/", http.FileServer(http.Dir(frogConfig.StoragePath))))

	err := http.ListenAndServe(":80", mux)
	checkErr(err, "ListenAndServe err")
}
