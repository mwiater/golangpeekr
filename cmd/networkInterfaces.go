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

// networkInterfacesCmd represents the networkInterfaces command
var networkInterfacesCmd = &cobra.Command{
	Use:   "networkInterfaces",
	Short: "Lists all network interfaces on the host.",
	Long: `Gathers information on each network interface available on the system,
which is important for network configuration and troubleshooting.`,
	Run: func(cmd *cobra.Command, args []string) {
		networkInterfaces, err := common.GetNetworkInterfaces()
		if err != nil {
			fmt.Println("Error:", err)
		}
		pp.Println(networkInterfaces)
	},
}

func init() {
	rootCmd.AddCommand(networkInterfacesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// networkInterfacesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// networkInterfacesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
