package models

import "time"

type User struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Email     string    `gorm:"column:email;uniqueIndex" json:"email"`
	Password  string    `gorm:"column:password" json:"-"`
	Name      *string   `gorm:"column:name" json:"name,omitempty"`
	Role      string    `gorm:"column:role;default:user" json:"role"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

func (User) TableName() string {
	return "User"
}
