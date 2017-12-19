package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xiam/exif"
)

var (
	randomCharacterTable = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
)

func newRandomString(length int) string {
	randString := ""
	var buffer bytes.Buffer
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		num := r.Intn(35)
		buffer.WriteString(randomCharacterTable[num])
	}
	randString = buffer.String()
	return randString
}

func getHashedPWByUsername(un string) (string, bool) {
	pw := ""
	exist := false
	err := db.QueryRow("SELECT password FROM userinfo WHERE username = ?", un).Scan(&pw)
	checkInfo(err, "select password from userinfo err")
	if err == nil {
		exist = true
	}
	return pw, exist
}

func getUIDByUsername(un string) (string, bool) {
	UID := ""
	exist := false
	err := db.QueryRow("SELECT id FROM userinfo WHERE username = ?", un).Scan(&UID)
	checkInfo(err, "select uid from userinfo err")
	if err == nil {
		exist = true
	}
	return UID, exist
}

func searchUsernameByUsername(un string) (string, bool) {
	var u string
	exist := false
	err := db.QueryRow("SELECT username FROM userinfo WHERE username = ?", un).Scan(&u)
	checkErr(err, "can not get username")
	if err == nil {
		exist = true
	}
	return u, exist
}

func getRecordsByRecordName(recordName string) Records {
	fmt.Println(recordName)
	recordIDs, records := []int{}, Records{}
	records.Records = make(map[int]Record)
	idrows, queryErr := db.Query("SELECT id FROM record WHERE recordname=?", recordName)
	checkErr(queryErr, "query organis id from comment with mysql error")
	defer idrows.Close()
	for idrows.Next() {
		var tmp int
		scanErr := idrows.Scan(&tmp)
		checkErr(scanErr, "scan organis id from comment with mysql error")
		recordIDs = append(recordIDs, tmp)
	}

	for index, id := range recordIDs {
		recordname, organismname, food, stage, season, status, habitat, note, createTime := "", "", "", "", "", "", "", "", ""
		db.QueryRow("SELECT recordname FROM record WHERE id = ?", id).Scan(&recordname)
		db.QueryRow("SELECT organismname FROM record WHERE id = ?", id).Scan(&organismname)
		db.QueryRow("SELECT food FROM record WHERE id = ?", id).Scan(&food)
		db.QueryRow("SELECT stage FROM record WHERE id = ?", id).Scan(&stage)
		db.QueryRow("SELECT season FROM record WHERE id = ?", id).Scan(&season)
		db.QueryRow("SELECT status FROM record WHERE id = ?", id).Scan(&status)
		db.QueryRow("SELECT habitat FROM record WHERE id = ?", id).Scan(&habitat)
		db.QueryRow("SELECT note FROM record WHERE id = ?", id).Scan(&note)
		db.QueryRow("SELECT createtime FROM record WHERE id = ?", id).Scan(&createTime)

		r := Record{
			ID:           id,
			RecordName:   recordname,
			OrganismName: organismname,
			Food:         food,
			Stage:        stage,
			Season:       season,
			Status:       status,
			Habitat:      habitat,
			Note:         note,
			CrateTime:    createTime}
		r.PhotoSrc = make(map[int]string)
		pathRows, queryErr := db.Query("SELECT path FROM photo WHERE recordid = ?", id)
		checkErr(queryErr, "query photo path from comment with mysql error")
		defer pathRows.Close()
		i := 0
		for pathRows.Next() {
			var tmp string
			scanErr := pathRows.Scan(&tmp)
			checkErr(scanErr, "scan photo path from comment with mysql error")

			r.PhotoSrc[i] = tmp
			i++
		}

		r.PhotoLatitude = make(map[int]string)
		latitudeRows, queryErr := db.Query("SELECT latitude FROM photo WHERE recordid = ?", id)
		checkErr(queryErr, "query photo latitude from comment with mysql error")
		defer latitudeRows.Close()
		j := 0
		for latitudeRows.Next() {
			var tmp string
			scanErr := latitudeRows.Scan(&tmp)
			checkErr(scanErr, "scan photo latitude from comment with mysql error")
			r.PhotoLatitude[i] = tmp
			j++
		}

		r.PhotoLongitude = make(map[int]string)
		longitudeRows, queryErr := db.Query("SELECT longitude FROM photo WHERE recordid = ?", id)
		checkErr(queryErr, "query photo longitude from comment with mysql error")
		defer longitudeRows.Close()
		k := 0
		for longitudeRows.Next() {
			var tmp string
			scanErr := longitudeRows.Scan(&tmp)
			checkErr(scanErr, "scan photo longitude from comment with mysql error")
			r.PhotoLongitude[i] = tmp
			k++
		}
		records.Records[index] = r
	}
	return records
}

