package repositories

import (
	models "BTG-Test/src/apibtg/models"

	logger "github.com/sirupsen/logrus"
)

type CustomerRepositoryImpl struct {
	commonRepository CommonRepository
}

func CreateCustomerRepositoryImpl(commonRepo CommonRepository) models.CustomerRepository {
	return &CustomerRepositoryImpl{
		commonRepository: commonRepo,
	}

}

func (cst *CustomerRepositoryImpl) FindAllCustomer() ([]models.Customer, error) {
	var (
		res    models.Customer
		result []models.Customer
	)
	query := "SELECT \"cst_id\",\"nationality_id\",\"cst_name\",\"cst_dob\",\"cst_phoneNum\",\"cst_email\" FROM \"BTG_Schema\".\"Customer\""
	rows, errRows := cst.commonRepository.FindAll(query)
	if errRows != nil {
		logger.Error(errRows)
		return nil, errRows
	}
	defer rows.Close()
	for rows.Next() {
		errRows := rows.Scan(&res.CustId, &res.NatId, &res.CustName, &res.CustDob, &res.CustPhone, &res.CustEmail)
		if errRows != nil {
			logger.Error(errRows)
			return nil, errRows
		}
		result = append(result, res)
	}
	return result, nil
}
