package deletereqs

import (
	"database/sql"
	"net/http"
)

// DeleteStudent - deletes student from the database by id
func DeleteStudent(r *http.Request, db *sql.DB) ([]byte, error) {
	var res []byte
	id := r.URL.Query().Get("id")
	_, err := db.Exec(`DELETE FROM students_info WHERE id=$1`, id)
	if err != nil {
		res = []byte("Could not execute sql query")
		return res, err
	}
	res = []byte("Successfully deleted student!")
	return res, nil
}
