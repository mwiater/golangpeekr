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

// diskUsageCmd represents the diskUsage command
var diskUsageCmd = &cobra.Command{
	Use:   "diskUsage",
	Short: "Obtains disk usage data for a given path.",
	Long: `This function uses the gopsutil package to find out the disk usage
statistics like free space and total space for a specified path, which is
crucial for disk space management.`,
	Run: func(cmd *cobra.Command, args []string) {
		diskUsage, err := common.GetCurrentDiskUsage("/")
		if err != nil {
			fmt.Println("Error:", err)
		}
		pp.Println(diskUsage)
	},
}

func init() {
	rootCmd.AddCommand(diskUsageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diskUsageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diskUsageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
