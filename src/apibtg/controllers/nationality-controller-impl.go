package controllers

import (
	models "BTG-Test/src/apibtg/models"
	util "BTG-Test/src/apibtg/utils"

	"net/http"

	"github.com/gorilla/mux"
)

type NatController struct {
	nationalityService models.NationalityService
}

func CreateNationalityController(router *mux.Router, service models.NationalityService) models.NationalityController {
	natController := NatController{
		nationalityService: service,
	}
	router.HandleFunc("/nationalities", natController.GetAllNationalities).Methods("GET")

	return &natController
}

func (n *NatController) GetAllNationalities(rw http.ResponseWriter, req *http.Request) {
	nationalities, errNat := n.nationalityService.GetAllNationality()
	if errNat != nil {
		util.HandleResponse(rw, http.StatusInternalServerError, "Couldn't get nationalities",
			map[string]interface{}{"results": errNat.Error()})
		return
	}
	if len(nationalities) < 1 {
		util.HandleResponse(rw, http.StatusNotFound, "No data in database",
			map[string]interface{}{"results": nil})
		return

	}

	util.HandleResponse(rw, http.StatusOK, "Data Found", map[string]interface{}{"results": nationalities})

}
