package controllers

import (
	"net/http"

	"github.com/rafaLino/couple-wishes-api/api/common"
)

type WishRouter struct {
	routes []common.Route
}

func NewWishRouter() common.Bundle {
	ctrl := NewWishController()

	r := []common.Route{
		{
			Method:  http.MethodGet,
			Path:    "/wishes",
			Handler: ctrl.ProtectedHandler(ctrl.GetAll),
		},
		{
			Method:  http.MethodGet,
			Path:    `/wishes/{id}`,
			Handler: ctrl.ProtectedHandler(ctrl.Get),
		},
		{
			Method:  http.MethodPost,
			Path:    "/wishes",
			Handler: ctrl.ProtectedHandler(ctrl.Save),
		},
		{
			Method:  http.MethodPut,
			Path:    `/wishes/{id}`,
			Handler: ctrl.ProtectedHandler(ctrl.Update),
		},
		{
			Method:  http.MethodDelete,
			Path:    `/wishes/{id}`,
			Handler: ctrl.ProtectedHandler(ctrl.Delete),
		},
		{
			Method:  http.MethodPost,
			Path:    `/wishes/byurl`,
			Handler: ctrl.ProtectedHandler(ctrl.Create),
		},
	}

	return &WishRouter{routes: r}
}

func (w *WishRouter) GetRoutes() []common.Route {
	return w.routes
}
