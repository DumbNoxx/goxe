package webhooks

import (
	"log"

	"github.com/DumbNoxx/goxe/internal/options"
)

func sentData(data []byte, err error, url string) {
	options.SentWebhook(url, data)
	if err != nil {
		log.Print("Convert json fail")
		return
	}
}
