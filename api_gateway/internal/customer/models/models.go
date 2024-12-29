package models

import (
	"time"

	customer_proto "github.com/depri11/technical_kreditplus/protos"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetCustomerResponse struct {
	Nik          string  `json:"nik"`
	FullName     string  `json:"full_name"`
	LegalName    string  `json:"legal_name"`
	PlaceOfBirth string  `json:"place_of_birth"`
	DateOfBirth  string  `json:"date_of_birth"`
	Salary       float64 `json:"salary"`
	PhotoKtp     string  `json:"photo_ktp"`
	PhotoSelfie  string  `json:"photo_selfie"`
}

type CreateCustomerRequest struct {
	Nik          string  `json:"nik"`
	FullName     string  `json:"full_name"`
	LegalName    string  `json:"legal_name"`
	PlaceOfBirth string  `json:"place_of_birth"`
	DateOfBirth  string  `json:"date_of_birth"`
	Salary       float64 `json:"salary"`
	PhotoKtp     string  `json:"photo_ktp"`
	PhotoSelfie  string  `json:"photo_selfie"`
}

func (c *CreateCustomerRequest) ToProto() (*customer_proto.CreateCustomerRequest, error) {
	date, err := time.Parse("2006-01-02", c.DateOfBirth)
	if err != nil {
		return nil, err
	}
	return &customer_proto.CreateCustomerRequest{
		Nik:          c.Nik,
		FullName:     c.FullName,
		LegalName:    c.LegalName,
		PlaceOfBirth: c.PlaceOfBirth,
		DateOfBirth:  timestamppb.New(date),
		Salary:       c.Salary,
		PhotoKtp:     c.PhotoKtp,
		PhotoSelfie:  c.PhotoSelfie,
	}, nil
}

type UpdateCustomerRequest struct {
	Nik          string  `json:"nik"`
	FullName     string  `json:"full_name"`
	LegalName    string  `json:"legal_name"`
	PlaceOfBirth string  `json:"place_of_birth"`
	DateOfBirth  string  `json:"date_of_birth"`
	Salary       float64 `json:"salary"`
	PhotoKtp     string  `json:"photo_ktp"`
	PhotoSelfie  string  `json:"photo_selfie"`
}

func (c *UpdateCustomerRequest) ToProto() (*customer_proto.UpdateCustomerRequest, error) {
	date, err := time.Parse("2006-01-02", c.DateOfBirth)
	if err != nil {
		return nil, err
	}
	return &customer_proto.UpdateCustomerRequest{
		Nik:          c.Nik,
		FullName:     c.FullName,
		LegalName:    c.LegalName,
		PlaceOfBirth: c.PlaceOfBirth,
		DateOfBirth:  timestamppb.New(date),
		Salary:       c.Salary,
		PhotoKtp:     c.PhotoKtp,
		PhotoSelfie:  c.PhotoSelfie,
	}, nil
}
