package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
	"strings"
	"sync"
	"syscall"
	"time"

	pkg "github.com/DumbNoxx/goxe/pkg/options"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

// getVersion returns the current version of the application.
//
// Returns:
//
//   - string: the application version. It prioritizes the value of the global 'version'
//     variable (typically set at link time). If empty, it attempts to retrieve the
//     version from build information (debug.ReadBuildInfo). If still unavailable,
//     it returns "vDev-build".
//
// The function performs:
//
//   - If the global 'version' variable is not empty, it returns it directly.
//   - Otherwise, it attempts to fetch build information using debug.ReadBuildInfo.
//   - If available and the main version is neither empty nor "(devel)", it returns it.
//   - In any other case, it returns "vDev-build" (indicating a development build).
func getVersion() string {
	if version != "" {
		return version
	}

	if info, ok := debug.ReadBuildInfo(); ok {
		if info.Main.Version != "" && info.Main.Version != "(devel)" {
			return info.Main.Version
		}
	}
	return "vDev-build"
}

// getVersionLatest queries the GitHub API to retrieve information about the latest release of the repository.
//
// Parameters:
//
//   - req: input parameter, currently unused (overwritten internally).
//   - res: input parameter, currently unused (overwritten internally).
//   - ctx: context for the HTTP request (allows for cancellation and timeouts).
//
// Returns:
//
//   - response: a pkg.ResponseGithubApi structure containing the latest version data.
//
// The function performs:
//
//   - Constructs a new HTTP GET request to "https://api.github.com/repos/DumbNoxx/goxe/releases/latest" using the provided context.
//   - If an error occurs during request creation, it logs it via log.Println and continues (returning a zero-value response).
//   - Executes the request using http.DefaultClient.Do.
//   - If an execution error occurs, it is logged.
//   - Ensures the response body is closed using a deferred call (defer).
//   - Reads the entire response body with io.ReadAll.
//   - If the HTTP status code is greater than 299, it logs an error message including the status code and the body content.
//   - If an error occurs while reading the body, it is logged.
//   - Attempts to deserialize (unmarshal) the JSON body into the response variable.
//   - If deserialization fails, it logs the error.
//   - Returns the structure (which may be partially populated if errors occurred).
func getVersionLatest(req *http.Request, res *http.Response, ctx context.Context) (response pkg.ResponseGithubApi) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.github.com/repos/DumbNoxx/goxe/releases/latest", nil)
	if err != nil {
		log.Println(err)
	}
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		log.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("failed to unmarshal github response:", err)
	}
	return response
}

// viewNewVersion periodically checks if a new version of the application is available on GitHub.
//
// Parameters:
//
//   - ctx: context for cancellation; when cancelled, the function terminates.
//   - wg: WaitGroup to notify the caller that the goroutine has finished (calls wg.Done() at the start).
//
// Returns:
//
//   - void: the function runs in an infinite loop until ctx is cancelled.
//
// The function performs:
//
//   - Retrieves the current version by calling getVersion().
//
//   - Creates a ticker that fires every 60 minutes.
//
//   - On every ticker tick:
//
//     -Calls getVersionLatest to fetch the latest release information from the GitHub API.
//
//     -If the retrieved version tag is "vDev-build" (development indicator), it ignores it and continues.
//
//     -If the tag matches the current version, it ignores it and continues.
//
//     -If a different version is found, it prints a message showing both the current and the new version,
//     followed by the release notes.
//
//   - If the context is cancelled, the function returns.
func viewNewVersion(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		res            *http.Response
		req            *http.Request
		currentVersion = getVersion()
	)

	ticker := time.NewTicker(60 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			release := getVersionLatest(req, res, ctx)

			if release.Tag_name == "vDev-build" {
				continue
			}

			if release.Tag_name == currentVersion {
				continue
			}

			fmt.Printf("Update available: %s -> %s\n", currentVersion, release.Tag_name)

			fmt.Println("--- Release Notes ---")
			fmt.Printf("\n%v\n", release.Body)
			fmt.Println("----------------------")
		case <-ctx.Done():
			return
		}
	}
}

