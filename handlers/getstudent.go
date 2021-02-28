package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bektosh/studentsDatabase/models"
)

// GetStudent - queries info about students by name
func GetStudent(w http.ResponseWriter, r *http.Request, db *sql.DB, param string) ([]byte, error) {
	value := r.URL.Query().Get(param)
	var (
		students []models.Student
		student  models.Student
	)
	sqlQuery := fmt.Sprintf("SELECT * FROM students_info WHERE %s=$1", param)
	rows, err := db.Query(sqlQuery, value)
	if err != nil {
		fmt.Println("Error while querying")
		return []byte("Sorry, internal error occurred"), err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&student.Id, &student.Name, &student.Surname, &student.Age, &student.Level, &student.Field,
			&student.Gpa, &student.Email, &student.Address)
		if err != nil {
			fmt.Println("Error while Scanning")
			return []byte("Sorry, internal error occurred"), err
		}
		students = append(students, student)
	}
	res, err := json.Marshal(students)
	if err != nil {
		fmt.Println("Error while converting into json")
		return []byte("Sorry, internal error occurred"), err
	}
	w.Header().Set("content-type", "application/json")
	return res, nil
}
