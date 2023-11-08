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

// memInfoCmd represents the memInfo command
var memInfoCmd = &cobra.Command{
	Use:   "memInfo",
	Short: "Gathers current memory usage statistics.",
	Long: `Calls upon the gopsutil package to retrieve current virtual memory
statistics, including total and available memory, providing a snapshot of
memory usage.`,
	Run: func(cmd *cobra.Command, args []string) {
		mememoryInfo, err := common.GetCurrentMemoryInfo()
		if err != nil {
			fmt.Println("Error:", err)
		}
		pp.Println(mememoryInfo)
	},
}

func init() {
	rootCmd.AddCommand(memInfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// memInfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// memInfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