func getRecordsByOrganismName(organismName string) Records {
	fmt.Println(organismName)
	recordIDs, records := []int{}, Records{}
	records.Records = make(map[int]Record)
	idrows, queryErr := db.Query("SELECT id FROM record WHERE organismname=?", organismName)
	checkErr(queryErr, "query organis id from comment with mysql error")
	defer idrows.Close()
	for idrows.Next() {
		var tmp int
		scanErr := idrows.Scan(&tmp)
		checkErr(scanErr, "scan organis id from comment with mysql error")
		recordIDs = append(recordIDs, tmp)
	}

	for index, id := range recordIDs {
		recordname, organismname, food, stage, season, status, habitat, note := "", "", "", "", "", "", "", ""
		db.QueryRow("SELECT recordname FROM record WHERE id = ?", id).Scan(&recordname)
		db.QueryRow("SELECT organismname FROM record WHERE id = ?", id).Scan(&organismname)
		db.QueryRow("SELECT food FROM record WHERE id = ?", id).Scan(&food)
		db.QueryRow("SELECT stage FROM record WHERE id = ?", id).Scan(&stage)
		db.QueryRow("SELECT season FROM record WHERE id = ?", id).Scan(&season)
		db.QueryRow("SELECT status FROM record WHERE id = ?", id).Scan(&status)
		db.QueryRow("SELECT habitat FROM record WHERE id = ?", id).Scan(&habitat)
		db.QueryRow("SELECT note FROM record WHERE id = ?", id).Scan(&note)

		r := Record{
			ID:           id,
			RecordName:   recordname,
			OrganismName: organismname,
			Food:         food,
			Stage:        stage,
			Season:       season,
			Status:       status,
			Habitat:      habitat,
			Note:         note}
		r.PhotoSrc = make(map[int]string)
		pathRows, queryErr := db.Query("SELECT path FROM photo WHERE recordid = ?", id)
		checkErr(queryErr, "query photo path from comment with mysql error")
		defer pathRows.Close()
		i := 0
		for pathRows.Next() {
			var tmp string
			scanErr := pathRows.Scan(&tmp)
			checkErr(scanErr, "scan photo path from comment with mysql error")
			r.PhotoSrc[i] = tmp
			i++
		}

		r.PhotoLatitude = make(map[int]string)
		latitudeRows, queryErr := db.Query("SELECT latitude FROM photo WHERE recordid = ?", id)
		checkErr(queryErr, "query photo latitude from comment with mysql error")
		defer latitudeRows.Close()
		j := 0
		for latitudeRows.Next() {
			var tmp string
			scanErr := latitudeRows.Scan(&tmp)
			checkErr(scanErr, "scan photo latitude from comment with mysql error")
			r.PhotoLatitude[i] = tmp
			j++
		}

		r.PhotoLongitude = make(map[int]string)
		longitudeRows, queryErr := db.Query("SELECT longitude FROM photo WHERE recordid = ?", id)
		checkErr(queryErr, "query photo longitude from comment with mysql error")
		defer longitudeRows.Close()
		k := 0
		for longitudeRows.Next() {
			var tmp string
			scanErr := longitudeRows.Scan(&tmp)
			checkErr(scanErr, "scan photo longitude from comment with mysql error")
			r.PhotoLongitude[i] = tmp
			k++
		}
		records.Records[index] = r
	}
	return records
}

