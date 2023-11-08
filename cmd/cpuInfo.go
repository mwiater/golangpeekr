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

// cpuInfoCmd represents the cpuInfo command
var cpuInfoCmd = &cobra.Command{
	Use:   "cpuInfo",
	Short: "Retrieves detailed CPU information.",
	Long: `This function calls the gopsutil package to obtain comprehensive
details about the CPU, such as model, cores, and speed. It's useful for
understanding the CPU's capabilities and current specifications.`,
	Run: func(cmd *cobra.Command, args []string) {
		cpuInfo, err := common.GetCurrentCPUInfo()
		if err != nil {
			fmt.Println("Error:", err)
		}
		pp.Println(cpuInfo)
	},
}

func init() {
	rootCmd.AddCommand(cpuInfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cpuInfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cpuInfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
