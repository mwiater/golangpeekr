/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/mwiater/golangcommon/common"
	"github.com/spf13/cobra"
)

// systemLoadAverageCmd represents the systemLoadAverage command
var systemLoadAverageCmd = &cobra.Command{
	Use:   "systemLoadAverage",
	Short: "Provides the system's load average.",
	Long: `Acquires the system's load average over the last 1, 5, and 15 minutes,
offering insights into system performance under load.`,
	Run: func(cmd *cobra.Command, args []string) {
		systemLoadAverage, err := common.GetSystemLoadAverage()
		if err != nil {
			fmt.Println("Error:", err)
		}
		pp.Println(systemLoadAverage)
	},
}

func init() {
	rootCmd.AddCommand(systemLoadAverageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// systemLoadAverageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// systemLoadAverageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
