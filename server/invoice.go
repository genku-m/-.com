package server

import (
	"fmt"
	"time"

	errpkg "github.com/genku-m/upsider-cording-test/invoice/errors"
	"github.com/genku-m/upsider-cording-test/models"
	"github.com/gin-gonic/gin"
)

type CreateInvoiceRequest struct {
	CompanyGUID       string    `json:"company_guid" binding:"required"`
	CustomerGUID      string    `json:"customer_guid" binding:"required"`
	PublishDate       time.Time `json:"publish_date" binding:"required"`
	Payment           uint64    `json:"payment" binding:"required"`
	CommissionTaxRate float64   `json:"commission_tax_rate" binding:"required"`
	TaxRate           float64   `json:"tax_rate" binding:"required"`
	PaymentDate       time.Time `json:"payment_date" binding:"required"`
}

type InvoiceResponse struct {
	GUID              string    `json:"guid" binding:"required"`
	CompanyGUID       string    `json:"company_guid" binding:"required"`
	CustomerGUID      string    `json:"customer_guid" binding:"required"`
	PublishDate       time.Time `json:"publish_date" binding:"required"`
	Payment           uint64    `json:"payment" binding:"required"`
	CommissionTax     uint64    `json:"commission_tax" binding:"required"`
	CommissionTaxRate float64   `json:"commission_tax_rate" binding:"required"`
	ConsumptionTax    uint64    `json:"consumption_tax" binding:"required"`
	TaxRate           float64   `json:"tax_rate" binding:"required"`
	BillingAmount     uint64    `json:"billing_amount" binding:"required"`
	PaymentDate       time.Time `json:"payment_date" binding:"required"`
	Status            string    `json:"status" binding:"required"`
}

func (s *Server) CreateInvoice(ctx *gin.Context, loginInfo *models.LoginInfo) (*InvoiceResponse, error) {
	var cir CreateInvoiceRequest
	if err := ctx.ShouldBindJSON(&cir); err != nil {
		return nil, errpkg.NewInvalidArgumentError(err)
	}
	if loginInfo.CompanyGUID != cir.CompanyGUID {
		return nil, errpkg.NewInvalidArgumentError(fmt.Errorf("company_guid is invalid: %v", cir.CompanyGUID))
	}

	invoice, err := s.invoiceUsecase.Create(ctx, cir.CompanyGUID, cir.CustomerGUID, cir.PublishDate, cir.Payment, cir.CommissionTaxRate, cir.TaxRate, cir.PaymentDate)
	if err != nil {
		return nil, err
	}

	return &InvoiceResponse{
		GUID:              invoice.GUID,
		CompanyGUID:       invoice.CompanyGUID,
		CustomerGUID:      invoice.CustomerGUID,
		PublishDate:       invoice.PublishDate,
		Payment:           invoice.Payment,
		CommissionTax:     invoice.CommissionTax,
		CommissionTaxRate: invoice.CommissionTaxRate,
		ConsumptionTax:    invoice.ConsumptionTax,
		TaxRate:           invoice.TaxRate,
		BillingAmount:     invoice.BillingAmount,
		PaymentDate:       invoice.PaymentDate,
		Status:            string(invoice.Status),
	}, nil
}

type ListInvoiceRequest struct {
	CompanyGUID      string    `json:"company_guid" binding:"required"`
	FirstPaymentDate time.Time `json:"first_payment_date" binding:"required"`
	LastPaymentDate  time.Time `json:"last_payment_date" binding:"required"`
}

func (s *Server) ListInvoice(ctx *gin.Context, loginInfo *models.LoginInfo) ([]*InvoiceResponse, error) {
	firstPaymentDateString := ctx.Query("first_payment_date")
	lastPaymentDateString := ctx.Query("last_payment_date")
	firstPaymentDate, err := time.Parse(time.RFC3339, firstPaymentDateString)
	if err != nil {
		return nil, errpkg.NewInvalidArgumentError(err)
	}
	lastPaymentDate, err := time.Parse(time.RFC3339, lastPaymentDateString)
	if err != nil {
		return nil, errpkg.NewInvalidArgumentError(err)
	}

	invoices, err := s.invoiceUsecase.List(ctx, loginInfo.CompanyGUID, firstPaymentDate, lastPaymentDate)
	if err != nil {
		return nil, err
	}

	var res []*InvoiceResponse
	for _, invoice := range invoices {
		res = append(res, &InvoiceResponse{
			GUID:              invoice.GUID,
			CompanyGUID:       invoice.CompanyGUID,
			CustomerGUID:      invoice.CustomerGUID,
			PublishDate:       invoice.PublishDate,
			Payment:           invoice.Payment,
			CommissionTax:     invoice.CommissionTax,
			CommissionTaxRate: invoice.CommissionTaxRate,
			ConsumptionTax:    invoice.ConsumptionTax,
			TaxRate:           invoice.TaxRate,
			BillingAmount:     invoice.BillingAmount,
			PaymentDate:       invoice.PaymentDate,
			Status:            string(invoice.Status),
		})
	}

	return res, nil
}
