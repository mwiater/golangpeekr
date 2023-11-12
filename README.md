# golangpeekr

A repository for common, reusable functions

## Application

To see a list of available functions in the package as a demo, just: `go run .`

This will print out all of the exported functions with comments, arguments, and return types, e.g.:

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

### Windows
If your source machine is Windows, use appropriate backward slashes and build natively:
Build: `go build -o bin\peekr.exe`
Run:   `bin\peekr.exe --help`

### CLI Options / flags

Show both functions and structs:

* `./bin/peekr list -d "/home/matt/projects/golangpeekr" -p "helpers"`
* `./bin/peekr list -sf -d "/home/matt/projects/golangpeekr" -p "helpers"`

Show functions only:

* `./bin/peekr list -f -d "/home/matt/projects/golangpeekr" -p "helpers"`

Show structs only:
* `./bin/peekr list -s -d "/home/matt/projects/golangpeekr" -p "helpers"`

## Tests

`go install gotest.tools/gotestsum@latest`

`gotestsum --format=testname`

```

EMPTY .
EMPTY cmd
EMPTY config
EMPTY helpers
PASS peekr.TestExtractFuncParams/No_params (0.00s)
PASS peekr.TestExtractFuncParams/One_param (0.00s)
PASS peekr.TestExtractFuncParams/Multiple_params (0.00s)
PASS peekr.TestExtractFuncParams/Unnamed_params (0.00s)
PASS peekr.TestExtractFuncParams (0.00s)
PASS peekr.TestExtractFuncResults/No_results (0.00s)
PASS peekr.TestExtractFuncResults/One_result (0.00s)
PASS peekr.TestExtractFuncResults/Multiple_results (0.00s)
PASS peekr.TestExtractFuncResults/Named_results (0.00s)
PASS peekr.TestExtractFuncResults (0.00s)
PASS peekr.TestExprToString/Identifier (0.00s)
PASS peekr.TestExprToString (0.00s)
PASS peekr.TestCommentify/Single_line_comment (0.00s)
PASS peekr.TestCommentify/Multi-line_comment (0.00s)
PASS peekr.TestCommentify/Empty_comment (0.00s)
PASS peekr.TestCommentify (0.00s)
PASS peekr (cached)

DONE 16 tests in 0.167s
```