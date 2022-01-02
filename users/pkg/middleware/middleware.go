package middleware

import (
	"net/http"
	"podbilling/users/model"
	token "podbilling/users/pkg/token"
)

type PrivatePageHandler func(user model.CustomClaims, wr http.ResponseWriter, req *http.Request)

func JwtAuthMW(key []byte, handler PrivatePageHandler) http.HandlerFunc {

	return func(wr http.ResponseWriter, req *http.Request) {

		user, err := token.ParseFromRequest(key, req)

		if err != nil {
			http.Error(wr, err.Error(), http.StatusUnauthorized)
			return
		}
		
		handler(user, wr, req)
	}
}