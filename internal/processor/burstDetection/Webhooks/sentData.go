package webhooks

import (
	"log"

	"github.com/DumbNoxx/goxe/internal/options"
)

// sentData sends and HTTP request with JSON data to the specified URL.
//
// Parameters:
//
//   - data: []byte containing the json payload to be sent.
//   - err: previus serialization error (if not nil, the function logs the error and return without sending).
//   - url: the webhook address to which the request will be sent.
//
// Returns:
//
//   - void: the function returns nothing: it logs errors using log.Print.
//
// The function performs:
//
//   - If err is not nil, it prints 'Convert JSON fail' using log.Print and returns immediately.
//   - If err is nil, it calls 'options.SentWebhook(url, data)' to performs the HTTP request.
func sentData(data []byte, err error, url string) {
	if err != nil {
		log.Print("Convert json fail")
		return
	}
	options.SentWebhook(url, data)
}
