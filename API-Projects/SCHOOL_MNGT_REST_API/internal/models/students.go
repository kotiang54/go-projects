package models

type Student struct {
	ID        int    `json:"id,omitempty" db:"id"`
	FirstName string `json:"first_name,omitempty" db:"first_name"`
	LastName  string `json:"last_name,omitempty" db:"last_name"`
	Email     string `json:"email,omitempty" db:"email"`
	Class     string `json:"class,omitempty" db:"class"`
	TeacherID int    `json:"teacher_id,omitempty" db:"teacher_id"`
}
