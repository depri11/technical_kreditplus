package models

import "time"

type Customer struct {
	Id           uint      `gorm:"primaryKey;autoIncrement"`
	Nik          string    `gorm:"column:nik"`
	FullName     string    `gorm:"column:full_name"`
	LegalName    string    `gorm:"column:legal_name"`
	PlaceOfBirth string    `gorm:"column:place_of_birth"`
	DateOfBirth  time.Time `gorm:"column:date_of_birth"`
	Salary       float64   `gorm:"column:salary"`
	PhotoKtp     string    `gorm:"column:photo_ktp"`
	PhotoSelfie  string    `gorm:"column:photo_selfie"`
}
