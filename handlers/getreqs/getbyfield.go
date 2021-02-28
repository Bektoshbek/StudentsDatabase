package getreqs

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bektosh/studentsDatabase/models"
)

// GetByField - queries info about students by field
func GetByField(w http.ResponseWriter, r *http.Request, db *sql.DB) ([]byte, error) {
	field := r.URL.Query().Get("field")
	var (
		students []models.Student
		student  models.Student
	)
	rows, err := db.Query("SELECT * FROM students_info WHERE field=$1", field)
	if rows.Err() == sql.ErrNoRows {
		return []byte("There are no records for this field!"), nil
	} else if err != nil {
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
