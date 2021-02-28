package getreqs

import (
	"database/sql"
	"encoding/json"
	"github.com/bektosh/studentsDatabase/models"
	"net/http"
)

// GetByID - queries for information of the student by his/her id
func GetByID(w http.ResponseWriter, r *http.Request, db *sql.DB) ([]byte, error) {
	id := r.URL.Query().Get("id")
	student := models.Student{}

	err := db.QueryRow("SELECT * FROM students_info WHERE id=$1", id).Scan(
		&student.Id, &student.Name, &student.Surname, &student.Age, &student.Level,
		&student.Field, &student.Gpa, &student.Email, &student.Address,
	)
	if err == sql.ErrNoRows {
		return []byte("There is no rows with this id!"), err
	} else if err != nil {
		return []byte("Sorry, internal error occurred"), err
	}

	res, err := json.Marshal(student)
	if err != nil {
		return []byte("Sorry, internal error occurred"), err
	}
	w.Header().Set("content-type", "application/json")
	return res, nil
}
