package models

import (
	"database/sql"
	"github.com/rapando/finance-api/api/entities"
	"github.com/rapando/finance-api/api/utils"
)

func ListExpenditures(db *sql.DB) (data []entities.Expenditure, err error) {
	utils.Log("List expenditures")
	query := "SELECT id, name, transaction_type, active, created" +
		"FROM expenditure " +
		"ORDER BY name, transaction_type;"
	utils.Log("Running query ", query)
	rows, err := db.Query(query)
	if err != nil {
		utils.Log("Unable to fetch expenditures because ", err)
		return
	}
	for rows.Next() {
		var d entities.Expenditure
		if err = rows.Scan(&d.ID, &d.Name, &d.TransactionType, &d.Active, &d.Created); err != nil {
			utils.Log("Unable to scan row because ", err)
			continue
		}
		data = append(data, d)
	}
	utils.Log("Found ", len(data), " expenditures")
	return
}
