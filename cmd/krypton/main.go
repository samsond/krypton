package main

import (
    "os"
    "github.com/spf13/cobra"
    ver "github.com/samsond/krypton/cmd/krypton/version"
)

func main() {
    var rootCmd = &cobra.Command{
        Use:   "kptn",
        Short: "kptn is a CLI for managing Kubernetes resources using the Krypton DSL",
    }

    // Add the version command using the alias
    rootCmd.AddCommand(ver.NewVersionCommand())
    

    // Execute the root command
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
