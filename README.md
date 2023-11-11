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

```
The Peekr command by itself doesn't do anything at the moment. Please
see the Peekr list subcommand via: 'peekr list --help'

Usage:
  peekr [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        List the functions and structs within a package.

Flags:
  -d, --directory string   Absolute path of directory to scan.
  -h, --help               help for peekr
  -p, --package string     Name of package to scan.

Use "peekr [command] --help" for more information about a command.
```

Run:   `./bin/peekr list --help`

```
Peek into the source code for a high-level view of how a package
is constructed. By default, the 'list' command will print both
functions and structs in the specified package. You can filter
out one or the other by specifying '-s' and '-f' flags.

Usage:
  peekr list [flags]

Flags:
  -f, --functions   Only list package functions.
  -h, --help        help for list
  -s, --structs     Only list package structs.

Global Flags:
  -d, --directory string   Absolute path of directory to scan.
  -p, --package string     Name of package to scan.

```

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

`env GOOS=windows GOARCH=amd64 go build -o bin\peekr.exe && bin\peekr.exe list -d "/home/matt/projects/golangpeekr" -p "helpers"`

`go build -o bin/peekr && ./bin/peekr list --help`


### Build

`go build -o bin/peekr`

### Run

./bin/peekr list -d "/home/matt/projects/golangpeekr" -p "helpers"

```
Functions in the 'helpers' package:

File: /home/matt/projects/golangpeekr/helpers/slices.go

  // SliceContains checks if an item is present in the given slice.
  SliceContains(slice []T, item T) bool

  // SliceIntersection returns a new slice containing the common elements of two slices.
  SliceIntersection(slice1 []T, slice2 []T) []T


File: /home/matt/projects/golangpeekr/helpers/terminal.go

  // ClearTerminal clears the terminal screen based on the operating system.
  ClearTerminal() error

  // TerminalColor prints the given string to the terminal in the color corresponding to the error level
  TerminalColor(message string, level ErrorLevel)

  // TerminalInfo prints collects various information about the current terminal.
  TerminalInfo() (*Terminal, error)


Structs in the 'helpers' package:

File: /home/matt/projects/golangpeekr/helpers/terminal.go

  // Terminal struct holds information about the user's terminal environment.
  Height                     int
  Width                      int
  OutputType                 string
  NumberOfSupportedColors    int
  TERM                       string
  SHELL                      string
  COLORTERM                  string

```

./bin/peekr list -f -d "/home/matt/projects/golangpeekr" -p "helpers"

./bin/peekr list -s -d "/home/matt/projects/golangpeekr" -p "helpers"

./bin/peekr list -sf -d "/home/matt/projects/golangpeekr" -p "helpers"

## Tests

`go install gotest.tools/gotestsum@latest`

`go get gotest.tools/gotestsum`

`gotestsum --format=testname`