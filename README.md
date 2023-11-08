# golangcommon

A repository for common, reusable functions

## Application

To see a list of available functions in the package, just: `go run .`

This will print out allfo the "public" functions with comments, arguments, and return types, e.g.:

```
Functions in the 'common' package:

// ProgressBar creates and returns a new progress bar with the specified size and custom options.
// It enables color codes, displays the progress in bytes, shows the count,
// predicts the remaining time, sets a description, and defines a custom theme for the progress bar.
ProgressBar(size int) *progressbar.ProgressBar


// GetCurrentCPUInfo provides current CPU information using the gopsutil package.
GetCurrentCPUInfo() ([]cpu.InfoStat, error)

// GetCurrentDiskUsage provides current disk usage information.
GetCurrentDiskUsage(path string) (*disk.UsageStat, error)

// GetCurrentMemoryInfo provides current memory information using the gopsutil package.
GetCurrentMemoryInfo() (*mem.VirtualMemoryStat, error)

// GetDiskPartitions lists all the disk partitions.
GetDiskPartitions(all bool) ([]disk.PartitionStat, error)

// GetHostInfo provides detailed host information.
GetHostInfo() (*host.InfoStat, error)

// GetInternalIPv4 returns the first internal IPv4 address it finds,
// typically one that starts with "192.168". If no such address is found,
// it returns an error.
GetInternalIPv4() (string, error)

// GetNetworkIOCounters provides network I/O counters.
GetNetworkIOCounters(pernic bool) ([]gopsutilNet.IOCountersStat, error)

// GetNetworkInterfaces lists all the network interfaces.
GetNetworkInterfaces() ([]gopsutilNet.InterfaceStat, error)

// GetSystemLoadAverage provides the system load average.
GetSystemLoadAverage() (*load.AvgStat, error)


// ClearTerminal clears the terminal screen based on the operating system.
ClearTerminal() error

// TerminalColor prints the given string to the terminal in the color corresponding to the error level
TerminalColor(message string, level ErrorLevel)

// TerminalInfo prints collects various information about the current terminal.
TerminalInfo() (*Terminal, error)
```

## CLI

go build -o bin/golangcommon && ./bin/golangcommon files --help

env GOOS=windows GOARCH=amd64 go build -o bin\golangcommon.exe && bin\golangcommon.exe --help

```
A utitlity belt of common functions I fing useful during golang development.

Usage:
  golangcommon [command]

Available Commands:
  completion        Generate the autocompletion script for the specified shell
  cpuInfo           Retrieves detailed CPU information.
  cpuUsage          Fetches current CPU utilization statistics.
  diskPartitions    Lists all disk partitions.
  diskUsage         Obtains disk usage data for a given path.
  help              Help about any command
  hostInfo          Delivers comprehensive host system information.
  list              List all the functions available within the 'common' package.
  localIP           Finds an internal IPv4 address.
  memInfo           Gathers current memory usage statistics.
  networkInterfaces Lists all network interfaces on the host.
  systemLoadAverage Provides the system's load average.
  terminalInfo      Provides details about the current terminal or shell environment.

Flags:
  -h, --help   help for golangcommon

Use "golangcommon [command] --help" for more information about a command.
```

## Demo

Runs through the current common Functions: `go run demo\demo.go`

## Cobra CLI

`go build -o bin/golangcommon && ./bin/golangcommon --help`

## Godoc

`go install golang.org/x/tools/cmd/godoc@latest`

`godoc -play -http={ip-address}:6060`, e.g.: `godoc -play -http=192.168.221.244:6060`

## Testing

This is pretty sparse at the moment, but...

`go install gotest.tools/gotestsum@latest`

`go clean -testcache && gotestsum --format testname`
