package putreqs

import (
	"database/sql"
	"encoding/json"
	"github.com/bektosh/studentsDatabase/handlers/getreqs"
	"github.com/bektosh/studentsDatabase/models"
	"io/ioutil"
	"net/http"
)

//UpdateStudent - updated specified info about a student
func UpdateStudent(w http.ResponseWriter, r *http.Request, db *sql.DB) ([]byte, error) {
	var (
		res     []byte
		student models.Student
	)
	body, err := getreqs.GetByID(w, r, db)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(body, &student)
	if err != nil {
		res = []byte("Could not decode json")
		return res, err
	}
	id := r.URL.Query().Get("id")
	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		res = []byte("Could not read the body of the request")
		return res, err
	}
	err = json.Unmarshal(body, &student)
	if err != nil {
		res = []byte("Could not decode json")
		return res, err
	}
	_, err = db.Exec(`UPDATE students_info SET name=$1, surname=$2, age=$3, level=$4, field=$5, gpa=$6, email=$7, address=$8 WHERE id=$9`,
		student.Name, student.Surname, student.Age, student.Level, student.Field, student.Gpa, student.Email, student.Address, id)
	if err != nil {
		res = []byte("Could not execute sql query")
		return res, err
	}
	res = []byte("Successfully updated student info")
	return res, nil
}
