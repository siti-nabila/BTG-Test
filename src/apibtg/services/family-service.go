package services

import (
	models "BTG-Test/src/apibtg/models"
	"BTG-Test/src/apibtg/utils"
)

type FamilyServiceImpl struct {
	familyRepo models.FamilyRepository
}

func CreateFamilyServiceImpl(familyRepo models.FamilyRepository) models.FamilyService {
	return &FamilyServiceImpl{
		familyRepo: familyRepo,
	}
}

func (cst *FamilyServiceImpl) GetAllFamily() ([]models.Family, error) {
	return cst.familyRepo.FindAllFamily()
}

func (cst *FamilyServiceImpl) FindFamilyById(id int) (models.Family, error) {
	return cst.familyRepo.FindFamilyById(id)
}

func (cst *FamilyServiceImpl) AddFamily(data models.Family) (models.Family, error) {
	// Check For Existing Family ID on database
	for {
		fid := utils.GenerateID()
		family, _ := cst.FindFamilyById(fid)
		if family == (models.Family{}) {
			data.FamId = fid
			break
		}
	}
	id, errCust := cst.familyRepo.AddFamily(data)
	if errCust != nil {
		return models.Family{}, errCust
	}
	return cst.FindFamilyById(id)
}

func (cst *FamilyServiceImpl) DeleteFamilyById(id int) error {
	errCust := cst.familyRepo.DeleteById(id)
	if errCust != nil {
		return errCust
	}
	return nil
}

func (cst *FamilyServiceImpl) UpdateFamilyById(data models.Family) (models.Family, error) {
	id, errCust := cst.familyRepo.UpdateById(data)
	if errCust != nil {
		return models.Family{}, errCust
	}

	res, errRes := cst.FindFamilyById(id)
	if errRes != nil {
		return models.Family{}, errRes
	}

	return res, nil

}
