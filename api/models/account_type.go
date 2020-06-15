package models

import (
	"database/sql"
	"github.com/rapando/finance-api/api/entities"
	"github.com/rapando/finance-api/api/utils"
)

func ListAccountTypes(db *sql.DB) (data []entities.AccountType, err error) {
	utils.Log("List account types")
	query := "SELECT id, name, active, created FROM account_type ORDER BY IF (name='Other',1,0), name ASC"
	utils.Log("Running query ", query)

	rows, err := db.Query(query)
	if err != nil {
		utils.Log("Unable to read account_types because ", err)
		return
	}
	for rows.Next() {
		var at entities.AccountType
		if err = rows.Scan(&at.ID, &at.Name, &at.Active, &at.Created); err != nil {
			utils.Log("Unable to scan row because ", err)
			continue
		}
		data = append(data, at)
	}
	utils.Log("Found ", len(data), " account types")
	return
}

func AddAccountType(db *sql.DB, data entities.AccountType) (ID int64, err error) {
	utils.Log("Inserting a new account type")
	query := "INSERT INTO account_type SET name=?"
	row, err := db.Exec(query, data.Name)
	if err != nil {
		utils.Log("Unable to save account type ", data.Name)
		return
	}
	ID, _ = row.LastInsertId()
	utils.Log("Inserted ", ID)
	return
}
