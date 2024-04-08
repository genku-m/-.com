package server_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/genku-m/upsider-cording-test/models"
	"github.com/genku-m/upsider-cording-test/server"
	"github.com/genku-m/upsider-cording-test/server/mock/mock_usecase"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateInvoice(t *testing.T) {
	type want struct {
		res *server.InvoiceResponse
		err error
	}

	tests := []struct {
		description string
		args        *gin.Context
		setup       func(mockInvoiceUsecase *mock_usecase.MockInvoiceUsecase)
		want        want
	}{
		{
			description: "正常系",
			args: func() *gin.Context {
				ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
				body := bytes.NewBufferString("{\"company_guid\": \"company_guid\",\"customer_guid\": \"customer_guid\",\"publish_date\": \"2024-04-01T00:00:00Z\",\"payment\": 10000,\"commission_tax_rate\": 0.04,\"tax_rate\": 0.1,\"payment_date\": \"2024-04-05T00:00:00Z\"}")
				req, _ := http.NewRequest("POST", "/api/invoices", body)
				ginContext.Request = req
				return ginContext
			}(),
			setup: func(mockInvoiceUsecase *mock_usecase.MockInvoiceUsecase) {
				mockInvoiceUsecase.EXPECT().Create(
					gomock.Any(),
					"company_guid",
					"customer_guid",
					time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
					uint64(10000),
					0.04,
					0.10,
					time.Date(2024, 4, 5, 0, 0, 0, 0, time.UTC),
				).Return(&models.Invoice{
					GUID:              "guid",
					CompanyGUID:       "company_guid",
					CustomerGUID:      "customer_guid",
					PublishDate:       time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
					Payment:           10000,
					CommissionTax:     400,
					CommissionTaxRate: 0.04,
					ConsumptionTax:    40,
					TaxRate:           0.10,
					BillingAmount:     10440,
					PaymentDate:       time.Date(2024, 4, 5, 0, 0, 0, 0, time.UTC),
					Status:            models.InvoiceStatusUnprocessed,
				}, nil)
			},
			want: want{
				res: &server.InvoiceResponse{
					GUID:              "guid",
					CompanyGUID:       "company_guid",
					CustomerGUID:      "customer_guid",
					PublishDate:       time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
					Payment:           10000,
					CommissionTax:     400,
					CommissionTaxRate: 0.04,
					ConsumptionTax:    40,
					TaxRate:           0.10,
					BillingAmount:     10440,
					PaymentDate:       time.Date(2024, 4, 5, 0, 0, 0, 0, time.UTC),
					Status:            string(models.InvoiceStatusUnprocessed),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)

			mockInvoiceUsecase := mock_usecase.NewMockInvoiceUsecase(ctrl)
			tt.setup(mockInvoiceUsecase)
			server := server.NewServer(mockInvoiceUsecase, &server.ServerConfig{})
			res, err := server.CreateInvoice(tt.args)
			if err != nil {
				assert.Equal(t, tt.want.err, err)
			}
			assert.Equal(t, tt.want.res, res)
		})
	}
}

func TestListInvoice(t *testing.T) {
	type want struct {
		res []*server.InvoiceResponse
		err error
	}

	tests := []struct {
		description string
		args        *gin.Context
		setup       func(mockInvoiceUsecase *mock_usecase.MockInvoiceUsecase)
		want        want
	}{
		{
			description: "正常系",
			args: func() *gin.Context {
				ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
				body := bytes.NewBufferString("{\"company_guid\": \"company_guid\",\"first_payment_date\": \"2024-04-01T00:00:00Z\",\"last_payment_date\": \"2024-04-05T00:00:00Z\"}")
				req, _ := http.NewRequest("GET", "/api/invoices", body)
				ginContext.Request = req
				return ginContext
			}(),
			setup: func(mockInvoiceUsecase *mock_usecase.MockInvoiceUsecase) {
				mockInvoiceUsecase.EXPECT().List(
					gomock.Any(),
					"company_guid",
					time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2024, 4, 5, 0, 0, 0, 0, time.UTC),
				).Return([]*models.Invoice{
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
				res: []*server.InvoiceResponse{
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
						Status:            string(models.InvoiceStatusPaied),
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
						Status:            string(models.InvoiceStatusUnprocessed),
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
						Status:            string(models.InvoiceStatusError),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)

			mockInvoiceUsecase := mock_usecase.NewMockInvoiceUsecase(ctrl)
			tt.setup(mockInvoiceUsecase)
			server := server.NewServer(mockInvoiceUsecase, &server.ServerConfig{})
			res, err := server.ListInvoice(tt.args)
			if err != nil {
				assert.Equal(t, tt.want.err, err)
			}
			assert.Equal(t, tt.want.res, res)
		})
	}
}
