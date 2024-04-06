package models

type BankAccount struct {
	GUID           string
	CustomerGUID   string
	CustomerBankID string
	BankName       string
	BranchName     string
	AccountNumber  string
	AccountName    string
}
