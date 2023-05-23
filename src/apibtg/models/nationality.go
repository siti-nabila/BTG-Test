package models

import "net/http"

type NationalityController interface {
	GetAllNationalities(resp http.ResponseWriter, req *http.Request)
}

type NationalityService interface {
	GetAllNationality() ([]Nationality, error)
}
type NationalityRepository interface {
	FindAllNationality() ([]Nationality, error)
}

type (
	Nationality struct {
		NatId   int    `json:"NationalityId"`
		NatName string `json:"Nationality"`
		NatCode string `json:"NationalityCode"`
	}
)
