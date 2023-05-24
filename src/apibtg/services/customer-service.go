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

func (cst *CustomerServiceImpl) FindCustomerById(id int) (models.Customer, error) {
	return cst.customerRepo.FindCustomerById(id)
}

func (cst *CustomerServiceImpl) AddCustomer(data models.Customer) (models.Customer, error) {
	id, errCust := cst.customerRepo.AddCustomer(data)
	if errCust != nil {
		return models.Customer{}, errCust
	}
	return cst.FindCustomerById(id)
}

func (cst *CustomerServiceImpl) DeleteCustomerById(id int) error {
	errCust := cst.customerRepo.DeleteById(id)
	if errCust != nil {
		return errCust
	}
	return nil
}

func (cst *CustomerServiceImpl) UpdateCustomerById(data models.Customer) (models.Customer, error) {
	id, errCust := cst.customerRepo.UpdateById(data)
	if errCust != nil {
		return models.Customer{}, errCust
	}

	res, errRes := cst.FindCustomerById(id)
	if errRes != nil {
		return models.Customer{}, errRes
	}

	return res, nil

}

func (cst *CustomerServiceImpl) FindCustomerByIdWithFamily(id int) (models.CustomerFamily, error) {
	res, errCust := cst.customerRepo.FindCustomerByIdWithChild(id)
	if errCust != nil {
		return models.CustomerFamily{}, errCust
	}

	return res, nil

}
