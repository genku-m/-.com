package models

type User struct {
	GUID        string
	CompanyGUID string
	Name        string
	Email       string
	Password    string
}

type LoginInfo struct {
	GUID        string `json:"guid"`
	CompanyGUID string `json:"company_guid"`
}
