package models

type User struct {
	Nickname string `gorm:"primaryKey"`
	Email    string
	Password string
	Descr    string
}