func getRecordsByTag(tag string) Records {
	fmt.Println(tag)
	recordIDs, records := []int{}, Records{}
	records.Records = make(map[int]Record)
	idrows, queryErr := db.Query("SELECT id FROM record WHERE tag=?", tag)
	checkErr(queryErr, "query records id from comment with mysql error")
	defer idrows.Close()
	for idrows.Next() {
		var tmp int
		scanErr := idrows.Scan(&tmp)
		checkErr(scanErr, "scan records id from comment with mysql error")
		recordIDs = append(recordIDs, tmp)
	}

	for index, id := range recordIDs {
		recordname, organismname, food, stage, season, status, habitat, note := "", "", "", "", "", "", "", ""
		db.QueryRow("SELECT recordname FROM record WHERE id = ?", id).Scan(&recordname)
		db.QueryRow("SELECT organismname FROM record WHERE id = ?", id).Scan(&organismname)
		db.QueryRow("SELECT food FROM record WHERE id = ?", id).Scan(&food)
		db.QueryRow("SELECT stage FROM record WHERE id = ?", id).Scan(&stage)
		db.QueryRow("SELECT season FROM record WHERE id = ?", id).Scan(&season)
		db.QueryRow("SELECT status FROM record WHERE id = ?", id).Scan(&status)
		db.QueryRow("SELECT habitat FROM record WHERE id = ?", id).Scan(&habitat)
		db.QueryRow("SELECT note FROM record WHERE id = ?", id).Scan(&note)

		r := Record{
			ID:           id,
			RecordName:   recordname,
			OrganismName: organismname,
			Food:         food,
			Stage:        stage,
			Season:       season,
			Status:       status,
			Habitat:      habitat,
			Note:         note}
		r.PhotoSrc = make(map[int]string)
		idrows, queryErr := db.Query("SELECT path FROM photo WHERE recordid = ?", id)
		checkErr(queryErr, "query photo path from comment with mysql error")
		defer idrows.Close()
		i := 0
		for idrows.Next() {
			var tmp string
			scanErr := idrows.Scan(&tmp)
			checkErr(scanErr, "scan photo path from comment with mysql error")
			r.PhotoSrc[i] = tmp
			i++
		}
		records.Records[index] = r
	}

	return records
}

func getLibraryDataByLabel(label string) Library {
	dataIDs, library := []int{}, Library{}
	library.LibraryDatas = make(map[int]LibraryData)

	idrows, queryErr := db.Query("SELECT id FROM library WHERE label=?", label)
	checkErr(queryErr, "query library id from comment with mysql error")
	defer idrows.Close()
	for idrows.Next() {
		var tmp int
		scanErr := idrows.Scan(&tmp)
		checkErr(scanErr, "scan library id from comment with mysql error")
		dataIDs = append(dataIDs, tmp)
	}

	for index, id := range dataIDs {
		organismname, family, food, season, status, habitat, note := "", "", "", "", "", "", ""
		db.QueryRow("SELECT organismname FROM library WHERE id = ?", id).Scan(&organismname)
		db.QueryRow("SELECT family FROM library WHERE id = ?", id).Scan(&family)
		db.QueryRow("SELECT food FROM library WHERE id = ?", id).Scan(&food)
		db.QueryRow("SELECT season FROM library WHERE id = ?", id).Scan(&season)
		db.QueryRow("SELECT status FROM library WHERE id = ?", id).Scan(&status)
		db.QueryRow("SELECT habitat FROM library WHERE id = ?", id).Scan(&habitat)
		db.QueryRow("SELECT note FROM library WHERE id = ?", id).Scan(&note)

		ld := LibraryData{
			ID:           id,
			OrganismName: organismname,
			Family:       family,
			Food:         food,
			Season:       season,
			Status:       status,
			Habitat:      habitat,
			Note:         note}
		library.LibraryDatas[index] = ld
	}
	return library
}

