package controllers

import (
	"net/http"

	"github.com/rafaLino/couple-wishes-api/api/common"
)

type UserRouter struct {
	routes []common.Route
}

func NewUserRouter() common.Bundle {
	ctrl := NewUserController()

	r := []common.Route{
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: (ctrl.ProtectedHandler(ctrl.GetAll)),
		},
		{
			Method:  http.MethodGet,
			Path:    `/users/{id}`,
			Handler: ctrl.ProtectedHandler(ctrl.Get),
		},
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: (ctrl.Create),
		},
		{
			Method:  http.MethodPut,
			Path:    `/users/{id}`,
			Handler: ctrl.ProtectedHandler(ctrl.Update),
		},
		{
			Method:  http.MethodDelete,
			Path:    `/users/{id}`,
			Handler: ctrl.ProtectedHandler(ctrl.Delete),
		},
		{
			Method:  http.MethodPost,
			Path:    `/users/checkusername/{username}`,
			Handler: ctrl.ProtectedHandler(ctrl.CheckUsername),
		},
		{
			Method:  http.MethodPost,
			Path:    `/users/updatepassword`,
			Handler: ctrl.ProtectedHandler(ctrl.UpdatePassword),
		},
		{
			Method:  http.MethodPost,
			Path:    `/users/couple`,
			Handler: ctrl.ProtectedHandler(ctrl.CreateCouple),
		},
		{
			Method:  http.MethodDelete,
			Path:    `/users/couple/{id}`,
			Handler: ctrl.ProtectedHandler(ctrl.DeleteCouple),
		},
		{
			Method:  http.MethodPost,
			Path:    `/users/login`,
			Handler: ctrl.Login,
		},
		{
			Method:  http.MethodPost,
			Path:    `/users/refresh`,
			Handler: ctrl.RefreshToken,
		},
	}

	return &UserRouter{routes: r}
}

func (w *UserRouter) GetRoutes() []common.Route {
	return w.routes
}
