package http

import (
	"strconv"
	"encoding/json"
	"podbilling/users/model"
	"net/http"
	"github.com/gorilla/mux"
	"podbilling/users/internal"
)

func (h *UsersHandler) GetSelf(user model.CustomClaims, wr http.ResponseWriter, req *http.Request) {

	userModel, err := h.UseCase.Get(user.ID)

	if err != nil {
		http.Error(wr, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(wr).Encode(userModel)
}

func (h *UsersHandler) Get(user model.CustomClaims, wr http.ResponseWriter, req *http.Request) {

	if user.Permissions != model.ADMINISTRATOR {
		http.Error(wr, users.NoPermissionsError.Error(), http.StatusForbidden)
		return
	}

	userID, _ := strconv.Atoi(mux.Vars(req)["id"])

	userModel, err := h.UseCase.Get(uint(userID))

	if err != nil {
		http.Error(wr, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(wr).Encode(userModel)
}

func (h *UsersHandler) Create(user model.CustomClaims, wr http.ResponseWriter, req *http.Request) {

	if user.Permissions != model.ADMINISTRATOR {
		http.Error(wr, users.NoPermissionsError.Error(), http.StatusForbidden)
		return
	}

	var newUser NewUserRequest

	err := json.NewDecoder(req.Body).Decode(&newUser)

	if err != nil {
		http.Error(wr, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.UseCase.Create(
		newUser.Login,
		newUser.Password,
		newUser.Permissions,
	)

	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	wr.WriteHeader(http.StatusCreated)
}

func (h *UsersHandler) Update(user model.CustomClaims, wr http.ResponseWriter, req *http.Request) {

	if user.Permissions != model.ADMINISTRATOR {
		http.Error(wr, users.NoPermissionsError.Error(), http.StatusForbidden)
		return
	}

	userID, _ := strconv.Atoi(mux.Vars(req)["id"])

	userModel, err := h.UseCase.Get(uint(userID))

	if err != nil {
		http.Error(wr, err.Error(), http.StatusNotFound)
		return
	}

	if err = json.NewDecoder(req.Body).Decode(&userModel); err != nil {
		http.Error(wr, err.Error(), http.StatusBadRequest)
		return
	}

	if err = h.UseCase.Update(userModel); err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	wr.WriteHeader(http.StatusOK)
}

func (h *UsersHandler) Delete(user model.CustomClaims, wr http.ResponseWriter, req *http.Request) {

	if user.Permissions != model.ADMINISTRATOR {
		http.Error(wr, users.NoPermissionsError.Error(), http.StatusForbidden)
		return
	}
	
	userID, _ := strconv.Atoi(mux.Vars(req)["id"])

	_, err := h.UseCase.Get(uint(userID))

	if err != nil {
		http.Error(wr, err.Error(), http.StatusNotFound)
		return
	}

	if err = h.UseCase.Delete(uint(userID)); err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	wr.WriteHeader(http.StatusOK)
}

func (h *UsersHandler) GetWorkers(user model.CustomClaims, wr http.ResponseWriter, req *http.Request) {

	if user.Permissions != model.ADMINISTRATOR {
		http.Error(wr, users.NoPermissionsError.Error(), http.StatusForbidden)
		return
	}

	workers, err := h.UseCase.GetWorkers()

	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(wr).Encode(workers)
}