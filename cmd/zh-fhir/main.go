package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/zs-health/zh-fhir-go/cmd/zh-fhir/internal/cli"
	"github.com/zs-health/zh-fhir-go/internal/ig"
	"github.com/zs-health/zh-fhir-go/internal/server"
)

// Version information (injected at build time via ldflags)
var (
	Version = "develop"
	Commit  = ""
	Date    = ""
)

func main() {
	serverMode := flag.Bool("server", false, "Start the full FHIR server")
	termServer := flag.Bool("term-server", false, "Start the legacy terminology server")
	port := flag.Int("port", 8080, "Port for the server")
	igPath := flag.String("ig", "./BD-Core-FHIR-IG", "Path to the Bangladesh FHIR IG")
	flag.Parse()

	if *serverMode {
		loader := ig.NewLoader()
		log.Printf("Loading IG data from %s...", *igPath)
		if err := loader.LoadFromIG(*igPath); err != nil {
			log.Printf("Warning: Failed to load IG: %v", err)
		}
		log.Printf("Loaded %d CodeSystems and %d ValueSets", len(loader.CodeSystems), len(loader.ValueSets))

		s := server.NewServer(loader)
		s.Start(*port)
		return
	}

	if *termServer {
		StartTerminologyServer(*port)
		return
	}

	if err := cli.Run(Version, Commit, Date); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
