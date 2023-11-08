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

// cpuUsageCmd represents the cpuUsage command
var cpuUsageCmd = &cobra.Command{
	Use:   "cpuUsage",
	Short: "Fetches current CPU utilization statistics.",
	Long: `Utilizes the gopsutil package to gather timing statistics of
the CPU, indicating the time spent in various states. It's essential for
performance monitoring and analysis.`,
	Run: func(cmd *cobra.Command, args []string) {
		cpuUsage, err := common.GetCurrentCPUUsage()
		if err != nil {
			fmt.Println("Error:", err)
		}
		pp.Println(cpuUsage)
	},
}

func init() {
	rootCmd.AddCommand(cpuUsageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cpuUsageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cpuUsageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
