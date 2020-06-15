package models

import (
	"database/sql"
	"github.com/rapando/finance-api/api/entities"
	"github.com/rapando/finance-api/api/utils"
)

func ListAccounts(db *sql.DB) (data []entities.Account, err error) {
	utils.Log("List accounts")
	query := "select a.id, a.name, t.id as tid, t.name as tname, a.balance, a.active, a.created " +
		"from account a inner join account_type t on a.account_type_id = t.id " +
		"order by a.name"
	rows, err := db.Query(query)
	if err != nil {
		utils.Log("Unable to read accounts because ", err)
		return
	}

	for rows.Next() {
		var d entities.Account
		if err = rows.Scan(&d.ID, &d.Name, &d.AccountType.ID, &d.AccountType.Name, &d.Balance, &d.Active, &d.Created);
		err != nil {
			utils.Log("Unable to scan record because ", err)
			continue
		}
		data = append(data, d)
	}
	utils.Log("Found ", len(data), " accounts")
	return
}

func CreateAccount(db *sql.DB, d entities.NewAccount) (id int64, err error) {
	utils.Log("Creating a new account")
	query := "INSERT INTO account SET name=?, account_type_id=?, balance=?"
	utils.Log("Running query ", query, " with params : ",d.Name, d.AccountTypeID, d.Balance )
	row, err := db.Exec(query, d.Name, d.AccountTypeID, d.Balance)
	if err != nil {
		utils.Log("Unable to save account because ", err)
		return
	}
	id, _ = row.LastInsertId()
	utils.Log("Inserted account ", id)
	return
}