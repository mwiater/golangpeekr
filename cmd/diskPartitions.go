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

// diskPartitionsCmd represents the diskPartitions command
var diskPartitionsCmd = &cobra.Command{
	Use:   "diskPartitions",
	Short: "Lists all disk partitions.",
	Long: `Retrieves a list of disk partitions, optionally including all mount
points, using the gopsutil package. This information is key for disk
management and partition analysis.`,
	Run: func(cmd *cobra.Command, args []string) {
		diskPartitions, err := common.GetDiskPartitions(true)
		if err != nil {
			fmt.Println("Error:", err)
		}
		pp.Println(diskPartitions)
	},
}

func init() {
	rootCmd.AddCommand(diskPartitionsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diskPartitionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diskPartitionsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
