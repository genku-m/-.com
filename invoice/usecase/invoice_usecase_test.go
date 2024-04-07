package invoice_usecase_test

import (
	"context"
	"testing"
	"time"

	invoice_usecase "github.com/genku-m/upsider-cording-test/invoice/usecase"
	"github.com/genku-m/upsider-cording-test/invoice/usecase/mock/mock_guid"
	"github.com/genku-m/upsider-cording-test/invoice/usecase/mock/mock_repository"
	"github.com/genku-m/upsider-cording-test/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	PublishDate, _ := time.Parse("2006-01-02T15:04:05Z07:00", "2024-04-01T00:00:00Z")
	PaymentDate, _ := time.Parse("2006-01-02T15:04:05Z07:00", "2024-04-05T00:00:00Z")
	type args struct {
		CompanyGUID       string
		CustomerGUID      string
		PublishDate       time.Time
		Payment           uint64
		CommissionTaxRate float64
		TaxRate           float64
		PaymentDate       time.Time
	}
	type want struct {
		err     error
		invoice *models.Invoice
	}
	tests := []struct {
		descrition string
		args       args
		setup      func(mockGuid *mock_guid.MockGuid, MocknvoiceRepository *mock_repository.MockInvoiceRepository)
		want       want
	}{
		{
			descrition: "正常系",
			args: args{
				CompanyGUID:       "company_guid",
				CustomerGUID:      "customer_guid",
				PublishDate:       PublishDate,
				Payment:           10000,
				CommissionTaxRate: 0.04,
				TaxRate:           0.10,
				PaymentDate:       PaymentDate,
			},
			setup: func(mockGuid *mock_guid.MockGuid, repo *mock_repository.MockInvoiceRepository) {
				mockGuid.EXPECT().New().Return("guid")
				repo.EXPECT().Create(gomock.Any(), &models.Invoice{
					GUID:              "guid",
					CompanyGUID:       "company_guid",
					CustomerGUID:      "customer_guid",
					PublishDate:       PublishDate,
					Payment:           10000,
					CommissionTax:     400,
					CommissionTaxRate: 0.04,
					ConsumptionTax:    40,
					TaxRate:           0.10,
					BillingAmount:     10440,
					PaymentDate:       PaymentDate,
					Status:            models.InvoiceStatusUnprocessed,
				}).Return(nil)
			},
			want: want{
				err: nil,
				invoice: &models.Invoice{
					GUID:              "guid",
					CompanyGUID:       "company_guid",
					CustomerGUID:      "customer_guid",
					PublishDate:       PublishDate,
					Payment:           10000,
					CommissionTax:     400,
					CommissionTaxRate: 0.04,
					ConsumptionTax:    40,
					TaxRate:           0.10,
					BillingAmount:     10440,
					PaymentDate:       PaymentDate,
					Status:            models.InvoiceStatusUnprocessed,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.descrition, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)

			repo := mock_repository.NewMockInvoiceRepository(ctrl)
			guid := mock_guid.NewMockGuid(ctrl)
			if tt.setup != nil {
				tt.setup(guid, repo)
			}
			uc := invoice_usecase.NewInvoiceUsecase(guid, repo)
			res, err := uc.Create(context.Background(), tt.args.CompanyGUID, tt.args.CustomerGUID, tt.args.PublishDate, tt.args.Payment, tt.args.CommissionTaxRate, tt.args.TaxRate, tt.args.PaymentDate)
			if err != nil {
				assert.Equal(t, tt.want.err, err)
			}
			assert.Equal(t, tt.want.invoice, res)
		})
	}
}
