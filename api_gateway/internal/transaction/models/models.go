package models

import "time"

type GetTransactionResponse struct {
	Id                string
	CustomerId        string
	ContractNumber    string
	Otr               float64
	AdminFee          float64
	InstallmentAmount float64
	InterestAmount    float64
	AssetName         string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