// autoUpdate manages the automatic update process when a new version is detected.
//
// Parameters:
//
//   - ctx: main context (used to fetch the version from the API).
//   - cancel: function to cancel the main context (passed to executeHandoff).
//   - pipe: log input channel (passed to executeHandoff to be closed).
//   - wgProcessor: WaitGroup for processors (for handoff coordination).
//   - wgProducer: WaitGroup for producers (for handoff coordination).
//   - once: sync.Once to ensure the handoff logic is executed only once.
//
// Returns:
//
//   - void: the function attempts to replace the current binary. On success, the current process
//     is replaced by the new version via syscall.Exec. On failure, it logs the error and exits with os.Exit(1).
//
// The function performs:
//
//   - Retrieves the current executable path (os.Executable) and the user's home directory.
//
//   - Fetches the latest version from GitHub via getVersionLatest and compares it with getVersion.
//
//   - If the latest version matches the current one, it does nothing and waits for context cancellation.
//
//   - If a new version is available:
//
//     -If the current executable is outside GOPATH (local build), it attempts to recompile
//     the binary temporarily using "go build" and then renames it.
//
//     -If it is within GOPATH (installed via go install), it executes "go install ...@latest".
//
//     -Subsequently, calls executeHandoff to gracefully stop producers and processors.
//
//     -Finally, uses syscall.Exec to replace the current process with the new binary,
//     passing the "-is-upgrade" flag to indicate an active update.
//
//   - If the binary is located in /usr/bin/goxe (package manager install), it displays a message
//     recommending the use of the system's package manager and skips the auto-update.
//
//   - If any error occurs during compilation, installation, or handoff, it logs the error and exits.
func autoUpdate(ctx context.Context, cancel context.CancelFunc, pipe chan<- *pipelines.LogEntry, wgProcessor, wgProducer *sync.WaitGroup, once *sync.Once) {
	var (
		req *http.Request
		res *http.Response
	)
	currentLocation, err := os.Executable()
	home, _ := os.UserHomeDir()
	gopath := filepath.Join(home, "go")
	version := getVersionLatest(req, res, ctx)
	v := getVersion()
	if err != nil {
		log.Fatal(err)
	}
	if version.Tag_name != v {
		if !strings.HasPrefix(currentLocation, gopath) {
			fmt.Println("[Test] Local build detected, recompiling...")
			tempBin := currentLocation + ".tmp"
			cmd := exec.Command("go", "build", "-o", tempBin, "./cmd/goxe")
			output, err := cmd.CombinedOutput()
			if err != nil {
				log.Printf("Build failed: %v\n", err)
				log.Printf("Compiler says: %s\n", string(output))
				return
			}
			if err := os.Rename(tempBin, currentLocation); err != nil {
				fmt.Printf("[Error] Failed to swap binary: %v\n", err)
				return
			}

			fmt.Println("[System] Preparing handoff, flushing buffers...")
			executeHandoff(once, cancel, pipe, wgProcessor, wgProducer)
			err = syscall.Exec(currentLocation, []string{currentLocation, "-is-upgrade"}, os.Environ())

			fmt.Printf("\n[Error] the handoff failed!: %v\n", err)
			os.Exit(1)

			<-ctx.Done()
			return
		}
		if strings.HasPrefix(currentLocation, gopath) {
			cmd := exec.Command("go", "install", "github.com/DumbNoxx/goxe/cmd/goxe@latest")
			err := cmd.Run()
			if err != nil {
				log.Println(err)
				return
			}
		}
		fmt.Println("[System] Preparing handoff, flushing buffers...")

		executeHandoff(once, cancel, pipe, wgProcessor, wgProducer)
		err = syscall.Exec(currentLocation, []string{currentLocation, "-is-upgrade"}, os.Environ())

		fmt.Printf("\n[Error] ¡El salto a V2 falló!: %v\n", err)
		os.Exit(1)

		<-ctx.Done()
		return
	}

	if strings.HasPrefix(currentLocation, "/usr/bin/goxe") {
		fmt.Println("Goxe was installed via a package manager. Please use your package manager to update it to avoid versioning conflicts.")
	}
	<-ctx.Done()
}
