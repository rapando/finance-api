package controllers

import "github.com/gorilla/mux"

func (s *Server) initRoutes() {
	s.Router = mux.NewRouter()
	s.Log("Initializing routes")

	s.Router.HandleFunc("/", s.Home).Methods("GET")

	//	account types
	s.Router.HandleFunc("/account-types", s.AccountTypes).Methods("GET", "POST")
	s.Router.HandleFunc("/accounts", s.Accounts).Methods("GET", "POST")
	s.Router.HandleFunc("/expenditures", s.Expenditures).Methods("GET", "POST")
}
