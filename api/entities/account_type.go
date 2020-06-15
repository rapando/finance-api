package entities

import (
	"errors"
	"github.com/rapando/finance-api/api/utils"
)

type AccountType struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Active bool `json:"active"`
	Created string `json:"created"`
}

func (a AccountType) Validate() (error) {
	if a.Name == "" {
		utils.Log("Name is required")
		return errors.New("Name is required")
	}
	return nil
}