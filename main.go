package main

import (
	"SkyRP/config"
	"SkyRP/server"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use: "start",
	Short: "Starts your reverse proxy",
	Run: func(cmd *cobra.Command, args []string) {
		server.HttpServer()
	},
}

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

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Shows your current config",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := config.Get_config_data()
		if err != nil {
			log.Fatalf("Error fetching config: %v", err)
			return
		}

		titleStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205"))

		labelStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("99")).
			Bold(true)

		valueStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("15")).
			PaddingLeft(2)

		fmt.Println(titleStyle.Render("Configuration Info"))

		for _, upstream := range data.Server.Upstreams {
			fmt.Println(labelStyle.Render("Worker:"), valueStyle.Render(upstream.ID))
			fmt.Println(labelStyle.Render("URL:   "), valueStyle.Render(upstream.URL))
			fmt.Println()
		}
	},
}

func main() {
	rootCmd := &cobra.Command{
		Use: "sky",
		Short: "🔧 Sky is a CLI tool",
	}

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(editConfig)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(infoCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}