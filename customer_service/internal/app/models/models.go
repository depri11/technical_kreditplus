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

type CustomerLimit struct {
	CustomerId  uint    `gorm:"primaryKey,column:customer_id"`
	Tenor1Month float64 `gorm:"column:tenor_1_month"`
	Tenor2Month float64 `gorm:"column:tenor_2_months"`
	Tenor3Month float64 `gorm:"column:tenor_3_months"`
	Tenor4Month float64 `gorm:"column:tenor_4_months"`
}

type CreateCustomerLimit struct {
	CustomerId  uint    `gorm:"column:customer_id"`
	Tenor1Month float64 `gorm:"column:tenor_1_month"`
	Tenor2Month float64 `gorm:"column:tenor_2_months"`
	Tenor3Month float64 `gorm:"column:tenor_3_months"`
	Tenor4Month float64 `gorm:"column:tenor_4_months"`
}

type UpdateCustomerLimit struct {
	CustomerId  uint    `gorm:"column:customer_id"`
	Tenor1Month float64 `gorm:"column:tenor_1_month"`
	Tenor2Month float64 `gorm:"column:tenor_2_months"`
	Tenor3Month float64 `gorm:"column:tenor_3_months"`
	Tenor4Month float64 `gorm:"column:tenor_4_months"`
}
