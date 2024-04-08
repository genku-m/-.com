//go:generate mockgen -source=$GOFILE -package=mock_guid -destination=mock/mock_guid/$GOFILE
package invoice_usecase

type Guid interface {
	Generate() string
}
