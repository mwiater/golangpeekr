/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/mwiater/peekr/peekr"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var FunctionsOnly bool
var StructsOnly bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "TODO: -d directory, -s structOnly, -f functionOnly",
	Long:  `TODO: -d directory, -s structOnly, -f functionOnly`,
	Run: func(cmd *cobra.Command, args []string) {

		dir := viper.GetString("directory")
		pkg := viper.GetString("package")

		// Call ListPackageFunctions if FunctionsOnly is true or if neither FunctionsOnly nor StructsOnly is true.
		if FunctionsOnly || (!FunctionsOnly && !StructsOnly) {
			peekr.ListPackageFunctions(dir, pkg)
		}

		// Call ListPackageStructs if StructsOnly is true or if neither FunctionsOnly nor StructsOnly is true.
		if StructsOnly || (!FunctionsOnly && !StructsOnly) {
			peekr.ListPackageStructs(dir, pkg)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Flags for List command
	listCmd.Flags().BoolVarP(&FunctionsOnly, "functions", "f", false, "Only list package functions.")
	viper.BindPFlag("functions", rootCmd.PersistentFlags().Lookup("functions"))

	listCmd.Flags().BoolVarP(&StructsOnly, "structs", "s", false, "Only list package structs.")
	viper.BindPFlag("structs", rootCmd.PersistentFlags().Lookup("structs"))
}
