package models

import "net/http"

type (
	CustomerController interface {
		GetAllCustomer(resp http.ResponseWriter, req *http.Request)
	}

	CustomerService interface {
		GetAllCustomer() ([]Customer, error)
	}
	CustomerRepository interface {
		FindAllCustomer() ([]Customer, error)
		// FindOne(id int) (Customer, error)
		// StoreOne(data Customer) (Customer, error)
		// Update(data Customer) (Customer, error)
		// DeleteById(id int) error
	}

	RequestCustomer struct {
	}

	Customer struct {
		CustId    int    `json:"CustId"` // will use util generate ID
		NatId     int    `json:"NationalityId"`
		CustName  string `json:"Name"`
		CustDob   string `json:"DoB"` // format : YYYY-MM-DD
		CustPhone string `json:"PhoneNo"`
		CustEmail string `json:"Email"`
	}
)
