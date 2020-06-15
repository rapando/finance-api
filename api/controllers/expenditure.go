package controllers

import (
	"github.com/rapando/finance-api/api/middleware"
	"github.com/rapando/finance-api/api/models"
	"net/http"
)

func (s *Server) Expenditures(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, err := models.ListExpenditures(s.Db)
		if err != nil {
			middleware.ErrorResponse(w)
			return
		}
		middleware.OkResponse(w, 200, data)
	}
}