package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zs-health/zh-fhir-go/cmd/zh-fhir/internal/cli"
)

// Version information (injected at build time via ldflags)
var (
	Version = "develop"
	Commit  = ""
	Date    = ""
)

func main() {
	termServer := flag.Bool("term-server", false, "Start the terminology server")
	port := flag.Int("port", 8080, "Port for the terminology server")
	flag.Parse()

	if *termServer {
		StartTerminologyServer(*port)
		return
	}

	if err := cli.Run(Version, Commit, Date); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
