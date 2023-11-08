package main

import (
	"os"

	"github.com/mwiater/golangcommon/cmd"
	"github.com/mwiater/golangcommon/common"
	"github.com/spf13/cobra"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Check to see if we're executing the binary or calling via Cobra CLI
	var rootCmd = &cobra.Command{Use: "golangcommon"}
	cobraCommands := cmd.ListAllCobraCommands(rootCmd)
	intersection := common.SliceIntersection(os.Args, cobraCommands)
	if len(intersection) > 0 || common.SliceContains(os.Args, "--help") {
		cmd.Execute()
	} else {
		common.ClearTerminal()
		common.ListPackageFunctions("./common/")
	}
}
