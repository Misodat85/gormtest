package model

type User struct {
	ID		string			`gorm:"primary_key""column;user_id"`
	NAME		string		`gorm:"column:user_name"`
	EMAIL		string		`gorm:"column:user_email"`
	PASSWORD	string		`gorm:"column:user_password"`
}