func getLibraryDataBySpecies(species string) Library {
	dataIDs, library := []int{}, Library{}
	library.LibraryDatas = make(map[int]LibraryData)

	idrows, queryErr := db.Query("SELECT id FROM library WHERE species=?", species)
	checkErr(queryErr, "query library id from comment with mysql error")
	defer idrows.Close()
	for idrows.Next() {
		var tmp int
		scanErr := idrows.Scan(&tmp)
		checkErr(scanErr, "scan library id from comment with mysql error")
		dataIDs = append(dataIDs, tmp)
	}

	for index, id := range dataIDs {
		organismname, family, food, season, status, habitat, note := "", "", "", "", "", "", ""
		db.QueryRow("SELECT organismname FROM library WHERE id = ?", id).Scan(&organismname)
		db.QueryRow("SELECT family FROM library WHERE id = ?", id).Scan(&family)
		db.QueryRow("SELECT food FROM library WHERE id = ?", id).Scan(&food)
		db.QueryRow("SELECT season FROM library WHERE id = ?", id).Scan(&season)
		db.QueryRow("SELECT status FROM library WHERE id = ?", id).Scan(&status)
		db.QueryRow("SELECT habitat FROM library WHERE id = ?", id).Scan(&habitat)
		db.QueryRow("SELECT note FROM library WHERE id = ?", id).Scan(&note)

		ld := LibraryData{
			ID:           id,
			OrganismName: organismname,
			Family:       family,
			Food:         food,
			Season:       season,
			Status:       status,
			Habitat:      habitat,
			Note:         note}
		library.LibraryDatas[index] = ld
	}
	return library
}

func getRecordsByAddress(address string) Records {
	records := Records{}
	return records
}

func getRecordsByGPS(longitude, latitude string) Records {
	records := Records{}
	return records
}

func getRecordsBySeason(season string) Records {
	fmt.Println(season)
	recordIDs, records := []int{}, Records{}
	records.Records = make(map[int]Record)
	idrows, queryErr := db.Query("SELECT id FROM record WHERE season=?", season)
	checkErr(queryErr, "query record id from comment with mysql error")
	defer idrows.Close()
	for idrows.Next() {
		var tmp int
		scanErr := idrows.Scan(&tmp)
		checkErr(scanErr, "scan record id from comment with mysql error")
		recordIDs = append(recordIDs, tmp)
	}

	for index, id := range recordIDs {
		recordname, organismname, food, stage, season, status, habitat, note := "", "", "", "", "", "", "", ""
		db.QueryRow("SELECT recordname FROM record WHERE id = ?", id).Scan(&recordname)
		db.QueryRow("SELECT organismname FROM record WHERE id = ?", id).Scan(&organismname)
		db.QueryRow("SELECT food FROM record WHERE id = ?", id).Scan(&food)
		db.QueryRow("SELECT stage FROM record WHERE id = ?", id).Scan(&stage)
		db.QueryRow("SELECT season FROM record WHERE id = ?", id).Scan(&season)
		db.QueryRow("SELECT status FROM record WHERE id = ?", id).Scan(&status)
		db.QueryRow("SELECT habitat FROM record WHERE id = ?", id).Scan(&habitat)
		db.QueryRow("SELECT note FROM record WHERE id = ?", id).Scan(&note)

		r := Record{
			ID:           id,
			RecordName:   recordname,
			OrganismName: organismname,
			Food:         food,
			Stage:        stage,
			Season:       season,
			Status:       status,
			Habitat:      habitat,
			Note:         note}
		r.PhotoSrc = make(map[int]string)
		idrows, queryErr := db.Query("SELECT path FROM photo WHERE recordid = ?", id)
		checkErr(queryErr, "query photo path from comment with mysql error")
		defer idrows.Close()
		i := 0
		for idrows.Next() {
			var tmp string
			scanErr := idrows.Scan(&tmp)
			checkErr(scanErr, "scan photo path from comment with mysql error")
			r.PhotoSrc[i] = tmp
			i++
		}
		records.Records[index] = r
	}
	return records
}

