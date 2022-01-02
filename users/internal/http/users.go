package http

import (
	"strconv"
	"encoding/json"
	"podbilling/users/model"
	common "podbilling/users/pkg/common"
	"net/http"
	"github.com/gorilla/mux"
	"podbilling/users/internal"
)

func (h *UsersHandler) GetSelfUser(user model.CustomClaims, wr http.ResponseWriter, req *http.Request) {

	userModel, err := h.UseCase.GetUser(user.ID)

	if err != nil {
		http.Error(wr, err.Error(), http.StatusNotFound)
		return
	}

	common.ServeJSON(wr, userModel)
}

func (h *UsersHandler) GetUserByID(user model.CustomClaims, wr http.ResponseWriter, req *http.Request) {

	if user.Permissions != model.ADMINISTRATOR {
		http.Error(wr, users.NoPermissionsError.Error(), http.StatusForbidden)
		return
	}

	userID, _ := strconv.Atoi(mux.Vars(req)["id"])

	userModel, err := h.UseCase.GetUser(uint(userID))

	if err != nil {
		http.Error(wr, err.Error(), http.StatusNotFound)
		return
	}

	common.ServeJSON(wr, userModel)
}

func (h *UsersHandler) CreateUser(user model.CustomClaims, wr http.ResponseWriter, req *http.Request) {

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

	err = h.UseCase.CreateUser(
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

func (h *UsersHandler) UpdateUser(user model.CustomClaims, wr http.ResponseWriter, req *http.Request) {

	if user.Permissions != model.ADMINISTRATOR {
		http.Error(wr, users.NoPermissionsError.Error(), http.StatusForbidden)
		return
	}

	userID, _ := strconv.Atoi(mux.Vars(req)["id"])

	userModel, err := h.UseCase.GetUser(uint(userID))

	if err != nil {
		http.Error(wr, err.Error(), http.StatusNotFound)
		return
	}

	if err = json.NewDecoder(req.Body).Decode(&userModel); err != nil {
		http.Error(wr, err.Error(), http.StatusBadRequest)
		return
	}

	if err = h.UseCase.UpdateUser(userModel); err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	wr.WriteHeader(http.StatusOK)
}

func (h *UsersHandler) DeleteUser(user model.CustomClaims, wr http.ResponseWriter, req *http.Request) {

	if user.Permissions != model.ADMINISTRATOR {
		http.Error(wr, users.NoPermissionsError.Error(), http.StatusForbidden)
		return
	}
	
	userID, _ := strconv.Atoi(mux.Vars(req)["id"])

	_, err := h.UseCase.GetUser(uint(userID))

	if err != nil {
		http.Error(wr, err.Error(), http.StatusNotFound)
		return
	}

	if err = h.UseCase.DeleteUser(uint(userID)); err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	wr.WriteHeader(http.StatusOK)
}

func (h *UsersHandler) GetAllWorkers(user model.CustomClaims, wr http.ResponseWriter, req *http.Request) {

	if user.Permissions != model.ADMINISTRATOR {
		http.Error(wr, users.NoPermissionsError.Error(), http.StatusForbidden)
		return
	}

	workers, err := h.UseCase.GetAllWorkers()

	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	common.ServeJSON(wr, workers)
}