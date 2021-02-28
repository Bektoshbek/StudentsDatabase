package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/bektosh/studentsDatabase/models"
	"io/ioutil"
	"net/http"
)

// AddStudent - adds student to the database
func AddStudent(r *http.Request, db *sql.DB) ([]byte, error) {
	var (
		res     []byte
		student models.Student
		id      string
	)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res = []byte("Could not read the body of the request")
		return res, err
	}
	err = json.Unmarshal(body, &student)
	if err != nil {
		res = []byte("Could not decode json into student")
		return res, err
	}
	rows, err := db.Query(
		`INSERT INTO students_info (name, surname, age, level, field, gpa, email, address)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`,
		student.Name, student.Surname, student.Age, student.Level, student.Field, student.Gpa, student.Email, student.Address,
	)
	defer rows.Close()
	if err != nil {
		res = []byte("Could not execute sql query")
		return res, err
	}
	rows.Next()
	err = rows.Scan(&id)
	if err != nil {
		res = []byte("Could not scan an id")
		return res, err
	}
	res = []byte("Successfully added student at this id: " + id)
	return res, nil
}
