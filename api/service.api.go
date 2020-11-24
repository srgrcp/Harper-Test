package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/srgrcp/Harper-Test/orm"
)

// GetServiceOrder rest-api endpoint
func GetServiceOrder(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	uuid := params["uuid"]
	service := orm.GetServiceOrder(uuid)
	json, _ := json.Marshal(*service)
	response.Header().Set("Content-Type", "application/json")
	response.Write(json)
}
