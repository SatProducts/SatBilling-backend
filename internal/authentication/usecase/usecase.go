package usecase

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"

	auth "podbilling/internal/authentication"
	"podbilling/model"
)

type UseCase struct {
	SignKey    []byte
	Repository auth.Repository
}

func NewUseCase(repo auth.Repository, signKey []byte) *UseCase {
	return &UseCase{
		Repository: repo,
		SignKey:    signKey,
	}
}

func (uc *UseCase) GenerateJWT(ctx context.Context, login, password string) (string, error) {

	user, err := uc.Repository.Get(
		login,
		password,
	)

	if err != nil {
		return "", err
	}

	claims := model.Claims{
		ID:          user.ID,
		Login:       user.Login,
		Permissions: user.Permissions,

		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1000).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := t.SignedString(uc.SignKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
