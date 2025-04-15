package controllers

import (
	"net/http"
	"os"

	"github.com/golobby/container/v3"
	"github.com/rafaLino/couple-wishes-api/api/common"
	"github.com/rafaLino/couple-wishes-api/entities"
	"github.com/rafaLino/couple-wishes-api/ports"
)

type WebhookController struct {
	common.Controller
}

func NewWebhookController() *WebhookController {
	return &WebhookController{}
}

var WEBHOOK_VERIFY_TOKEN = os.Getenv("WEBHOOK_VERIFY_TOKEN")

func (c *WebhookController) Get(w http.ResponseWriter, r *http.Request) {
	mode := c.GetQuery(r, "hub.mode")
	token := c.GetQuery(r, "hub.verify_token")
	challenge := c.GetQuery(r, "hub.challenge")

	if mode == "subscribe" && token == WEBHOOK_VERIFY_TOKEN {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(challenge))
		return
	}

	w.WriteHeader(http.StatusForbidden)
}

func (c *WebhookController) Post(w http.ResponseWriter, r *http.Request) {
	var service ports.IWishService
	container.Resolve(&service)

	var input entities.WhatsAppMessage
	err := c.GetContent(&input, r)

	if err != nil {
		c.SendError(nil, http.StatusBadRequest, w)
		return
	}

	err = service.CreateFromWhatsApp(input)

	if err != nil {
		c.SendError(err, http.StatusBadRequest, w)
		return
	}

	c.SendJSON(w, nil, http.StatusOK)
}
