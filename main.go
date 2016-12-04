package gocleanarchitecture

import (
	"github.com/gorilla/mux"
	"goclean-architecture/interfaceadapter/controller"
	"goclean-architecture/interfaceadapter/repository"
	"goclean-architecture/usecase"
	"net/http"
)

func main() {

	// create reposilories
	userRepo := repository.NewUserRepo()

	// Create use case
	userUseCase := usecase.NewUserUseCase(userRepo)

	// Create controller
	userCtrl := controller.NewUserCtrl(userUseCase)
	ctrl := controller.NewController(userCtrl)

	r := mux.NewRouter()
	ctrl.Register(r)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
