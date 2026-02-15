package cli

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/charmbracelet/log"
	"github.com/zs-health/zh-fhir-go/cmd/zh-fhir/internal/build"
	"github.com/zs-health/zh-fhir-go/cmd/zh-fhir/internal/config"
)

const (
	appName        = "zh-fhir"
	appDescription = "FHIR utility CLI for go-zh-fhir"
)

// CLI represents the root command structure.
type CLI struct {
	config.GlobalConfig
}

// Run executes the zh-fhir CLI with the provided build info.
func Run(version, commit, date string) error {
	// Set build info
	build.SetBuildInfo(version, commit, date)

	// Check for --version or -V flag before Kong parsing
	// This allows version flag to work without requiring a command
	for _, arg := range os.Args[1:] {
		if arg == "--version" || arg == "-V" {
			build.PrintBuildInfo()
			return nil
		}
	}

	// Parse command-line arguments
	cli := &CLI{}
	ctx := kong.Parse(cli,
		kong.Name(appName),
		kong.Description(appDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{
			"version": version,
			"commit":  commit,
			"date":    date,
		},
	)

	// If no command was selected, show help
	if ctx.Command() == "" {
		_ = ctx.PrintUsage(false)
		return fmt.Errorf("no command specified")
	}

	// Setup logging
	logger := setupLogger(&cli.GlobalConfig)

	// Log build info
	logger.Debug("zh-fhir CLI starting",
		"version", version,
		"commit", commit,
		"build_date", date,
	)

	// Run the selected command
	err := ctx.Run(&cli.GlobalConfig)
	if err != nil {
		logger.Error("command failed", "error", err)
		return err
	}

	return nil
}

// setupLogger configures the global logger based on config.
func setupLogger(cfg *config.GlobalConfig) *log.Logger {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    cfg.Debug,
		ReportTimestamp: true,
		TimeFormat:      "15:04:05",
	})

	// Set log level
	switch cfg.LogLevel {
	case "trace":
		logger.SetLevel(log.DebugLevel) // log package doesn't have trace, use debug
	case "debug":
		logger.SetLevel(log.DebugLevel)
	case "info":
		logger.SetLevel(log.InfoLevel)
	case "warn":
		logger.SetLevel(log.WarnLevel)
	case "error":
		logger.SetLevel(log.ErrorLevel)
	case "fatal":
		logger.SetLevel(log.FatalLevel)
	default:
		logger.SetLevel(log.InfoLevel)
	}

	// Set output format (pretty vs JSON)
	if !cfg.Pretty {
		logger.SetFormatter(log.JSONFormatter)
	}

	// Set as default logger
	log.SetDefault(logger)

	return logger
}

// ParseArgs is a convenience function for testing.
// It parses arguments and returns the CLI struct and Kong context.
func ParseArgs(args []string, version, commit, date string) (*CLI, *kong.Context, error) {
	build.SetBuildInfo(version, commit, date)

	cli := &CLI{}
	parser, err := kong.New(cli,
		kong.Name(appName),
		kong.Description(appDescription),
		kong.Vars{
			"version": version,
			"commit":  commit,
			"date":    date,
		},
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create parser: %w", err)
	}

	ctx, err := parser.Parse(args)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse arguments: %w", err)
	}

	return cli, ctx, nil
}