func getRecordsByDateRange(dateFrom, dateTo string) Records {
	fmt.Println(dateFrom, dateTo)
	recordIDs, records := []int{}, Records{}
	records.Records = make(map[int]Record)

	idrows, queryErr := db.Query("SELECT id FROM record WHERE createtime BETWEEN " + dateFrom + " AND " + dateTo + "")
	checkErr(queryErr, "query record id from comment with mysql error")
	defer idrows.Close()
	for idrows.Next() {
		var tmp int
		scanErr := idrows.Scan(&tmp)
		checkErr(scanErr, "scan record id from comment with mysql error")
		recordIDs = append(recordIDs, tmp)
	}

	for index, id := range recordIDs {
		recordname, organismname, food, stage, season, status, habitat, note := "", "", "", "", "", "", "", ""
		db.QueryRow("SELECT recordname FROM record WHERE id = ?", id).Scan(&recordname)
		db.QueryRow("SELECT organismname FROM record WHERE id = ?", id).Scan(&organismname)
		db.QueryRow("SELECT food FROM record WHERE id = ?", id).Scan(&food)
		db.QueryRow("SELECT stage FROM record WHERE id = ?", id).Scan(&stage)
		db.QueryRow("SELECT season FROM record WHERE id = ?", id).Scan(&season)
		db.QueryRow("SELECT status FROM record WHERE id = ?", id).Scan(&status)
		db.QueryRow("SELECT habitat FROM record WHERE id = ?", id).Scan(&habitat)
		db.QueryRow("SELECT note FROM record WHERE id = ?", id).Scan(&note)

		r := Record{
			ID:           id,
			RecordName:   recordname,
			OrganismName: organismname,
			Food:         food,
			Stage:        stage,
			Season:       season,
			Status:       status,
			Habitat:      habitat,
			Note:         note}
		r.PhotoSrc = make(map[int]string)
		idrows, queryErr := db.Query("SELECT path FROM photo WHERE recordid = ?", id)
		checkErr(queryErr, "query photo path from comment with mysql error")
		defer idrows.Close()
		i := 0
		for idrows.Next() {
			var tmp string
			scanErr := idrows.Scan(&tmp)
			checkErr(scanErr, "scan photo path from comment with mysql error")
			r.PhotoSrc[i] = tmp
			i++
		}
		records.Records[index] = r
	}
	return records
}

func checkOrganismNameExistByOrganismName(organismName string) bool {
	var o string
	exist := false
	err := db.QueryRow("SELECT organismname FROM record WHERE organismname = ?", organismName).Scan(&o)
	checkErr(err, "can not get username")
	if err == nil && o != "" {
		exist = true
	}
	return exist
}

func storeLibraryWithCSV(r *http.Request, UID string) bool {
	successStoreRecord := false
	/*
		尚未實作：
		解析表單
		儲存csv檔
		開啟csv檔案
		分析並記錄到 MySQL
	*/
	csvPath := "蝶蛾.csv"
	file, err := os.Open(frogConfig.StoragePath + "csv/" + csvPath)
	checkErr(err, "open csv file err")
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("read", csvPath, err)
		}

		result, storeRecordErr := db.Exec("INSERT INTO library SET organismname=?, label=?, species=?, status=?, habitat=?,createtime=CURRENT_TIMESTAMP", record[2], "lepidoptera", record[2], record[3], record[4])
		fmt.Println(storeRecordErr)
		if storeRecordErr == nil {
			id, getRecordIDErr := result.LastInsertId()
			if getRecordIDErr == nil {
				recordID := strconv.FormatInt(id, 10)
				if record[0] != "" {
					updateFamilyCommand := "UPDATE library SET `family`=?" + " WHERE id=?"
					db.Exec(updateFamilyCommand, record[0], recordID) // 科
				}
				/*
					if record[1] != "" {
						updateGenusCommand := "UPDATE library SET `genus`=?" + " WHERE id=?"
						db.Exec(updateGenusCommand, record[1], recordID) // 屬
					}
				*/
			}
		}

	}
	return successStoreRecord
}

