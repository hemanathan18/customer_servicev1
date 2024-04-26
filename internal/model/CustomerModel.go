package model

import "time"

type Customer struct {
	Customer_ID uint64 `gorm:"primaryKey"`
	Name        string
	Gender      string
	DOB         time.Time `gorm:"type:date"`
}
