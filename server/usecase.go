//go:generate mockgen -source=$GOFILE -package=mock_usecase -destination=mock/mock_usecase/$GOFILE
package server

import (
	"context"
	"time"

	"github.com/genku-m/upsider-cording-test/models"
)

type InvoiceUsecase interface {
	Create(ctx context.Context, PublishDate time.Time, Payment uint64, CommissionTaxRate float64, TaxRate float64, PaymentDate time.Time) (*models.Invoice, error)
	List(ctx context.Context, companyGUID string, firstPaymentDate, lastPaymentDate time.Time) ([]*models.Invoice, error)
}
