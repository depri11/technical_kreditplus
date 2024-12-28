package models

import "time"

type Transaction struct {
	Id                uint64    `gorm:"column:id"`
	CustomerId        string    `gorm:"column:customer_id"`
	ContractNumber    string    `gorm:"column:contract_number"`
	Otr               float64   `gorm:"column:otr"`
	AdminFee          float64   `gorm:"column:admin_fee"`
	InstallmentAmount float64   `gorm:"column:installment_amount"`
	InterestAmount    float64   `gorm:"column:interest_amount"`
	AssetName         string    `gorm:"column:asset_name"`
	CreatedAt         time.Time `gorm:"column:created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at"`
}
