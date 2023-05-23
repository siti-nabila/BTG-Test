package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func HandleResponse(rw http.ResponseWriter, status int, msg string, data map[string]interface{}) {
	response := Response{
		Status:  status,
		Message: msg,
		Data:    data,
	}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(response)
}

func HandleRequest(rw http.ResponseWriter, req *http.Request) *json.Decoder {
	rw.Header().Set("Context-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// decode the json request to user
	decoded := json.NewDecoder(req.Body)

	return decoded
}
