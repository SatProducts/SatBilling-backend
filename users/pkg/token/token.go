package token

import (
	"github.com/golang-jwt/jwt"
	"podbilling/users/model"
	"net/http"
	"strings"
)

func ParseToken(unparsedToken string, key []byte) (model.CustomClaims, error) {

	var claims model.CustomClaims

	_, err := jwt.ParseWithClaims(
		unparsedToken,
		&claims,
		func(t *jwt.Token) (interface{}, error) {
			return key, nil
		},
	)

	if err != nil {
		return model.CustomClaims{}, err
	}

	return claims, nil
}

func ParseFromRequest(key []byte, req *http.Request) (model.CustomClaims, error) {

	parsedClaims, err := ParseToken(
		strings.Split(req.Header.Get("Authorization"), "Bearer ")[1],
		key,
	)

	if err != nil {
		return model.CustomClaims{}, err
	}

	return parsedClaims, nil
}
