package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Controller interface {
	Register(router *mux.Router)
}

func NewController(userCtrl UserCtrl) Controller {
	return &controllerImpl{
		mdw: NewMdw(),
		userCtrl: userCtrl,
	}
}

type controllerImpl struct {
	userCtrl UserCtrl
	mdw      Mdw
}

func (c *controllerImpl) Register(r *mux.Router) {
	mdw := c.mdw

	r.Path("/users/{userId}").Methods("GET").Handler(
		mdw.CORS(mdw.Log(mdw.Header(http.HandlerFunc(c.userCtrl.GetUser)))),
	)
}