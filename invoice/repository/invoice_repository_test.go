package repository_test

import (
	"context"
	"testing"
	"time"

	errpkg "github.com/genku-m/upsider-cording-test/invoice/errors"
	"github.com/genku-m/upsider-cording-test/invoice/repository"
	"github.com/genku-m/upsider-cording-test/models"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	publishDate, _ := time.Parse(time.RFC3339, "2024-04-01T00:00:00Z")
	paymentDate, _ := time.Parse(time.RFC3339, "2024-04-05T00:00:00Z")
	type args struct {
		invoice *models.Invoice
	}
	type want struct {
		wantErr errpkg.ErrCode
	}
	tests := []struct {
		description string
		args        args
		want        want
	}{
		{
			description: "正常系",
			args: args{
				invoice: &models.Invoice{
					GUID:              "new_guid",
					CompanyGUID:       "company-1",
					CustomerGUID:      "customer-1",
					PublishDate:       publishDate,
					Payment:           1000,
					CommissionTax:     100,
					CommissionTaxRate: 0.1,
					ConsumptionTax:    10,
					TaxRate:           0.01,
					BillingAmount:     1110,
					PaymentDate:       paymentDate,
					Status:            models.InvoiceStatusUnprocessed,
				},
			},
		},
		{
			description: "cumpanyGUIDが不適切な場合",
			args: args{
				invoice: &models.Invoice{
					GUID:              "new_guid",
					CompanyGUID:       "not-found-company",
					CustomerGUID:      "customer-1",
					PublishDate:       publishDate,
					Payment:           1000,
					CommissionTax:     100,
					CommissionTaxRate: 0.1,
					ConsumptionTax:    10,
					TaxRate:           0.01,
					BillingAmount:     1110,
					PaymentDate:       paymentDate,
					Status:            models.InvoiceStatusUnprocessed,
				},
			},
			want: want{
				wantErr: errpkg.ErrInvalidArgument,
			},
		},
		{
			description: "custemerGUIDが存在しない場合",
			args: args{
				invoice: &models.Invoice{
					GUID:              "new_guid",
					CompanyGUID:       "company-1",
					CustomerGUID:      "not-found-customer",
					PublishDate:       publishDate,
					Payment:           1000,
					CommissionTax:     100,
					CommissionTaxRate: 0.1,
					ConsumptionTax:    10,
					TaxRate:           0.01,
					BillingAmount:     1110,
					PaymentDate:       paymentDate,
					Status:            models.InvoiceStatusUnprocessed,
				},
			},
			want: want{
				wantErr: errpkg.ErrNotFound,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			prepareTestDatabase()
			r := repository.NewInvoiceRepository(testDB)
			if err := r.Create(context.Background(), tt.args.invoice); err != nil {
				serverError, ok := err.(*errpkg.ServerError)
				if !ok {
					t.FailNow()
				}
				assert.Equal(t, tt.want.wantErr, serverError.ErrCode)
			} else {
				invoice := getInvoice(t, tt.args.invoice.GUID)
				assert.Equal(t, tt.args.invoice, invoice)
			}
		})
	}
}

func getInvoice(t *testing.T, guid string) *models.Invoice {
	var invoice repository.Invoice
	err := testDB.QueryRow("SELECT * FROM invoice WHERE guid=?", guid).Scan(
		&invoice.ID,
		&invoice.GUID,
		&invoice.CompanyID,
		&invoice.CustomerID,
		&invoice.PublishDate,
		&invoice.Payment,
		&invoice.CommissionTax,
		&invoice.CommissionTaxRate,
		&invoice.ConsumptionTax,
		&invoice.TaxRate,
		&invoice.BillingAmount,
		&invoice.PaymentDate,
		&invoice.Status,
	)
	if err != nil {
		t.Fatal(err)
	}

	var customerGUID, companyGUID string
	err = testDB.QueryRow("SELECT guid FROM customer WHERE id=?", invoice.CustomerID).Scan(&customerGUID)
	if err != nil {
		t.Fatal(err)
	}

	err = testDB.QueryRow("SELECT guid FROM company WHERE id=?", invoice.CompanyID).Scan(&companyGUID)
	if err != nil {
		t.Fatal(err)
	}

	return &models.Invoice{
		GUID:              invoice.GUID,
		CompanyGUID:       companyGUID,
		CustomerGUID:      customerGUID,
		PublishDate:       invoice.PublishDate,
		Payment:           invoice.Payment,
		CommissionTax:     invoice.CommissionTax,
		CommissionTaxRate: invoice.CommissionTaxRate,
		ConsumptionTax:    invoice.ConsumptionTax,
		TaxRate:           invoice.TaxRate,
		BillingAmount:     invoice.BillingAmount,
		PaymentDate:       invoice.PaymentDate,
		Status:            models.InvoiceStatusUnprocessed,
	}
}

