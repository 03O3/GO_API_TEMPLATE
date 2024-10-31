package models

type User struct {
	ID       int    `json:"id" gorm:"column:id"`
	Email    string `json:"email" gorm:"column:email"`
	Username string `json:"username" gorm:"column:username"`
	Telegram string `json:"telegram" gorm:"column:telegram"`
	Admin    bool   `json:"admin" gorm:"column:admin"`
}
