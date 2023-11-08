/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/mwiater/peekr/peekr"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "TODO: -d directory, -s structOnly, -f functionOnly",
	Long:  `TODO: -d directory, -s structOnly, -f functionOnly`,
	Run: func(cmd *cobra.Command, args []string) {
		peekr.ListPackageFunctions("/home/matt/projects/golangpeekr", "helpers")
		peekr.ListPackageStructs("/home/matt/projects/golangpeekr", "helpers")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	listCmd.Flags().BoolP("functionsOnly", "f", false, "Only list package functions.")
	listCmd.Flags().BoolP("structsOnly", "s", false, "Only list package structs.")
}
