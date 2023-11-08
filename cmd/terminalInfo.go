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

// terminalInfoCmd represents the terminalInfo command
var terminalInfoCmd = &cobra.Command{
	Use:   "terminalInfo",
	Short: "Provides details about the current terminal or shell environment.",
	Long:  `Provides details about the current terminal or shell environment.`,
	Run: func(cmd *cobra.Command, args []string) {
		terminalInfo, err := common.TerminalInfo()
		if err != nil {
			fmt.Println("Error:", err)
		}
		pp.Println(terminalInfo)
	},
}

func init() {
	rootCmd.AddCommand(terminalInfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// terminalInfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// terminalInfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
