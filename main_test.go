package gocleanarchitecture

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"goclean/interfaceadapter/controller"
	"goclean/interfaceadapter/repository"
	"goclean/usecase"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Sample integration to test the whole code
func TestIntegration(t *testing.T) {
	// re-create system
	userRepo := repository.NewUserRepo()
	userUseCase := usecase.NewUserUseCase(userRepo)
	userCtrl := controller.NewUserCtrl(userUseCase)
	ctrl := controller.NewController(userCtrl)

	router := mux.NewRouter()
	ctrl.Register(router)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected response code %v but got %v", 200, w.Code)
	}

	respBody, _ := ioutil.ReadAll(w.Body)
	respData := map[string]string{}
	_ = json.Unmarshal(respBody, &respData)

	if respData["id"] != "1" {
		t.Errorf("Expeccted id %v but got %v", "1", respData["id"])
	}
}
