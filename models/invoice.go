package models

import "time"

type Invoice struct {
	GUID           string
	CompanyGUID    string
	CustomerGUID   string
	PublishDate    time.Time
	Payment        uint32
	Commission     float64
	CommissionTax  float64
	ConsumptionTax uint32
	TaxRate        uint32
	BillingAmount  uint32
	PaymentDate    time.Time
	Status         InvoiceStatus
}
