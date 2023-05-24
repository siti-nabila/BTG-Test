package repositories

import (
	database "BTG-Test/src/apibtg/database"
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

func (cst *CustomerRepositoryImpl) FindCustomerById(id int) (models.Customer, error) {
	var (
		result models.Customer
	)
	query := "SELECT \"cst_id\",\"nationality_id\",\"cst_name\",\"cst_dob\",\"cst_phoneNum\",\"cst_email\" FROM \"BTG_Schema\".\"Customer\" WHERE \"cst_id\" = $1"
	row, errQuery := cst.commonRepository.FindById(query, id)
	if errQuery != nil {
		logger.Error(errQuery)
		return models.Customer{}, errQuery
	}
	row.Scan(&result.CustId, &result.NatId, &result.CustName, &result.CustDob, &result.CustPhone, &result.CustEmail)

	return result, nil
}

func (cst *CustomerRepositoryImpl) AddCustomer(data models.Customer) (int, error) {
	var (
		custId int
	)
	req := map[int]interface{}{
		1: data.NatId,
		2: data.CustName,
		3: data.CustDob,
		4: data.CustPhone,
		5: data.CustEmail,
	}
	query := "INSERT INTO \"BTG_Schema\".\"Customer\"(\"nationality_id\", \"cst_name\", \"cst_dob\",\"cst_phoneNum\", \"cst_email\") "

	row, errQuery := cst.commonRepository.StoreOne(query, "cst_id", req)
	if errQuery != nil {
		logger.Error(errQuery)
		return 0, errQuery
	}
	row.Scan(&custId)
	return custId, nil
}

func (cst *CustomerRepositoryImpl) UpdateById(data models.Customer) (int, error) {
	var (
		id int
	)
	db, errDb := database.GetConnectionBTG()
	if errDb != nil {
		logger.Error(errDb)
		return 0, errDb
	}

	query := "UPDATE \"BTG_Schema\".\"Customer\" SET \"nationality_id\"=$1, \"cst_name\"=$2, \"cst_dob\"=$3, \"cst_phoneNum\"=$4, \"cst_email\"=$5 WHERE \"cst_id\"=$6 RETURNING \"cst_id\""

	errQuery := db.QueryRow(query, data.NatId, data.CustName, data.CustDob, data.CustPhone, data.CustEmail, data.CustId).Scan(&id)
	if errQuery != nil {
		logger.Error(errQuery)
		return 0, errQuery
	}

	return id, nil

}

func (cst *CustomerRepositoryImpl) DeleteById(id int) error {
	errQuery := cst.commonRepository.DeleteById("Customer", "cst_id", id)
	if errQuery != nil {
		logger.Error(errQuery)
		return errQuery
	}

	return nil
}

func (cst *CustomerRepositoryImpl) FindCustomerByIdWithChild(id int) (models.CustomerFamily, error) {
	var (
		result models.CustomerFamily
		family models.Family
	)

	db, errDb := database.GetConnectionBTG()
	if errDb != nil {
		logger.Error(errDb)
		return models.CustomerFamily{}, errDb
	}

	query := `SELECT cst."cst_id", "nationality_id", "cst_name" AS "CustName", "cst_dob" AS "CustDoB", "cst_phoneNum" as "PhoneNo", "cst_email" as "CustEmail", "fl_id", fl."cst_id", "fl_relation", "fl_name" AS "FamName", "fl_dob" AS "FamDoB" FROM "BTG_Schema"."Customer" cst INNER JOIN "BTG_Schema"."family_list" fl ON fl."cst_id" = cst."cst_id" WHERE cst."cst_id" = $1`

	rows, errQuery := db.Query(query, id)
	if errQuery != nil {
		logger.Error(errQuery)
		return models.CustomerFamily{}, nil
	}
	defer rows.Close()
	for rows.Next() {
		errRows := rows.Scan(&result.Customer.CustId, &result.Customer.NatId, &result.Customer.CustName, &result.Customer.CustDob, &result.Customer.CustPhone, &result.Customer.CustEmail, &family.FamId, &family.CustId, &family.FamRel, &family.FamName, &family.FamDob)
		if errRows != nil {
			logger.Error(errRows)
			return models.CustomerFamily{}, errRows
		}

		result.Family = append(result.Family, family)
	}
	return result, nil
}
