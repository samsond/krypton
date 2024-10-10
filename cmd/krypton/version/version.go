package version

import (
	_ "embed"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

//go:embed VERSION
var rawVersion string

var stage = "dev"

// AppVersion holds the main version number
var appVersion string

// Prerelease holds pre-release information like "-beta"
var prerelease string

// Initialize version information
func init() {
	setVersionAndStage(rawVersion, stage)
}

// setVersionAndStage sets the version and stage variables
func setVersionAndStage(version, buildStage string) {
	// Check if an environment variable for the stage is set
	envStage := os.Getenv("KPTN_STAGE")
	if envStage != "" {
		buildStage = envStage
	}

	appVersion = strings.TrimSpace(version)
	stage = buildStage

	// Append "-beta" for "dev" stage
	if stage == "dev" {
		prerelease = "-beta"
	} else {
		prerelease = "" // No suffix for prod release builds.
	}
}

// versionString returns the complete version string
func versionString() string {
	return fmt.Sprintf("%s%s", appVersion, prerelease)
}

// NewVersionCommand returns a new cobra command for the version
func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of kptn",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "kptn version: %s\n", versionString())
		},
	}
}
