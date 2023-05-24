package repositories

import (
	database "BTG-Test/src/apibtg/database"
	models "BTG-Test/src/apibtg/models"

	logger "github.com/sirupsen/logrus"
)

type FamilyRepositoryImpl struct {
	commonRepository CommonRepository
}

func CreateFamilyRepositoryImpl(commonRepo CommonRepository) models.FamilyRepository {
	return &FamilyRepositoryImpl{
		commonRepository: commonRepo,
	}

}

func (cst *FamilyRepositoryImpl) FindAllFamily() ([]models.Family, error) {
	var (
		res    models.Family
		result []models.Family
	)
	query := "SELECT \"fl_id\",\"cst_id\",\"fl_relation\",\"fl_name\",\"fl_dob\" FROM \"BTG_Schema\".\"family_list\""
	rows, errRows := cst.commonRepository.FindAll(query)
	if errRows != nil {
		logger.Error(errRows)
		return nil, errRows
	}
	defer rows.Close()
	for rows.Next() {
		errRows := rows.Scan(&res.FamId, &res.CustId, &res.FamRel, &res.FamName, &res.FamDob)
		if errRows != nil {
			logger.Error(errRows)
			return nil, errRows
		}
		result = append(result, res)
	}
	return result, nil
}

func (cst *FamilyRepositoryImpl) FindFamilyById(id int) (models.Family, error) {
	var (
		res models.Family
	)
	query := "SELECT \"fl_id\",\"cst_id\",\"fl_relation\",\"fl_name\",\"fl_dob\" FROM \"BTG_Schema\".\"family_list\" WHERE \"fl_id\" = $1"
	row, errQuery := cst.commonRepository.FindById(query, id)
	if errQuery != nil {
		logger.Error(errQuery)
		return models.Family{}, errQuery
	}
	row.Scan(&res.FamId, &res.CustId, &res.FamRel, &res.FamName, &res.FamDob)

	return res, nil
}

func (cst *FamilyRepositoryImpl) AddFamily(data models.Family) (int, error) {
	var (
		famId int
	)
	req := map[int]interface{}{
		1: data.FamId,
		2: data.CustId,
		3: data.FamRel,
		4: data.FamName,
		5: data.FamDob,
	}
	query := "INSERT INTO \"BTG_Schema\".\"family_list\"(\"fl_id\",\"cst_id\",\"fl_relation\",\"fl_name\",\"fl_dob\") "
	row, errQuery := cst.commonRepository.StoreOne(query, "fl_id", req)
	if errQuery != nil {
		logger.Error(errQuery)
		return 0, errQuery
	}
	row.Scan(&famId)
	return famId, nil
}

func (cst *FamilyRepositoryImpl) UpdateById(data models.Family) (int, error) {
	var (
		id int
	)
	db, errDb := database.GetConnectionBTG()
	if errDb != nil {
		logger.Error(errDb)
		return 0, errDb
	}

	query := "UPDATE \"BTG_Schema\".\"family_list\" SET \"cst_id\"=$1, \"fl_relation\"=$2, \"fl_name\"=$3, \"fl_dob\"=$4 WHERE \"fl_id\"=$5 RETURNING \"fl_id\""

	errQuery := db.QueryRow(query, data.CustId, data.FamRel, data.FamName, data.FamDob, data.FamId).Scan(&id)
	if errQuery != nil {
		logger.Error(errQuery)
		return 0, errQuery
	}

	return id, nil

}

func (cst *FamilyRepositoryImpl) DeleteById(id int) error {
	errQuery := cst.commonRepository.DeleteById("family_list", "fl_id", id)
	if errQuery != nil {
		logger.Error(errQuery)
		return errQuery
	}

	return nil
}
