package controllers

import (
	models "BTG-Test/src/apibtg/models"
	util "BTG-Test/src/apibtg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type FamilyController struct {
	familyService models.FamilyService
}

func CreateFamilyController(router *mux.Router, service models.FamilyService) models.FamilyController {
	flController := FamilyController{
		familyService: service,
	}
	router.HandleFunc("/families", flController.GetAllFamily).Methods("GET")
	router.HandleFunc("/family/{id}", flController.GetFamilyById).Methods("GET")
	router.HandleFunc("/family", flController.PostFamily).Methods("POST")
	router.HandleFunc("/family/{id}", flController.PutFamilyById).Methods("PUT")
	router.HandleFunc("/family/{id}", flController.DeleteFamilyById).Methods("DELETE")
	return &flController
}

func (fl *FamilyController) GetAllFamily(rw http.ResponseWriter, req *http.Request) {
	familys, errCst := fl.familyService.GetAllFamily()
	if errCst != nil {
		util.HandleResponse(rw, http.StatusInternalServerError, "Couldn't get family list",
			map[string]interface{}{"results": errCst.Error()})
		return
	}
	if len(familys) < 1 {
		util.HandleResponse(rw, http.StatusNotFound, "No data in database",
			map[string]interface{}{"results": nil})
		return

	}
	util.HandleResponse(rw, http.StatusOK, "Data Found", map[string]interface{}{"results": familys})

}

func (fl *FamilyController) GetFamilyById(rw http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	fid, _ := strconv.Atoi(id)

	family, errCst := fl.familyService.FindFamilyById(fid)
	if errCst != nil {
		util.HandleResponse(rw, http.StatusInternalServerError, "Couldn't get family",
			map[string]interface{}{"results": errCst.Error()})
		return
	}

	util.HandleResponse(rw, http.StatusOK, "Data Found", map[string]interface{}{"result": family})

}

func (fl *FamilyController) PostFamily(rw http.ResponseWriter, req *http.Request) {
	var (
		request models.Family
	)
	data := util.HandleRequest(rw, req)
	data.Decode(&request)
	family, errCst := fl.familyService.AddFamily(request)
	if errCst != nil {
		util.HandleResponse(rw, http.StatusInternalServerError, "Couldn't get family",
			map[string]interface{}{"results": errCst.Error()})
		return
	}
	util.HandleResponse(rw, http.StatusOK, "Data successfully added", map[string]interface{}{"result": family})

}

func (fl *FamilyController) PutFamilyById(rw http.ResponseWriter, req *http.Request) {
	var (
		request models.Family
	)
	params := mux.Vars(req)
	id := params["id"]
	fid, _ := strconv.Atoi(id)

	data := util.HandleRequest(rw, req)

	data.Decode(&request)
	request.FamId = fid

	family, errCst := fl.familyService.UpdateFamilyById(request)
	if errCst != nil {
		util.HandleResponse(rw, http.StatusInternalServerError, "Couldn't update family",
			map[string]interface{}{"results": errCst.Error()})
		return
	}
	util.HandleResponse(rw, http.StatusOK, "Data successfully updated", map[string]interface{}{"result": family})

}

func (fl *FamilyController) DeleteFamilyById(rw http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	fid, _ := strconv.Atoi(id)

	errCst := fl.familyService.DeleteFamilyById(fid)
	if errCst != nil {
		util.HandleResponse(rw, http.StatusInternalServerError, "Couldn't delete family",
			map[string]interface{}{"results": errCst.Error()})
		return
	}
	util.HandleResponse(rw, http.StatusOK, "Data successfully deleted", map[string]interface{}{"result": ""})

}
