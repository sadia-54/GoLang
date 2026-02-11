package models

import (
	"time"

)

type Student struct {
	ID  uint `gorm:"primaryKey"`
	Name  string `gorm:"size:100; not null"`
	Age  int `gorm:"not null"`
	Email  string `gorm:"size:150; unique; not null"`
	Department  string `gorm:"size:100; not null"`
	Session  string `gorm:"size:20; not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

}
