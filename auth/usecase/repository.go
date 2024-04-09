package auth_usecase

import (
	"context"

	"github.com/genku-m/upsider-cording-test/models"
)

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}
