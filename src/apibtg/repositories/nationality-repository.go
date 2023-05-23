package repositories

import (
	models "BTG-Test/src/apibtg/models"

	logger "github.com/sirupsen/logrus"
)

type NatRepositoryImpl struct {
	commonRepository CommonRepository
}

func CreateNationalityRepositoryImpl(commonRepo CommonRepository) models.NationalityRepository {
	return &NatRepositoryImpl{
		commonRepository: commonRepo,
	}

}

func (n *NatRepositoryImpl) FindAllNationality() ([]models.Nationality, error) {
	var (
		res    models.Nationality
		result []models.Nationality
	)

	query := "SELECT * FROM \"BTG_Schema\".\"Nationality\""

	rows, errQuery := n.commonRepository.FindAll(query)
	if errQuery != nil {
		logger.Error(errQuery)
		return nil, errQuery
	}
	defer rows.Close()
	for rows.Next() {
		errRows := rows.Scan(&res.NatId, &res.NatName, &res.NatCode)
		if errRows != nil {
			logger.Error(errRows)
			return nil, errRows
		}
		result = append(result, res)
	}

	return result, nil

}
