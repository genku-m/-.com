package models

import "time"

type Invoice struct {
	GUID              string
	CompanyGUID       string
	CustomerGUID      string
	PublishDate       time.Time
	Payment           uint64
	CommissionTax     uint64
	CommissionTaxRate float64
	ConsumptionTax    uint64
	TaxRate           float64
	BillingAmount     uint64
	PaymentDate       time.Time
	Status            InvoiceStatus
}
