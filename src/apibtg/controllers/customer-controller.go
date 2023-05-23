package controllers

import (
	models "BTG-Test/src/apibtg/models"
	util "BTG-Test/src/apibtg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type CustController struct {
	customerService models.CustomerService
}

func CreateCustomerController(router *mux.Router, service models.CustomerService) models.CustomerController {
	cstController := CustController{
		customerService: service,
	}
	router.HandleFunc("/customers", cstController.GetAllCustomer).Methods("GET")
	return &cstController
}

func (cst *CustController) GetAllCustomer(rw http.ResponseWriter, req *http.Request) {
	customers, errCst := cst.customerService.GetAllCustomer()
	if errCst != nil {
		util.HandleResponse(rw, http.StatusInternalServerError, "Couldn't get customer list",
			map[string]interface{}{"results": errCst.Error()})
		return
	}
	if len(customers) < 1 {
		util.HandleResponse(rw, http.StatusNotFound, "No data in database",
			map[string]interface{}{"results": nil})
		return

	}
	util.HandleResponse(rw, http.StatusOK, "Data Found", map[string]interface{}{"results": customers})

}
