package models

import "net/http"

type (
	CustomerController interface {
		GetAllCustomer(resp http.ResponseWriter, req *http.Request)
		GetCustomerById(resp http.ResponseWriter, req *http.Request)
		GetCustomerByIdWithFamily(resp http.ResponseWriter, req *http.Request)
		PostCustomer(resp http.ResponseWriter, req *http.Request)
		PostCustomerWithFamily(resp http.ResponseWriter, req *http.Request)
		PutCustomerById(resp http.ResponseWriter, req *http.Request)
		DeleteCustomerById(resp http.ResponseWriter, req *http.Request)
	}

	CustomerService interface {
		GetAllCustomer() ([]Customer, error)
		FindCustomerById(id int) (Customer, error)
		FindCustomerByIdWithFamily(id int) (CustomerFamily, error)
		AddCustomer(data Customer) (Customer, error)
		UpdateCustomerById(data Customer) (Customer, error)
		DeleteCustomerById(id int) error
	}
	CustomerRepository interface {
		FindAllCustomer() ([]Customer, error)
		FindCustomerById(id int) (Customer, error)
		FindCustomerByIdWithChild(id int) (CustomerFamily, error)
		AddCustomer(data Customer) (int, error)
		UpdateById(data Customer) (int, error)
		DeleteById(id int) error
	}

	RequestCustomer struct {
	}

	Customer struct {
		CustId    int    `json:"CustId,omitempty"`
		NatId     int    `json:"NationalityId"`
		CustName  string `json:"Name"`
		CustDob   string `json:"DoB"`
		CustPhone string `json:"PhoneNo"`
		CustEmail string `json:"Email"`
	}
	CustomerFamily struct {
		Customer Customer `json:"Customer"`
		Family   []Family `json:"Family"`
	}
)
