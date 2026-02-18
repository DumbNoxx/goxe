package main

import (
	"flag"
	"fmt"
	"os/exec"
)

var (
	versionFlag bool
	isUpgrade   *bool
	routeFile   string
	version     string
)

func init() {
	flag.BoolVar(&versionFlag, "v", false, "show program version")
	flag.BoolVar(&versionFlag, "version", false, "show program version")
	isUpgrade = flag.Bool("is-upgrade", false, "Internal use for hot-swap")
}

func updateArg() {
	fmt.Println("Sending update signal to the active instance...")
	cmd := exec.Command("pkill", "-SIGUSR1", "goxe")
	cmd.Run()
}

func brewFlag() {
	routeFile = flag.Arg(0)
	fmt.Println(routeFile)
}
