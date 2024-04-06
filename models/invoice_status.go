package models

type InvoiceStatus string

const (
	InvoiceStatusUnprocessed InvoiceStatus = "unprocessed"
	InvoiceStatusProcessing  InvoiceStatus = "processing"
	InvoiceStatusPaied       InvoiceStatus = "paied"
	InvoiceStatusError       InvoiceStatus = "error"
)
