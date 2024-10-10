package main

import (
	"os"

	"github.com/samsond/krypton/cmd/krypton/completion"
	"github.com/samsond/krypton/cmd/krypton/generate"
	ver "github.com/samsond/krypton/cmd/krypton/version"
	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{
		Use:   "kptn",
		Short: "kptn is a CLI for managing Kubernetes resources using the Krypton DSL",
	}

	// Add the version command using the alias
	rootCmd.AddCommand(ver.NewVersionCommand())
	rootCmd.AddCommand(generate.NewGenerateCommand())
	rootCmd.AddCommand(completion.NewCompletionCommand(rootCmd))

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
