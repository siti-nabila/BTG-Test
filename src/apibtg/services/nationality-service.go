package services

import (
	models "BTG-Test/src/apibtg/models"
)

type NationalityServiceImpl struct {
	nationalityRepo models.NationalityRepository
}

func CreateNationalityServiceImpl(nationalityRepo models.NationalityRepository) models.NationalityService {
	return &NationalityServiceImpl{
		nationalityRepo: nationalityRepo,
	}
}

func (n NationalityServiceImpl) GetAllNationality() ([]models.Nationality, error) {
	return n.nationalityRepo.FindAllNationality()
}
