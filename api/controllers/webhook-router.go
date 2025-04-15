package controllers

import (
	"net/http"

	"github.com/rafaLino/couple-wishes-api/api/common"
)

type WebhookRouter struct {
	routes []common.Route
}

func NewWebhookRouter() common.Bundle {
	ctrl := NewWebhookController()

	r := []common.Route{
		{
			Method:  http.MethodGet,
			Path:    "/webhook",
			Handler: ctrl.Get,
		},
		{
			Method:  http.MethodPost,
			Path:    `/webhook`,
			Handler: ctrl.Post,
		},
	}

	return &WebhookRouter{routes: r}
}

func (w *WebhookRouter) GetRoutes() []common.Route {
	return w.routes
}
