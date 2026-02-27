package integrations

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/DumbNoxx/goxe/internal/options"
	"github.com/DumbNoxx/goxe/pkg/exporter"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

// Integrations iterates through configured external services to send aggregated log data
// via HTTP POST requests.
//
// Parameters:
//
//   - logs: map containing log statistics grouped by source and message.
//
//   - Shipper: an interface of type exporter.Shipper used to format the data before sending.
//
// The function performs:
//
//   - Initializes an http.Client with a 10-second timeout.
//
//   - Iterates through the 'options.Config.Integrations' slice.
//
//   - For each integration, checks if 'OnAggregation' is enabled:
//
//     -Calls Shipper.PrepareShip(logs) to get the formatted log data.
//
//     -Creates a new HTTP POST request to the integration's URL using the encoded data.
//
//     -Populates the request headers with the specific headers defined in the integration config.
//
//     -Sets a custom 'User-Agent' header using the build information (goxe/version).
//
//     -Executes the HTTP request using the client.
//
//   - If an error occurs during data transformation, request creation, or execution,
//     logs the error and continues with the next integration.
//
//   - Upon successful execution, prints the HTTP status response to the console.
//
//   - Ensures the response body is closed to prevent resource leaks.
func Integrations(logs map[string]map[string]*pipelines.LogStats, Shipper exporter.Shipper) {
	var client = &http.Client{
		Timeout: time.Second * 10,
	}
	for _, integration := range options.Config.Integrations {
		if integration.OnAggregation {
			data, err := Shipper.PrepareShip(logs)
			if err != nil {
				log.Println(err)
				continue
			}
			req, err := http.NewRequest("POST", integration.Url, bytes.NewBuffer(data))
			if err != nil {
				log.Println(err)
				continue
			}
			for header, value := range integration.Headers {
				req.Header.Set(header, value)
			}
			version, _ := debug.ReadBuildInfo()
			req.Header.Set("User-Agent", fmt.Sprintf("goxe/%s", version.Main.Version))
			res, err := client.Do(req)

			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Printf("\n[System] Status Response of %s: %s\n", integration.Url, res.Status)
			res.Body.Close()
		}
	}
}
