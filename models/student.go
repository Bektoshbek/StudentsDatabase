package models

type Student struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Age     int     `json:"age"`
	Level   int     `json:"level"`
	Field   string  `json:"field"`
	Gpa     string  `json:"gpa"`
	Email   string  `json:"email"`
	Address string  `json:"address"`
}
