package main

import (
    "fmt"
    "github.com/spf13/cobra"
)

var version = "0.1.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Print the version number of kptn",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("kptn version: %s\n", version)
    },
}
