package auth_usecase

import (
	"context"
	"database/sql"
	"fmt"

	errpkg "github.com/genku-m/upsider-cording-test/invoice/errors"
	"github.com/genku-m/upsider-cording-test/models"
)

type User struct {
	ID        uint32 `db:"id"`
	GUID      string `db:"guid"`
	Name      string `db:"name"`
	CompanyID uint32 `db:"company_id"`
	Email     string `db:"email"`
	Password  string `db:"password"`
}

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var User User
	err := r.db.QueryRowContext(ctx, "SELECT * FROM user WHERE email=?", email).Scan(
		&User.ID,
		&User.GUID,
		&User.CompanyID,
		&User.Name,
		&User.Email,
		&User.Password,
	)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, errpkg.NewNotFoundError(fmt.Errorf("user not found: %v err: %v", email, err.Error()))
		default:
			return nil, errpkg.NewInternalError(err)
		}
	}

	var companyGUID string
	err = r.db.QueryRowContext(ctx, "SELECT guid FROM company WHERE id=?", User.CompanyID).Scan(&companyGUID)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, errpkg.NewNotFoundError(fmt.Errorf("user not found: %v err: %v", email, err.Error()))
		default:
			return nil, errpkg.NewInternalError(err)
		}
	}

	return &models.User{
		GUID:        User.GUID,
		Name:        User.Name,
		CompanyGUID: companyGUID,
		Email:       User.Email,
		Password:    User.Password,
	}, nil
}
