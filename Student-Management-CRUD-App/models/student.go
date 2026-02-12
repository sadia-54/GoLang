package models

import (
	"time"

)

type Student struct {
	ID  uint `gorm:"primaryKey"`
	Name  string `json:"name" validate:"required,min=2" gorm:"size:100; not null"`
	Age  int `json:"age" validate:"required,gt=0" gorm:"not null"`
	Email  string `json:"email" validate:"required,email" gorm:"size:150; unique; not null"`
	Department  string `json:"department" validate:"required,min=3" gorm:"size:100; not null"`
	Session  string `json:"session" validate:"required" gorm:"size:20; not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

}
