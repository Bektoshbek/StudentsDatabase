package main

import (
	"database/sql"
	"fmt"
	"github.com/bektosh/studentsDatabase/handlers"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strings"
)

// Constants for database connection
const (
	DbName     = "<database name>"
	DbUser     = "<database user>"
	DbPassword = "<user password>"
	DbHost     = "<database host>"
	DbPort     = "<database port>"
	SslMode    = "<ssl mode>"
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
			res, err = handlers.GetStudent(w, r, db, "id")
		} else if strings.Contains(r.URL.Path[1:], "getbyname") {
			res, err = handlers.GetStudent(w, r, db, "name")
		} else if strings.Contains(r.URL.Path[1:], "getbylevel") {
			res, err = handlers.GetStudent(w, r, db, "level")
		} else if strings.Contains(r.URL.Path[1:], "getbyage") {
			res, err = handlers.GetStudent(w, r, db, "age")
		} else if strings.Contains(r.URL.Path[1:], "getbyfield") {
			res, err = handlers.GetStudent(w, r, db, "field")
		} else {
			res = []byte("Unknown Request, cannot perform any action!")
		}
	case "POST":
		if strings.Contains(r.URL.Path[1:], "addstudent") {
			res, err = handlers.AddStudent(r, db)
		} else {
			res = []byte("Unknown Request, cannot perform any action!")
		}
		// Call Handler function for adding a new student to the database
	case "PUT":
		if strings.Contains(r.URL.Path[1:], "updatestudent") {
			res, err = handlers.UpdateStudent(w, r, db)
		} else {
			res = []byte("Unknown Request, cannot perform any action!")
		}
	case "DELETE":
		if strings.Contains(r.URL.Path[1:], "deletestudent") {
			res, err = handlers.DeleteStudent(r, db)
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
