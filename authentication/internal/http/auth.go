package http

import (
	"encoding/json"
	common "podbilling/authentication/pkg/common"
	"net/http"
)

func (h *AuthHandler) Login(wr http.ResponseWriter, req *http.Request) {

	var info AuthRequest

	err := json.NewDecoder(req.Body).Decode(&info)

	if err != nil {
		http.Error(wr, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.UseCase.GenerateJWT(info.Login, info.Password)

	if err != nil {
		http.Error(wr, err.Error(), http.StatusUnauthorized)
		return
	}

	common.ServeJSON(wr, AuthResponse{Token: token})
}
