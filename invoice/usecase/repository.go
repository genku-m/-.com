//go:generate mockgen -source=$GOFILE -package=mock_repository -destination=mock/mock_repository/$GOFILE
package invoice_usecase

import (
	"context"
	"time"

	"github.com/genku-m/upsider-cording-test/models"
)

type InvoiceRepository interface {
	Create(ctx context.Context, invoice *models.Invoice) error
	List(ctx context.Context, companyGUID string, firstPaymentDate, lastPaymentDate time.Time) ([]*models.Invoice, error)
}
