//go:generate mockgen -source=$GOFILE -package=mock_usecase -destination=mock/mock_usecase/$GOFILE
package server

import (
	"context"
	"time"

	"github.com/genku-m/upsider-cording-test/models"
)

type InvoiceUsecase interface {
	Create(ctx context.Context, companyGUID, CustomerGUID string, publishDate time.Time, payment uint64, commissionTaxRate float64, taxRate float64, paymentDate time.Time) (*models.Invoice, error)
	// 未実装なのでコメントアウト
	// List(ctx context.Context, companyGUID string, firstPaymentDate, lastPaymentDate time.Time) ([]*models.Invoice, error)
}
