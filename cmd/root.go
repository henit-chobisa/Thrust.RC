/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "thrust",
	Short: "Thrust helps you setup a test environment for your rocket.chat apps blazingly fast, where you can use and test your Rocket.chat apps",
	Long:  `Stop configuring your workspace everytime you wanna test Rocket.Chat Apps and waste the initial 20 mins of yours, let thrust handle it, just type thrust start <your_app_directory>. On launch it sets up everything which you would need and launches an RC Server and installs the app in that for you to test.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initializeConfig() {
	viper.SetConfigName(".thrust")
	viper.SetConfigType("yaml")
}

func init() {
	cobra.OnInitialize(initializeConfig)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
