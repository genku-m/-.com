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

func (i *invoiceUsecase) Create(ctx context.Context, companyGUID, customerGUID string, publishDate time.Time, payment uint64, commissionTaxRate, taxRate float64, paymentDate time.Time) (*models.Invoice, error) {
	paymentDecimal := decimal.NewFromInt(int64(payment))
	commissionTaxRateDecimal := decimal.NewFromFloat(commissionTaxRate)
	taxRateDecimal := decimal.NewFromFloat(taxRate)
	commissionTax := paymentDecimal.Mul(commissionTaxRateDecimal)
	consumptionTax := commissionTax.Mul(taxRateDecimal)
	billingAmount := paymentDecimal.Add(commissionTax).Add(consumptionTax)
	invoiceModel := &models.Invoice{
		GUID:              i.Guid.Generate(),
		CompanyGUID:       companyGUID,
		CustomerGUID:      customerGUID,
		PublishDate:       publishDate,
		Payment:           payment,
		CommissionTax:     commissionTax.BigInt().Uint64(),
		CommissionTaxRate: commissionTaxRate,
		ConsumptionTax:    consumptionTax.BigInt().Uint64(),
		TaxRate:           taxRate,
		BillingAmount:     billingAmount.BigInt().Uint64(),
		PaymentDate:       paymentDate,
		// 既に支払い済みのものを作るユースケースがあるかどうか
		Status: models.InvoiceStatusUnprocessed,
	}
	if err := i.InvoiceRepository.Create(ctx, invoiceModel); err != nil {
		return nil, err
	}

	return invoiceModel, nil
}

func (i *invoiceUsecase) List(ctx context.Context, companyGUID string, firstPaymentDate, lastPaymentDate time.Time) ([]*models.Invoice, error) {
	return i.InvoiceRepository.List(ctx, companyGUID, firstPaymentDate, lastPaymentDate)
}
