package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rafaLino/couple-wishes-api/ports"
)

type WhatsAppApiAdapter struct {
	client *http.Client
	ports.WhatsAppApiAdapter
}

func NewWhatsAppApiAdapter() ports.WhatsAppApiAdapter {
	return &WhatsAppApiAdapter{}
}

func (w *WhatsAppApiAdapter) Connect() {
	w.client = &http.Client{Timeout: time.Duration(1) * time.Second}
}

func (w *WhatsAppApiAdapter) MarkAsRead(phone_id string, message_id string) error {
	url := fmt.Sprintf("https://graph.facebook.com/v22.0/%s/messages", phone_id)
	message := map[string]string{
		"messaging_product": "whatsapp",
		"status":            "read",
		"message_id":        message_id,
	}

	body, err := json.Marshal(message)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GRAPH_API_TOKEN")))
	req.Header.Add("Accept", "application/json")

	_, err = w.client.Do(req)

	if err != nil {
		return err
	}

	return nil
}
