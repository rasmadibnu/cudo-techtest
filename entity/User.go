package entity

import "time"

type User struct {
	ID              int64     `gorm:"column:id" json:"id"`
	Name            string    `gorm:"column:name;NOT NULL" json:"name"`
	Email           string    `gorm:"column:email;NOT NULL" json:"email"`
	EmailVerifiedAt time.Time `gorm:"column:email_verified_at" json:"email_verified_at"`
	Password        string    `gorm:"column:password;NOT NULL" json:"password"`
	RememberToken   string    `gorm:"column:remember_token" json:"remember_token"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (m *User) TableName() string {
	return "users"
}
