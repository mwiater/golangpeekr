/*
Copyright © 2023 Matt J. Wiater
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "peekr",
	Short: "Peek under the hood",
	Long:  `Peek under the hood`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.golangcommon.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// ListAllCobraCommands prints all commands and subcommands recursively
func ListAllCobraCommands(cmd *cobra.Command) []string {
	var commands []string
	// Get a list of child commands
	subcommands := rootCmd.Commands()
	// Recursively print each subcommand
	for _, subcmd := range subcommands {
		commands = append(commands, subcmd.Use)
	}
	return commands
}
