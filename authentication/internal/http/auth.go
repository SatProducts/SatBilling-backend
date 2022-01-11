package http

import (
	"encoding/json"
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

	json.NewEncoder(wr).Encode(AuthResponse{Token: token})
}
