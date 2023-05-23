package services

import (
	models "BTG-Test/src/apibtg/models"
)

type CustomerServiceImpl struct {
	customerRepo models.CustomerRepository
}

func CreateCustomerServiceImpl(custRepo models.CustomerRepository) models.CustomerService {
	return &CustomerServiceImpl{
		customerRepo: custRepo,
	}
}

func (cst *CustomerServiceImpl) GetAllCustomer() ([]models.Customer, error) {
	return cst.customerRepo.FindAllCustomer()
}
