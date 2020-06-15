package controllers

import (
	"encoding/json"
	"github.com/rapando/finance-api/api/entities"
	"github.com/rapando/finance-api/api/middleware"
	"github.com/rapando/finance-api/api/models"
	"github.com/rapando/finance-api/api/utils"
	"net/http"
)

func (s *Server) AccountTypes(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		utils.Log("Listing all account types")
		response, err := models.ListAccountTypes(s.Db)
		if err != nil {
			middleware.ErrorResponse(w)
			return
		}
		middleware.OkResponse(w, 200, response)
	} else {
		utils.Log("Creating new account types")
		var data entities.AccountType

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			utils.Log("Unable to read payload because ", err)
			middleware.ErrorResponse(w)
			return
		}
		if err := data.Validate(); err != nil {
			middleware.ErrorResponse(w)
			return
		}
		id, err := models.AddAccountType(s.Db, data)
		if err != nil {
			middleware.ErrorResponse(w)
			return
		}
		middleware.OkResponse(w, 200, map[string]int64 {
			"id": id,
		})
	}
}
