package invoice_usecase

import (
	"context"
	"time"

	"github.com/genku-m/upsider-cording-test/models"
	"github.com/shopspring/decimal"
)

type invoiceUsecase struct {
	InvoiceRepository InvoiceRepository
	Guid              Guid
}

func NewInvoiceUsecase(guid Guid, invoiceRepo InvoiceRepository) *invoiceUsecase {
	return &invoiceUsecase{
		InvoiceRepository: invoiceRepo,
		Guid:              guid,
	}
}

func (i *invoiceUsecase) Create(ctx context.Context, CompanyGUID, CustomerGUID string, PublishDate time.Time, Payment uint64, CommissionTaxRate, TaxRate float64, PaymentDate time.Time) (*models.Invoice, error) {
	PaymentDecimal := decimal.NewFromInt(int64(Payment))
	CommissionTaxRateDecimal := decimal.NewFromFloat(CommissionTaxRate)
	TaxRateDecimal := decimal.NewFromFloat(TaxRate)
	CommissionTax := PaymentDecimal.Mul(CommissionTaxRateDecimal)
	ConsumptionTax := CommissionTax.Mul(TaxRateDecimal)
	BillingAmount := PaymentDecimal.Add(CommissionTax).Add(ConsumptionTax)
	invoiceModel := &models.Invoice{
		GUID:              i.Guid.New(),
		CompanyGUID:       CompanyGUID,
		CustomerGUID:      CustomerGUID,
		PublishDate:       PublishDate,
		Payment:           Payment,
		CommissionTax:     CommissionTax.BigInt().Uint64(),
		CommissionTaxRate: CommissionTaxRate,
		ConsumptionTax:    ConsumptionTax.BigInt().Uint64(),
		TaxRate:           TaxRate,
		BillingAmount:     BillingAmount.BigInt().Uint64(),
		PaymentDate:       PaymentDate,
		// 既に支払い済みのものを作るユースケースがあるかどうか
		Status: models.InvoiceStatusUnprocessed,
	}
	if err := i.InvoiceRepository.Create(ctx, invoiceModel); err != nil {
		return nil, err
	}

	return invoiceModel, nil
}
