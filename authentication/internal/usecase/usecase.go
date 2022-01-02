package usecase

import (
	"github.com/dgrijalva/jwt-go"
	"time"

	auth "podbilling/authentication/internal"
	"podbilling/authentication/model"
)

type AuthUseCase struct {
	SignKey    []byte
	Repository auth.Repository
}

func NewAuthUseCase(repo auth.Repository, signKey []byte) *AuthUseCase {
	return &AuthUseCase{
		Repository: repo,
		SignKey:    signKey,
	}
}

func (uc *AuthUseCase) GenerateJWT(login, password string) (string, error) {

	user, err := uc.Repository.GetUser(
		login,
		password,
	)

	if err != nil {
		return "", err
	}

	claims := model.CustomClaims{
		ID: user.ID,
		Login: user.Login,
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
