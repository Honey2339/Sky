package main

import (
	"SkyRP/config"
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use: "init",
	Short: "Initialize the config",
	Run: func(cmd *cobra.Command, args []string) {
		config.Create_config()
	},
}

var editConfig = &cobra.Command{
	Use: "edit",
	Short: "Edit the config",
	Run: func(cmd *cobra.Command, args []string) {
		config.Add_config()
	},
}

func main() {
	rootCmd := &cobra.Command{
		Use: "sky",
		Short: "Sky is a CLI tool",
	}

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(editConfig)

	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}