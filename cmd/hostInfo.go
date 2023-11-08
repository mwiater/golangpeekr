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

// hostInfoCmd represents the hostInfo command
var hostInfoCmd = &cobra.Command{
	Use:   "hostInfo",
	Short: "Delivers comprehensive host system information.",
	Long: ` Fetches detailed information about the host system, such as uptime,
boot time, and OS specifics, using the gopsutil package. It's a vital function
for system diagnostics and inventory.`,
	Run: func(cmd *cobra.Command, args []string) {
		hostInfo, err := common.GetHostInfo()
		if err != nil {
			fmt.Println("Error:", err)
		}
		pp.Println(hostInfo)
	},
}

func init() {
	rootCmd.AddCommand(hostInfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hostInfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hostInfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
