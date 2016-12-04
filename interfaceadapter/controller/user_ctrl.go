package controller

import (
	"net/http"
	"goclean/usecase"
	"github.com/gorilla/mux"
	"encoding/json"
)

type UserCtrl interface {
	GetUser(w http.ResponseWriter, r *http.Request)
}

func NewUserCtrl(userUsecase usecase.UserUseCase) UserCtrl {
	return &userCtrlImpl{
		userUsecase: userUsecase,
	}
}

type userCtrlImpl struct {
	userUsecase usecase.UserUseCase
}

func (c *userCtrlImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	// Get Uid in query
	vars := mux.Vars(r)
	userId, ok := vars["userId"]

	// Validate request data
	if !ok || userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Call usecase layer to get user
	userEntity, err := c.userUsecase.GetUser(userId)
	if err != nil {
		// Response error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Convert entity data to the new one that we will response to API
	userPresenter := NewUser(userEntity)

	js, err := json.Marshal(userPresenter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
