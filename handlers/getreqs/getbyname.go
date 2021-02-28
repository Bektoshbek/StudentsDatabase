package getreqs

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bektosh/studentsDatabase/models"
)

// GetByName - queries info about students by name
func GetByName(w http.ResponseWriter, r *http.Request, db *sql.DB) ([]byte, error) {
	name := r.URL.Query().Get("name")
	var (
		students []models.Student
		student  models.Student
	)
	rows, err := db.Query("SELECT * FROM students_info WHERE name=$1", name)
	if err != nil {
		fmt.Println("Error while querying")
		return []byte("Sorry, internal error occurred"), err
	}
	defer rows.Close()

	for rows.Next() {
		if rows.Err() == sql.ErrNoRows {
			return []byte("There are no records for this name!"), nil
		}
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
