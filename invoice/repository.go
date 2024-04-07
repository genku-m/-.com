//go:generate mockgen -source=$GOFILE -package=mock_invoice_usecase -destination=mock/$GOPACKAGE/mock_$GOFILE
package invoice

import (
	"context"
	"time"

	"github.com/genku-m/upsider-cording-test/models"
)

type Repository interface {
	Create(ctx context.Context, invoice models.Invoice) error
	List(ctx context.Context, companyGUID string, firstPaymentDate, lastPaymentDate time.Time) ([]models.Invoice, error)
}
