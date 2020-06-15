package controllers

import (
	"encoding/json"
	"github.com/rapando/finance-api/api/entities"
	"github.com/rapando/finance-api/api/middleware"
	"github.com/rapando/finance-api/api/models"
	"github.com/rapando/finance-api/api/utils"
	"net/http"
)

func (s *Server) Accounts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		response, err := models.ListAccounts(s.Db)
		if err != nil {
			middleware.ErrorResponse(w)
			return
		}
		middleware.OkResponse(w, 200, response)
	} else {
		var data entities.NewAccount
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			utils.Log("Unable to read payload because ", err)
			middleware.ErrorResponse(w)
			return
		}
		if err := data.Validate(); err != nil {
			middleware.ErrorResponse(w)
			return
		}

		id, err := models.CreateAccount(s.Db, data)
		if err != nil {
			middleware.ErrorResponse(w)
			return
		}
		middleware.OkResponse(w, 200, map[string]int64{"id": id})

	}
}
