package models

import "net/http"

type (
	FamilyController interface {
		GetAllFamily(resp http.ResponseWriter, req *http.Request)
		GetFamilyById(resp http.ResponseWriter, req *http.Request)
		PostFamily(resp http.ResponseWriter, req *http.Request)
		PutFamilyById(resp http.ResponseWriter, req *http.Request)
		DeleteFamilyById(resp http.ResponseWriter, req *http.Request)
	}

	FamilyService interface {
		GetAllFamily() ([]Family, error)
		FindFamilyById(id int) (Family, error)
		AddFamily(data Family) (Family, error)
		UpdateFamilyById(data Family) (Family, error)
		DeleteFamilyById(id int) error
	}
	FamilyRepository interface {
		FindAllFamily() ([]Family, error)
		FindFamilyById(id int) (Family, error)
		AddFamily(data Family) (int, error)
		UpdateById(data Family) (int, error)
		DeleteById(id int) error
	}

	Family struct {
		FamId   int    // will generate using util generate ID
		CustId  int    `json:"CustId"`
		FamRel  string `json:"Relationship"`
		FamName string `json:"Name"`
		FamDob  string `json:"DoB"` // format : YYYY-MM-DD
	}
)