func TestList(t *testing.T) {
	firstPaymentDate, _ := time.Parse(time.RFC3339, "2024-04-01T00:00:00Z")
	lastPaymentDate, _ := time.Parse(time.RFC3339, "2024-04-08T00:00:00Z")
	notFoundFirstPaymentDate, _ := time.Parse(time.RFC3339, "2022-04-01T00:00:00Z")
	notFoundLastPaymentDate, _ := time.Parse(time.RFC3339, "2022-04-08T00:00:00Z")
	publishDate1, _ := time.Parse(time.RFC3339, "2024-04-05T00:00:00Z")
	paymentDate1, _ := time.Parse(time.RFC3339, "2024-04-06T00:00:00Z")
	publishDate2, _ := time.Parse(time.RFC3339, "2024-04-01T00:00:00Z")
	paymentDate2, _ := time.Parse(time.RFC3339, "2024-04-02T00:00:00Z")
	type args struct {
		companyGUID      string
		firstPaymentDate time.Time
		lastPaymentDate  time.Time
	}
	type want struct {
		invoices []*models.Invoice
		wantErr  errpkg.ErrCode
	}
	tests := []struct {
		description string
		args        args
		want        want
	}{
		{
			description: "正常系",
			args: args{
				companyGUID:      "company-1",
				firstPaymentDate: firstPaymentDate,
				lastPaymentDate:  lastPaymentDate,
			},
			want: want{
				invoices: []*models.Invoice{
					{
						GUID:              "invoice-1",
						CompanyGUID:       "company-1",
						CustomerGUID:      "customer-1",
						PublishDate:       publishDate1,
						Payment:           100000,
						CommissionTax:     1000,
						CommissionTaxRate: 0.01,
						ConsumptionTax:    1000,
						TaxRate:           0.01,
						BillingAmount:     102000,
						PaymentDate:       paymentDate1,
						Status:            models.InvoiceStatusUnprocessed,
					},
					{
						GUID:              "invoice-2",
						CompanyGUID:       "company-1",
						CustomerGUID:      "customer-1",
						PublishDate:       publishDate2,
						Payment:           100000,
						CommissionTax:     1000,
						CommissionTaxRate: 0.01,
						ConsumptionTax:    1000,
						TaxRate:           0.01,
						BillingAmount:     102000,
						PaymentDate:       paymentDate2,
						Status:            models.InvoiceStatusProcessing,
					},
				},
			},
		},
		{
			description: "companyGUIDが存在しない場合",
			args: args{
				companyGUID:      "not-found-company",
				firstPaymentDate: firstPaymentDate,
				lastPaymentDate:  lastPaymentDate,
			},
			want: want{
				wantErr: errpkg.ErrNotFound,
			},
		},
		{
			description: "invoiceが存在しない場合",
			args: args{
				companyGUID:      "company-1",
				firstPaymentDate: notFoundFirstPaymentDate,
				lastPaymentDate:  notFoundLastPaymentDate,
			},
			want: want{
				invoices: []*models.Invoice{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			prepareTestDatabase()
			r := repository.NewInvoiceRepository(testDB)
			invoices, err := r.List(context.Background(), tt.args.companyGUID, tt.args.firstPaymentDate, tt.args.lastPaymentDate)
			if err != nil {
				serverError, ok := err.(*errpkg.ServerError)
				if !ok {
					t.FailNow()
				}
				assert.Equal(t, tt.want.wantErr, serverError.ErrCode)
			} else {
				assert.Equal(t, tt.want.invoices, invoices)
			}
		})
	}
}
