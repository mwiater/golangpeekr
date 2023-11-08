/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/mwiater/peekr/peekr"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "TODO: -d directory, -s structOnly, -f functionOnly",
	Long:  `TODO: -d directory, -s structOnly, -f functionOnly`,
	Run: func(cmd *cobra.Command, args []string) {
		peekr.ListPackageFunctions(viper.GetString("directory"), viper.GetString("package"))
		peekr.ListPackageStructs(viper.GetString("directory"), viper.GetString("package"))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Flags for List command
	listCmd.Flags().BoolP("functionsOnly", "f", false, "Only list package functions.")
	listCmd.Flags().BoolP("structsOnly", "s", false, "Only list package structs.")
}
