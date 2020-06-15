package entities

import (
	"errors"
	"github.com/rapando/finance-api/api/utils"
)

type Account struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	AccountType struct{
		ID int64 `json:"id"`
		Name string `json:"name"`
	} `json:"account_type"`
	Balance float64 `json:"balance"`
	Active bool `json:"active"`
	Created string `json:"created"`
}

type NewAccount struct {
	Name string `json:"name"`
	AccountTypeID int64 `json:"account_type_id"`
	Balance float64 `json:"balance"`
}

func (na NewAccount) Validate() (error) {
	if na.Name == "" {
		utils.Log("Name is required")
		return errors.New("Name is required")
	}
	if na.AccountTypeID == 0 {
		utils.Log("Account type id is required")
		return errors.New("Account ID is required")
	}
	return nil
}