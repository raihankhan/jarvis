package main

import (
	"github.com/raihankhan/jarvis/cmd"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "jarvis",
		Short: "Jarvis is a utility tool for workstation activities.",
	}

	rootCmd.AddCommand(cmd.NewAppUpdateCommand())

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
