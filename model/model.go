package model

type User struct {
	ID		int			`gorm:"column:reminder_id""AUTO_INCREMENT"`
	DATE	string		`gorm:"column:remind_date"`
	TODO		string		`gorm:"column:todolist"`
	//EMAIL		string		`gorm:"column:user_email"`
	//PASSWORD	string		`gorm:"column:user_password"`
}