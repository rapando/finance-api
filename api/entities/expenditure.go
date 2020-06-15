package entities

type Expenditure struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	TransactionType string `json:"transaction_type"`
	Active bool `json:"active"`
	Created string `json:"created"`
}
