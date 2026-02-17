package options

import (
	"bytes"
	"fmt"
	"net/http"
)

// SentWebhook sends an HTTP POST request with the provided payload to the specified URL.
// Parameters:
//   - url: the webhook address to which the request will be sent.
//   - payload: byte slice containing the JSON content to be sent in the request body.
//
// Returns:
//   - error: nil if the request completed successfully and the HTTP status code is between 200 and 299.
//     In case of a request error or status code out of range, it returns a descriptive error.
//
// The function performs:
//
//   - Performs an HTTP POST request using http.Post, with "application/json" as the Content-Type.
//   - If http.Post fails (network error, etc.), it returns a formatted error with fmt.Errorf wrapping the original error.
//   - Ensures the response body is closed using defer resp.Body.Close().
//   - Checks the status code: if it is between 200 and 299 (success), it returns nil.
//   - If the status code indicates an error, it returns an error containing the status code and message.
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
