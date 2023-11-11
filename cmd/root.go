/*
Copyright © 2023 Matt J. Wiater
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Directory string
var Package string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "peekr",
	Short: "Peek under the hood",
	Long:  `The Peekr command by itself doesn't do anything at the moment. Please
see the Peekr list subcommand via: 'peekr list --help'`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		// TO DO
		os.Exit(1)
	}
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringVarP(&Directory, "directory", "d", "", "Absolute path of directory to scan.")
	rootCmd.MarkPersistentFlagRequired("directory")
	viper.BindPFlag("directory", rootCmd.PersistentFlags().Lookup("directory"))

	rootCmd.PersistentFlags().StringVarP(&Package, "package", "p", "", "Name of package to scan.")
	rootCmd.MarkPersistentFlagRequired("package")
	viper.BindPFlag("package", rootCmd.PersistentFlags().Lookup("package"))
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
