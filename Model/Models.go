package model

type User struct {
	ID       int64  `gorm:"primarykey"`
	Username string `gorm:"unique_index;not null"`
	Password string `gorm:"not null"`
}
