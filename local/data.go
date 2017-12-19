package main

import "net/http"

// Coordinate .
type Coordinate struct {
	Latitude  float64
	Longitude float64
}

// LoginPage .
type LoginPage struct {
	LoginStatus  bool
	LoginProblem string
}

// PrivateIndex .
type PrivateIndex struct {
	Username string
	UserID   int
}

// UploadPage .
type UploadPage struct {
	UploadStatus bool
}

// LibraryPage .
type LibraryPage struct {
	ModifyStatus bool
}

// LibraryData .
type LibraryData struct {
	ID             int
	OrganismName   string
	Kingdom        string
	Phylum         string
	Class          string
	Order          string
	Family         string
	Genus          string
	Species        string
	Food           string
	Season         string
	Status         string
	Habitat        string
	Note           string
	CrateTime      string
	PhotoSrc       map[int]string // index, photo path
	PhotoLatitude  map[int]string // index, photo latitude
	PhotoLongitude map[int]string // index, photo longitude
}

// Library .
type Library struct {
	LibraryDatas map[int]LibraryData
}

// Record .
type Record struct {
	ID             int
	RecordName     string
	OrganismName   string //有空要把單筆記錄改成可以記錄多個生物 []string
	ISAnimal       bool
	Kingdom        string
	Phylum         string
	Class          string
	Order          string
	Family         string
	Genus          string
	Species        string
	Food           string
	Stage          string
	Season         string
	Status         string
	Address        string
	Habitat        string
	Note           string
	CrateTime      string
	PhotoSrc       map[int]string // index, photo path
	PhotoLatitude  map[int]string // index, photo latitude
	PhotoLongitude map[int]string // index, photo longitude
}

// Records .
type Records struct {
	Records map[int]Record
}

func loadPrivateIndexData(w http.ResponseWriter, r *http.Request) PrivateIndex {
	l := PrivateIndex{}
	cu, _ := r.Cookie("username")
	l.Username = cu.Value
	return l
}
