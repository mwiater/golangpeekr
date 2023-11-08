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

// localIPCmd represents the localIP command
var localIPCmd = &cobra.Command{
	Use:   "localIP",
	Short: "Finds an internal IPv4 address.",
	Long: `Searches for and returns the first internal IPv4 address, commonly
within the "192.168" subnet. If none is found, it returns an error. This is
useful for services that need to bind to an internal network interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		localIP, err := common.GetInternalIPv4()
		if err != nil {
			fmt.Println("Error:", err)
		}
		pp.Println(localIP)
	},
}

func init() {
	rootCmd.AddCommand(localIPCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// localIPCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// localIPCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
