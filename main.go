package main

import (
	"database/sql"
	"fmt"
	"github.com/bektosh/studentsDatabase/handlers/deletereqs"
	"github.com/bektosh/studentsDatabase/handlers/getreqs"
	"github.com/bektosh/studentsDatabase/handlers/postreqs"
	"github.com/bektosh/studentsDatabase/handlers/putreqs"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strings"
)

// Constants for database connection
const (
	DbName     = "students"
	DbUser     = "bektosh"
	DbPassword = "030409"
	DbHost     = "localhost"
	DbPort     = "5432"
	SslMode    = "disable"
)

var db *sql.DB

// InitDbConn - Initializes database connection
func InitDbConn() error {
	var err error = nil
	ConnStr := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=%s",
		DbName, DbUser, DbPassword, DbHost, DbPort, SslMode)

	db, err = sql.Open("postgres", ConnStr)
	return err
}

// Handler ...
func Handler(w http.ResponseWriter, r *http.Request) {
	var (
		res []byte
		err error
	)

	switch r.Method {
	case "GET":
		if strings.Contains(r.URL.Path[1:], "getbyid") {
			res, err = getreqs.GetByID(w, r, db)
		} else if strings.Contains(r.URL.Path[1:], "getbyname") {
			//name := r.URL.Query().Get("name")
			res, err = getreqs.GetByName(w, r, db /*name*/)
		} else if strings.Contains(r.URL.Path[1:], "getbylevel") {
			res, err = getreqs.GetByLevel(w, r, db)
		} else if strings.Contains(r.URL.Path[1:], "getbyage") {
			res, err = getreqs.GetByAge(w, r, db)
		} else if strings.Contains(r.URL.Path[1:], "getbyfield") {
			res, err = getreqs.GetByField(w, r, db)
		} else {
			res = []byte("Unknown Request, cannot perform any action!")
		}
	case "POST":
		if strings.Contains(r.URL.Path[1:], "addstudent") {
			res, err = postreqs.AddStudent(r, db)
		} else {
			res = []byte("Unknown Request, cannot perform any action!")
		}
		// Call Handler function for adding a new student to the database
	case "PUT":
		if strings.Contains(r.URL.Path[1:], "updatestudent") {
			res, err = putreqs.UpdateStudent(w, r, db)
		} else {
			res = []byte("Unknown Request, cannot perform any action!")
		}
	case "DELETE":
		if strings.Contains(r.URL.Path[1:], "deletestudent") {
			res, err = deletereqs.DeleteStudent(r, db)
		} else {
			res = []byte("Unknown Request, cannot perform any action!")
		}
	default:
		res = []byte("Unknown Request, cannot perform any action!")
	}
	if err != nil {
		fmt.Println(string(res), ":", err)
	}

	_, err = w.Write(res)
	if err != nil {
		log.Fatal("Could not write: ", err)
	}
}

func main() {
	err := InitDbConn()
	if err != nil {
		log.Fatal("Could not connect to the database: ", err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
	http.HandleFunc("/", Handler)
	fmt.Println("Server listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