func storeRecord(w http.ResponseWriter, r *http.Request, UID string) {
	successStore := true
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		checkWarn(err, "ParseMultipartForm err")
	}

	form := r.MultipartForm
	fmt.Println(form)
	recordName := form.Value["recordname"][0]
	organismName := form.Value["organismname"][0]
	result, storeRecordErr := db.Exec("INSERT INTO record SET userid=?, recordname=?, organismname=?, createtime=CURRENT_TIMESTAMP", UID, recordName, organismName)
	fmt.Println(storeRecordErr)
	if storeRecordErr == nil {
		id, getRecordIDErr := result.LastInsertId()
		if getRecordIDErr == nil {
			recordID := strconv.FormatInt(id, 10)
			//fmt.Println(recordID)
			for key, value := range r.MultipartForm.Value {
				if key == "recordname" || key == "organismname" || key == "photos" {
					continue
				}
				//fmt.Println(key, value)
				if len(value) == 1 {
					updateCommand := "UPDATE record SET `" + key + "`=?" + " WHERE id=?"
					//fmt.Println(updateCommand)
					_, updateErr := db.Exec(updateCommand, value[0], recordID)
					if updateErr != nil {
						successStore = false
					}
				} else {

				}
			}
			//fmt.Println(form)

			for _, fileHeaders := range r.MultipartForm.File {
				if len(fileHeaders) > 0 {
					for _, fileHeader := range fileHeaders {
						//fmt.Println(form)
						file, _ := fileHeader.Open()
						defer file.Close()
						//fmt.Println("filename:", fileHeader.Filename)
						//fmt.Println("bytes:", fileHeader.Size)
						randString := newRandomString(35)
						photoExt := filepath.Ext(fileHeader.Filename)
						photoPath := randString + photoExt

						//newFile, _ := os.Create(frogConfig.StoragePath + "photo/" + photoPath)
						//defer newFile.Close()
						//io.Copy(newFile, file)

						buf, _ := ioutil.ReadAll(file)
						ioutil.WriteFile(frogConfig.StoragePath+"photo/"+photoPath, buf, os.ModePerm)

						data, decodeErr := exif.Read(frogConfig.StoragePath + "photo/" + photoPath)
						//_, decodeErr := exif.Read(frogConfig.StoragePath + "photo/" + photoPath)

						checkWarn(decodeErr, "decode photo exif err")

						if decodeErr != nil {
							_, storeRecordPhotoErr := db.Exec("INSERT INTO photo SET userid=?, recordid=?, path=?, name=?, createtime=CURRENT_TIMESTAMP", UID, recordID, photoPath, fileHeader.Filename)
							checkErr(storeRecordPhotoErr, "store record photo err")
							if storeRecordPhotoErr != nil {
								successStore = false
							}
						} else {
							latitudePosition, longitudePosition, latitudeValue, longitudeValue, dateAndTime := "", "", "", "", ""
							for key, value := range data.Tags {
								switch key {
								case "North or South Latitude":
									latitudePosition = value
								case "East or West Longitude":
									longitudePosition = value
								case "Latitude":
									latitudeValue = value
								case "Longitude":
									longitudeValue = value
								case "Date and Time":
									dateAndTime = value
									charsDateAndTime := []rune(dateAndTime)
									charsDateAndTime[4], charsDateAndTime[7] = '-', '-'
									dateAndTime = string(charsDateAndTime)
								}
								//fmt.Println(key, value)
							}
							if latitudePosition != "" && longitudePosition != "" && latitudeValue != "" && longitudeValue != "" {
								latitude, longitude := parseCoordinate(latitudeValue, latitudePosition, longitudeValue, longitudePosition)
								result, storeRecordPhotoErr := db.Exec("INSERT INTO photo SET userid=?, recordid=?, path=?, name=?, longitude=?, latitude=?, createtime=CURRENT_TIMESTAMP", UID, recordID, photoPath, fileHeader.Filename, longitude, latitude)
								checkErr(storeRecordPhotoErr, "store record photo err")
								if storeRecordPhotoErr != nil {
									successStore = false
								} else {
									id, getPhotoIDErr := result.LastInsertId()
									if getPhotoIDErr == nil {
										photoID := strconv.FormatInt(id, 10)
										if dateAndTime != "" {
											updateCommand := "UPDATE photo SET `dateAndTime`=? WHERE id=?"
											_, updateErr := db.Exec(updateCommand, dateAndTime, photoID)
											if updateErr != nil {
												successStore = false
											}
										}
									}
								}
							}
							if latitudePosition == "" || longitudePosition == "" || latitudeValue == "" || longitudeValue == "" {
								result, storeRecordPhotoErr := db.Exec("INSERT INTO photo SET userid=?, recordid=?, path=?, name=?, createtime=CURRENT_TIMESTAMP", UID, recordID, photoPath, fileHeader.Filename)
								checkErr(storeRecordPhotoErr, "store record photo err")
								if storeRecordPhotoErr != nil {
									id, getPhotoIDErr := result.LastInsertId()
									if getPhotoIDErr != nil {
										successStore = false
									} else {
										photoID := strconv.FormatInt(id, 10)
										if dateAndTime != "" {
											updateCommand := "UPDATE photo SET `dateAndTime`=? WHERE id=?"
											_, updateErr := db.Exec(updateCommand, dateAndTime, photoID)
											if updateErr != nil {
												successStore = false
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	p := UploadPage{UploadStatus: successStore}
	b, _ := json.Marshal(p)
	w.Write(b)
}

func getAllRecords() Records {

	recordIDs, records := []int{}, Records{}
	records.Records = make(map[int]Record)
	idrows, queryErr := db.Query("SELECT id FROM record")
	checkErr(queryErr, "query organis id from comment with mysql error")
	defer idrows.Close()
	for idrows.Next() {
		var tmp int
		scanErr := idrows.Scan(&tmp)
		checkErr(scanErr, "scan organis id from comment with mysql error")
		recordIDs = append(recordIDs, tmp)
	}

	for index, id := range recordIDs {
		recordname, organismname, food, stage, season, status, habitat, note, createTime := "", "", "", "", "", "", "", "", ""
		db.QueryRow("SELECT recordname FROM record WHERE id = ?", id).Scan(&recordname)
		db.QueryRow("SELECT organismname FROM record WHERE id = ?", id).Scan(&organismname)
		db.QueryRow("SELECT food FROM record WHERE id = ?", id).Scan(&food)
		db.QueryRow("SELECT stage FROM record WHERE id = ?", id).Scan(&stage)
		db.QueryRow("SELECT season FROM record WHERE id = ?", id).Scan(&season)
		db.QueryRow("SELECT status FROM record WHERE id = ?", id).Scan(&status)
		db.QueryRow("SELECT habitat FROM record WHERE id = ?", id).Scan(&habitat)
		db.QueryRow("SELECT note FROM record WHERE id = ?", id).Scan(&note)
		db.QueryRow("SELECT createtime FROM record WHERE id = ?", id).Scan(&createTime)

		r := Record{
			ID:           id,
			RecordName:   recordname,
			OrganismName: organismname,
			Food:         food,
			Stage:        stage,
			Season:       season,
			Status:       status,
			Habitat:      habitat,
			Note:         note,
			CrateTime:    createTime}
		r.PhotoSrc = make(map[int]string)
		pathRows, queryErr := db.Query("SELECT path FROM photo WHERE recordid = ?", id)
		checkErr(queryErr, "query photo path from comment with mysql error")
		defer pathRows.Close()
		i := 0
		for pathRows.Next() {
			var tmp string
			scanErr := pathRows.Scan(&tmp)
			checkErr(scanErr, "scan photo path from comment with mysql error")

			r.PhotoSrc[i] = tmp
			i++
		}

		r.PhotoLatitude = make(map[int]string)
		latitudeRows, queryErr := db.Query("SELECT latitude FROM photo WHERE recordid = ?", id)
		checkErr(queryErr, "query photo latitude from comment with mysql error")
		defer latitudeRows.Close()
		j := 0
		for latitudeRows.Next() {
			var tmp string
			scanErr := latitudeRows.Scan(&tmp)
			checkErr(scanErr, "scan photo latitude from comment with mysql error")
			r.PhotoLatitude[i] = tmp
			j++
		}

		r.PhotoLongitude = make(map[int]string)
		longitudeRows, queryErr := db.Query("SELECT longitude FROM photo WHERE recordid = ?", id)
		checkErr(queryErr, "query photo longitude from comment with mysql error")
		defer longitudeRows.Close()
		k := 0
		for longitudeRows.Next() {
			var tmp string
			scanErr := longitudeRows.Scan(&tmp)
			checkErr(scanErr, "scan photo longitude from comment with mysql error")
			r.PhotoLongitude[i] = tmp
			k++
		}
		records.Records[index] = r
	}
	return records
}

func getRecordByRecordID(recordID string) Record {
	id, _ := strconv.Atoi(recordID)
	recordname, organismname, food, stage, season, status, habitat, note := "", "", "", "", "", "", "", ""
	db.QueryRow("SELECT recordname FROM record WHERE id = ?", recordID).Scan(&recordname)
	db.QueryRow("SELECT organismname FROM record WHERE id = ?", recordID).Scan(&organismname)
	db.QueryRow("SELECT food FROM record WHERE id = ?", recordID).Scan(&food)
	db.QueryRow("SELECT stage FROM record WHERE id = ?", recordID).Scan(&stage)
	db.QueryRow("SELECT season FROM record WHERE id = ?", recordID).Scan(&season)
	db.QueryRow("SELECT status FROM record WHERE id = ?", recordID).Scan(&status)
	db.QueryRow("SELECT habitat FROM record WHERE id = ?", recordID).Scan(&habitat)
	db.QueryRow("SELECT note FROM record WHERE id = ?", recordID).Scan(&note)

	r := Record{
		ID:           id,
		RecordName:   recordname,
		OrganismName: organismname,
		Food:         food,
		Stage:        stage,
		Season:       season,
		Status:       status,
		Habitat:      habitat,
		Note:         note}
	r.PhotoSrc = make(map[int]string)
	idrows, queryErr := db.Query("SELECT path FROM photo WHERE recordid = ?", recordID)
	checkErr(queryErr, "query photo path from comment with mysql error")
	defer idrows.Close()
	i := 0
	for idrows.Next() {
		var tmp string
		scanErr := idrows.Scan(&tmp)
		checkErr(scanErr, "scan photo path from comment with mysql error")
		r.PhotoSrc[i] = tmp
		i++
	}
	return r
}

func alterRecordByRecordID(r *http.Request) bool {
	successAlter := true
	r.ParseMultipartForm(32 << 20)
	recordID := r.Form.Get("recordid")
	for key, values := range r.Form {
		for _, value := range values {
			updateCommand := "UPDATE record SET " + key + "=?" + " WHERE id=?"
			// when value quantity == 1, can do this
			_, updateErr := db.Exec(updateCommand, value, recordID)
			if updateErr != nil {
				successAlter = false
			}
		}
	}
	return successAlter
}

func alterRecordPhotoByRecordID(r *http.Request) bool {
	successAlter := true
	return successAlter
}

func removeRecordByRecordID(recordID string) bool {
	successDelete := false
	_, deleteRecordWithMysqlErr := db.Exec("DELETE FROM record WHERE id=?", recordID)
	checkErr(deleteRecordWithMysqlErr, "deleteRecordWithMysqlErr")
	if deleteRecordWithMysqlErr == nil {
		successDelete = true
	}
	return successDelete
}
