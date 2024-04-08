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
		companyGUID       string
		customerGUID      string
		publishDate       time.Time
		payment           uint64
		commissionTaxRate float64
		taxRate           float64
		paymentDate       time.Time
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
				companyGUID:       "company_guid",
				customerGUID:      "customer_guid",
				publishDate:       PublishDate,
				payment:           10000,
				commissionTaxRate: 0.04,
				taxRate:           0.10,
				paymentDate:       PaymentDate,
			},
			setup: func(mockGuid *mock_guid.MockGuid, repo *mock_repository.MockInvoiceRepository) {
				mockGuid.EXPECT().Generate().Return("guid")
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
			res, err := uc.Create(context.Background(), tt.args.companyGUID, tt.args.customerGUID, tt.args.publishDate, tt.args.payment, tt.args.commissionTaxRate, tt.args.taxRate, tt.args.paymentDate)
			if err != nil {
				assert.Equal(t, tt.want.err, err)
			}
			assert.Equal(t, tt.want.invoice, res)
		})
	}
}

func TestList(t *testing.T) {
	type args struct {
		companyGUID                       string
		firstPaymentDate, lastPaymentDate time.Time
	}
	type want struct {
		err     error
		invoice []*models.Invoice
	}
	tests := []struct {
		descrition string
		args       args
		setup      func(MocknvoiceRepository *mock_repository.MockInvoiceRepository)
		want       want
	}{
		{
			descrition: "正常系",
			args: args{
				companyGUID:      "company_guid",
				firstPaymentDate: time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
				lastPaymentDate:  time.Date(2024, 4, 5, 0, 0, 0, 0, time.UTC),
			},
			setup: func(repo *mock_repository.MockInvoiceRepository) {
				repo.EXPECT().List(gomock.Any(),
					"company_guid",
					time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2024, 4, 5, 0, 0, 0, 0, time.UTC),
				).Return(
					[]*models.Invoice{
						{
							GUID:              "guid1",
							CompanyGUID:       "company_guid1",
							CustomerGUID:      "customer_guid1",
							PublishDate:       time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
							Payment:           10000,
							CommissionTax:     400,
							CommissionTaxRate: 0.04,
							ConsumptionTax:    40,
							TaxRate:           0.10,
							BillingAmount:     10440,
							PaymentDate:       time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
							Status:            models.InvoiceStatusPaied,
						},
						{
							GUID:              "guid2",
							CompanyGUID:       "company_guid2",
							CustomerGUID:      "customer_guid2",
							PublishDate:       time.Date(2024, 4, 2, 0, 0, 0, 0, time.UTC),
							Payment:           10000,
							CommissionTax:     400,
							CommissionTaxRate: 0.04,
							ConsumptionTax:    40,
							TaxRate:           0.10,
							BillingAmount:     10440,
							PaymentDate:       time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
							Status:            models.InvoiceStatusUnprocessed,
						},
						{
							GUID:              "guid3",
							CompanyGUID:       "company_guid3",
							CustomerGUID:      "customer_guid3",
							PublishDate:       time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
							Payment:           10000,
							CommissionTax:     400,
							CommissionTaxRate: 0.04,
							ConsumptionTax:    40,
							TaxRate:           0.10,
							BillingAmount:     10440,
							PaymentDate:       time.Date(2024, 4, 5, 0, 0, 0, 0, time.UTC),
							Status:            models.InvoiceStatusError,
						},
					}, nil)
			},
			want: want{
				invoice: []*models.Invoice{
					{
						GUID:              "guid1",
						CompanyGUID:       "company_guid1",
						CustomerGUID:      "customer_guid1",
						PublishDate:       time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
						Payment:           10000,
						CommissionTax:     400,
						CommissionTaxRate: 0.04,
						ConsumptionTax:    40,
						TaxRate:           0.10,
						BillingAmount:     10440,
						PaymentDate:       time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
						Status:            models.InvoiceStatusPaied,
					},
					{
						GUID:              "guid2",
						CompanyGUID:       "company_guid2",
						CustomerGUID:      "customer_guid2",
						PublishDate:       time.Date(2024, 4, 2, 0, 0, 0, 0, time.UTC),
						Payment:           10000,
						CommissionTax:     400,
						CommissionTaxRate: 0.04,
						ConsumptionTax:    40,
						TaxRate:           0.10,
						BillingAmount:     10440,
						PaymentDate:       time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
						Status:            models.InvoiceStatusUnprocessed,
					},
					{
						GUID:              "guid3",
						CompanyGUID:       "company_guid3",
						CustomerGUID:      "customer_guid3",
						PublishDate:       time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
						Payment:           10000,
						CommissionTax:     400,
						CommissionTaxRate: 0.04,
						ConsumptionTax:    40,
						TaxRate:           0.10,
						BillingAmount:     10440,
						PaymentDate:       time.Date(2024, 4, 5, 0, 0, 0, 0, time.UTC),
						Status:            models.InvoiceStatusError,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.descrition, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)

			repo := mock_repository.NewMockInvoiceRepository(ctrl)
			if tt.setup != nil {
				tt.setup(repo)
			}
			uc := invoice_usecase.NewInvoiceUsecase(nil, repo)
			res, err := uc.List(context.Background(), tt.args.companyGUID, tt.args.firstPaymentDate, tt.args.lastPaymentDate)
			if err != nil {
				assert.Equal(t, tt.want.err, err)
			}
			assert.Equal(t, tt.want.invoice, res)
		})
	}
}
