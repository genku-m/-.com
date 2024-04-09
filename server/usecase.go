//go:generate mockgen -source=$GOFILE -package=mock_usecase -destination=mock/mock_usecase/$GOFILE
package server

import (
	"context"
	"time"

	"github.com/genku-m/upsider-cording-test/models"
	"github.com/gin-gonic/gin"
)

type InvoiceUsecase interface {
	Create(ctx context.Context, companyGUID, CustomerGUID string, publishDate time.Time, payment uint64, commissionTaxRate float64, taxRate float64, paymentDate time.Time) (*models.Invoice, error)
	List(ctx context.Context, companyGUID string, firstPaymentDate, lastPaymentDate time.Time) ([]*models.Invoice, error)
}

type AuthUsecase interface {
	Login(ctx *gin.Context, email, password string) error
}
