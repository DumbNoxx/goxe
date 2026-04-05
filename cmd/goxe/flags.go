package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sync"

	"github.com/DumbNoxx/goxe/internal/factory"
	"github.com/DumbNoxx/goxe/internal/options"
)

func init() {
	flag.BoolVar(&versionFlag, "v", false, "Show program version")
	flag.BoolVar(&versionFlag, "version", false, "Show program version")
	isUpgrade = flag.Bool("is-upgrade", false, "Internal use for hot-swap")
	flag.BoolVar(&flagRouteFile, "brew", false, "Distill and normalize a raw log file into a structured report")
	flag.BoolVar(&flagRouteFile, "b", false, "Distill and normalize a raw log file into a structured report")
}

var (
	versionFlag   bool
	isUpgrade     *bool
	flagRouteFile bool
	routeFile     string
	version       string
)

func updateArg() {
	fmt.Println("[Goxe] Sending update signal to the active instance...")
	cmd := exec.Command("pkill", "-SIGUSR1", "goxe")
	cmd.Run()
}

func brewFlag(wg *sync.Mutex, getConfig options.ConfigProvider) error {
	routeFile = flag.Arg(0)
	idLog := flag.Arg(1)
	file, err := os.Open(routeFile)
	if os.IsNotExist(err) {
		return err
	}
	defer file.Close()
	manager, err := factory.GetIngestor(routeFile)
	if err != nil {
		return err
	}

	manager.FileNormalized(file, idLog, wg, routeFile, Shipper, getConfig)

	return nil
}
