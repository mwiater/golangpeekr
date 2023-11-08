# golangpeekr

A repository for common, reusable functions

## Application

To see a list of available functions in the package, just: `go run .`

This will print out allfo the "public" functions with comments, arguments, and return types, e.g.:

```
Functions in the 'helpers' package:

[ slices ]

  // SliceContains checks if an item is present in the given slice.
  SliceContains(slice []T, item T) bool

  // SliceIntersection returns a new slice containing the common elements of two slices.
  SliceIntersection(slice1 []T, slice2 []T) []T

[ terminal ]

  // ClearTerminal clears the terminal screen based on the operating system.
  ClearTerminal() error

  // TerminalColor prints the given string to the terminal in the color corresponding to the error level
  TerminalColor(message string, level ErrorLevel)

  // TerminalInfo prints collects various information about the current terminal.
  TerminalInfo() (*Terminal, error)


Structs in the 'helpers' package:

[ terminal ]

  // Terminal struct holds information about the user's terminal environment.
  Height int
  Width int
  OutputType string
  NumberOfSupportedColors int
  TERM string
  SHELL string
  COLORTERM string
```

## CLI

### Linux
If your source machine is linux, use appropriate forward slashes and build natively:
Build: `go build -o bin/peekr`
Run:   `./bin/peekr --help`

If your source machine is Windows, use appropriate backward slashes and build natively:
Build: `env GOOS=windows GOARCH=amd64 go build -o bin\peekr`
Run:   N/A

### Windows
If your source machine is Windows, use appropriate backward slashes and build natively:
Build: `go build -o bin\peekr.exe`
Run:   `bin\peekr.exe --help`

If your source machine is linux, use appropriate forward slashes:
Build: `env GOOS=windows GOARCH=amd64 go build -o bin/peekr.exe`
Run:   N/A

```
Peek under the hood

Usage:
  peekr [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        TODO: -d directory, -s structOnly, -f functionOnly

Flags:
  -h, --help   help for peekr

Use "peekr [command] --help" for more information about a command.

```


## Cobra CLI

`env GOOS=windows GOARCH=amd64 go build -o bin\peekr.exe && bin\peekr.exe list`

`go build -o bin/peekr && ./bin/peekr list --help`