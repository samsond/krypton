package main

import (
    "os"
    "github.com/spf13/cobra"
)

func main() {
    var rootCmd = &cobra.Command{
        Use:   "kptn",
        Short: "kptn is a CLI for managing Kubernetes resources using the Krypton DSL",
    }

    // Add the subcommands to the root command
    rootCmd.AddCommand(versionCmd)
    

    // Execute the root command
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
