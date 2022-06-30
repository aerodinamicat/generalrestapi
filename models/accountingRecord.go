package models

type AccountingRecord struct {
	Id         string  `json:"id"`
	Account    uint16  `json:"account"`
	Subaccount uint16  `json:"subaccount"`
	Concept    string  `json:"concept"`
	Type       string  `json:"type"`
	Amount     float32 `json:"amount"`
}
