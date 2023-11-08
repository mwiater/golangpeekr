package main

import (
	"fmt"
	"time"

	"github.com/k0kubun/pp"
	"github.com/mwiater/golangcommon/common"
)

func main() {
	common.ClearTerminal()
	fmt.Println("ClearTerminal(): Success!")
	fmt.Println()

	fmt.Println("Running ProgressBar() examples:")
	progressBarDataSize := 100
	bar := common.ProgressBar(progressBarDataSize)
	for i := 0; i <= progressBarDataSize; i++ {
		bar.Add(1)
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("-------------------------------------------")
	fmt.Println()

	fmt.Println("Running TerminalColor() examples:")
	common.TerminalColor("This is an alert message", common.Alert)
	common.TerminalColor("This is a critical message", common.Critical)
	common.TerminalColor("This is an error message", common.Error)
	common.TerminalColor("This is a warning message", common.Warn)
	common.TerminalColor("This is a notice message", common.Notice)
	common.TerminalColor("This is an info message", common.Info)
	common.TerminalColor("This is a debug message", common.Debug)
	fmt.Println("-------------------------------------------")
	fmt.Println()

	fmt.Println("Running GetInternalIPv4() example:")
	localIP, err := common.GetInternalIPv4()
	if err != nil {
		error := fmt.Sprintf("Error: %s", err)
		common.TerminalColor(error, common.Error)
	}
	common.TerminalColor("Result: "+localIP, common.Debug)
	pp.Println()
	fmt.Println("-------------------------------------------")
	fmt.Println()

	fmt.Println("Running TerminalInfo() example:")
	termInfo, err := common.TerminalInfo()
	if err != nil {
		error := fmt.Sprintf("Error: %s", err)
		common.TerminalColor(error, common.Error)
	}
	pp.Println("Result:", termInfo)
	fmt.Println("-------------------------------------------")
	fmt.Println()

	fmt.Println("Running GetCurrentCPUInfo() example:")
	cpu, _ := common.GetCurrentCPUInfo()
	pp.Println("Result", cpu)
	fmt.Println("-------------------------------------------")
	fmt.Println()

	fmt.Println("Running GetCurrentDiskUsage() example:")
	disk, _ := common.GetCurrentDiskUsage("/")
	pp.Println("Result", disk)
	fmt.Println("-------------------------------------------")
	fmt.Println()

	fmt.Println("Running GetCurrentMemoryInfo() example:")
	mem, _ := common.GetCurrentMemoryInfo()
	pp.Println("Result", mem)
	fmt.Println("-------------------------------------------")
	fmt.Println()

	fmt.Println("Running GetDiskPartitions() example:")
	part, _ := common.GetDiskPartitions(true)
	pp.Println("Result", part)
	fmt.Println("-------------------------------------------")
	fmt.Println()

	fmt.Println("Running GetHostInfo() example:")
	host, _ := common.GetHostInfo()
	pp.Println("Result", host)
	fmt.Println("-------------------------------------------")
	fmt.Println()

	fmt.Println("Running GetNetworkIOCounters() example:")
	netIO, _ := common.GetNetworkIOCounters(true)
	pp.Println("Result", netIO)
	fmt.Println("-------------------------------------------")
	fmt.Println()

	fmt.Println("Running GetNetworkInterfaces() example:")
	net, _ := common.GetNetworkInterfaces()
	pp.Println("Result", net)
	fmt.Println("-------------------------------------------")
	fmt.Println()

	fmt.Println("Running GetSystemLoadAverage() example:")
	load, _ := common.GetSystemLoadAverage()
	pp.Println("Result", load)
	fmt.Println("-------------------------------------------")
	fmt.Println()

}
