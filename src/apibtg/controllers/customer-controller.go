package controllers

import (
	models "BTG-Test/src/apibtg/models"
	util "BTG-Test/src/apibtg/utils"
	"net/http"
	"strconv"

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
	router.HandleFunc("/customer/{id}", cstController.GetCustomerById).Methods("GET")
	router.HandleFunc("/customer-fam/{id}", cstController.GetCustomerByIdWithFamily).Methods("GET")
	router.HandleFunc("/customer", cstController.PostCustomer).Methods("POST")
	router.HandleFunc("/customer/{id}", cstController.PutCustomerById).Methods("PUT")
	router.HandleFunc("/customer/{id}", cstController.DeleteCustomerById).Methods("DELETE")
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

func (cst *CustController) GetCustomerById(rw http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	uid, _ := strconv.Atoi(id)

	customer, errCst := cst.customerService.FindCustomerById(uid)
	if errCst != nil {
		util.HandleResponse(rw, http.StatusInternalServerError, "Couldn't get customer",
			map[string]interface{}{"results": errCst.Error()})
		return
	}

	util.HandleResponse(rw, http.StatusOK, "Data Found", map[string]interface{}{"result": customer})

}

func (cst *CustController) PostCustomer(rw http.ResponseWriter, req *http.Request) {
	var (
		request models.Customer
	)
	data := util.HandleRequest(rw, req)
	data.Decode(&request)
	customer, errCst := cst.customerService.AddCustomer(request)
	if errCst != nil {
		util.HandleResponse(rw, http.StatusInternalServerError, "Couldn't get customer",
			map[string]interface{}{"results": errCst.Error()})
		return
	}
	util.HandleResponse(rw, http.StatusOK, "Data successfully added", map[string]interface{}{"result": customer})

}

func (cst *CustController) PutCustomerById(rw http.ResponseWriter, req *http.Request) {
	var (
		request models.Customer
	)
	params := mux.Vars(req)
	id := params["id"]
	uid, _ := strconv.Atoi(id)

	data := util.HandleRequest(rw, req)

	data.Decode(&request)
	request.CustId = uid

	customer, errCst := cst.customerService.UpdateCustomerById(request)
	if errCst != nil {
		util.HandleResponse(rw, http.StatusInternalServerError, "Couldn't update customer",
			map[string]interface{}{"results": errCst.Error()})
		return
	}
	util.HandleResponse(rw, http.StatusOK, "Data successfully updated", map[string]interface{}{"result": customer})

}

func (cst *CustController) DeleteCustomerById(rw http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	uid, _ := strconv.Atoi(id)

	errCst := cst.customerService.DeleteCustomerById(uid)
	if errCst != nil {
		util.HandleResponse(rw, http.StatusInternalServerError, "Couldn't delete customer",
			map[string]interface{}{"results": errCst.Error()})
		return
	}
	util.HandleResponse(rw, http.StatusOK, "Data successfully deleted", map[string]interface{}{"result": ""})

}

func (cst *CustController) GetCustomerByIdWithFamily(rw http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	uid, _ := strconv.Atoi(id)

	customer, errCst := cst.customerService.FindCustomerByIdWithFamily(uid)
	if errCst != nil {
		util.HandleResponse(rw, http.StatusInternalServerError, "Couldn't get customer list",
			map[string]interface{}{"results": errCst.Error()})
		return
	}
	util.HandleResponse(rw, http.StatusOK, "Data Found", map[string]interface{}{"result": customer})
}
