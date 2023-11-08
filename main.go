package main

import (
	"os"

	"github.com/mwiater/peekr/cmd"
	"github.com/mwiater/peekr/config"
	"github.com/mwiater/peekr/helpers"
	"github.com/mwiater/peekr/peekr"
	"github.com/spf13/cobra"
)

var Logger = config.GetLogger()

func main() {
	Logger.Info("Info hello, world", "user", os.Getenv("USER"))
	Logger.Debug("Debug hello, world", "user", os.Getenv("USER"))
	Logger.Error("Error hello, world", "user", os.Getenv("USER"))
	Logger.Warn("Warn hello, world", "user", os.Getenv("USER"))

	// Check to see if we're executing the binary or calling via Cobra CLI
	var rootCmd = &cobra.Command{Use: "peekr"}
	cobraCommands := cmd.ListAllCobraCommands(rootCmd)
	intersection := helpers.SliceIntersection(os.Args, cobraCommands)
	if len(intersection) > 0 || helpers.SliceContains(os.Args, "--help") {
		helpers.ClearTerminal()
		cmd.Execute()
	} else {
		helpers.ClearTerminal()

		//peekr.ListPackageFunctions("./helpers")
		//peekr.ListPackageStructs("./helpers")

		peekr.ListPackageFunctions("/home/matt/projects/golangpeekr", "helpers")
		peekr.ListPackageStructs("/home/matt/projects/golangpeekr", "helpers")
	}
}
