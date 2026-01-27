package options

import (
	"bytes"
	"fmt"
	"net/http"
)

func SentWebhook(url string, payload []byte) error {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to send webhook: %w", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	return fmt.Errorf("webhook delivery failed with status: %s", resp.Status)
}
