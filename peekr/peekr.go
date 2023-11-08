package peekr

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/mwiater/peekr/helpers"
)

// FunctionInfo holds metadata about a function within a Go source file.
// It includes the file name, function signature, associated comments,
// parameter list, and return types.
type FunctionInfo struct {
	FileName string // Name of the file where the function is defined
	Function string // Name of the function
	Comments string // Comments associated with the function
	Params   string // List of parameters in the function signature
	Returns  string // Return types of the function
}

// StructInfo holds metadata about a struct type within a Go source file.
// It includes the struct name, slice of its fields, and associated comments.
type StructInfo struct {
	Name    string      // Name of the struct
	Fields  []FieldInfo // Slice of FieldInfo representing each field in the struct
	Comment string      // Comments associated with the struct
}

// FieldInfo holds metadata about a field within a struct.
// It includes the field name, field type, and associated comments.
type FieldInfo struct {
	Name    string // Name of the field
	Type    string // Type of the field
	Comment string // Comments associated with the field
}

// ListPackageFunctions prints a color-coded list of functions from the specified package.
// It retrieves function metadata using PackageFunctions and formats the output.
func ListPackageFunctions(dir, pkgName string) {
	// Retrieve function information from the specified package.
	groupedFunctions, err := PackageFunctions(dir, pkgName)
	if err != nil {
		panic(err)
	}

	// Prepare a sorted list of group names (file names) for output.
	var groupNames []string
	for groupName := range groupedFunctions {
		groupNames = append(groupNames, groupName)
	}
	sort.Strings(groupNames) // Sort the group names alphabetically.

	// Print the header for the function list output.
	header := fmt.Sprintf("\nFunctions in the %s package:", fmt.Sprintf("'%s'", pkgName))
	helpers.TerminalColor(header, helpers.Notice)

	// Iterate over each group and print the functions contained within.
	for _, groupName := range groupNames {
		functions := groupedFunctions[groupName]
		helpers.TerminalColor("\nFile: "+groupName+"\n", helpers.Info)
		for _, funcInfo := range functions {
			// Print function comments and signatures with color coding.
			helpers.TerminalColor("  "+strings.TrimSpace(funcInfo.Comments), helpers.Info)
			signature := fmt.Sprintf("  %s(%s) %s", funcInfo.Function, funcInfo.Params, funcInfo.Returns)
			helpers.TerminalColor(signature, helpers.Debug)
			fmt.Println()
		}
	}
}

// ListPackageStructs prints a color-coded list of structs from the specified package.
// It retrieves struct metadata using PackageStructs and formats the output.
func ListPackageStructs(dir, pkgName string) error {
	// Retrieve struct information from the specified package.
	structsMap, err := PackageStructs(dir, pkgName)
	if err != nil {
		return err
	}

	// Sort the group names (file names) for output.
	var groupNames []string
	for groupName := range structsMap {
		groupNames = append(groupNames, groupName)
	}
	sort.Strings(groupNames) // Sort the group names alphabetically.

	// Print the header for the struct list output.
	header := fmt.Sprintf("\nStructs in the %s package:", fmt.Sprintf("'%s'", pkgName))
	helpers.TerminalColor(header, helpers.Notice)

	// Iterate over each group and print the structs contained within.
	for _, groupName := range groupNames {
		structs := structsMap[groupName]
		if len(structs) > 0 {
			helpers.TerminalColor("\nFile: "+groupName+"\n", helpers.Info)
			for _, structInfo := range structs {
				// Print struct comments and definitions with color coding.
				if structInfo.Comment != "" {
					comment := fmt.Sprintf("  // %s", strings.TrimSpace(structInfo.Comment))
					helpers.TerminalColor(comment, helpers.Info)
				}
				structHeader := fmt.Sprintf("type %s struct {", structInfo.Name)
				helpers.TerminalColor("  "+structHeader, helpers.Debug)
				for _, field := range structInfo.Fields {
					fieldStr := fmt.Sprintf("    %s %s", field.Name, field.Type)
					helpers.TerminalColor(fieldStr, helpers.Debug)
					if field.Comment != "" {
						fieldComment := fmt.Sprintf("    // %s", strings.TrimSpace(field.Comment))
						helpers.TerminalColor(fieldComment, helpers.Info)
					}
				}
				helpers.TerminalColor("  }\n", helpers.Debug)
			}
		}
	}

	return nil
}

// PackageFunctions retrieves a map of FunctionInfo slices indexed by file path.
// It scans the specified package directory for Go files and extracts metadata
// for each exported function.
func PackageFunctions(dir, pkgName string) (map[string][]FunctionInfo, error) {
	fset := token.NewFileSet()                 // Create a new file set for parsing.
	funcMap := make(map[string][]FunctionInfo) // Initialize a map to store function information.

	// Walk the directory tree recursively to find Go files.
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and test files.
		if info.IsDir() || strings.HasSuffix(info.Name(), "_test.go") {
			return nil
		}

		// Parse only Go files that match the specified package name.
		if strings.HasSuffix(info.Name(), ".go") {
			// Parse the Go source file
			f, parseErr := parser.ParseFile(fset, path, nil, parser.ParseComments)
			if parseErr != nil {
				return parseErr
			}

			// Check if the file's package name matches the desired package.
			if f.Name.Name != pkgName {
				return nil
			}

			// Use the file name without the extension for grouping
			groupName := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))

			// Process declarations within the file.
			for _, decl := range f.Decls {
				if fn, ok := decl.(*ast.FuncDecl); ok && fn.Name.IsExported() {
					// Extract comments, parameters, and return types for exported functions.
					var comments string
					if fn.Doc != nil {
						comments = Commentify(fn.Doc.Text())
					}

					// Extract the function's parameters and return types
					params := ExtractFuncParams(fn.Type.Params)
					returns := ExtractFuncResults(fn.Type.Results)

					// Store the function information in the map.
					funcInfo := FunctionInfo{
						FileName: groupName,
						Function: fn.Name.Name,
						Comments: comments,
						Params:   params,
						Returns:  returns,
					}
					funcMap[path] = append(funcMap[path], funcInfo)
				}
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Sort the functions within each group alphabetically
	for _, functions := range funcMap {
		sort.Slice(functions, func(i, j int) bool {
			return functions[i].Function < functions[j].Function
		})
	}

	return funcMap, nil
}

// PackageStructs retrieves a map of StructInfo slices indexed by file path.
// It scans the specified package directory for Go files and extracts metadata
// for each exported struct.
func PackageStructs(dir, pkgName string) (map[string][]StructInfo, error) {
	fset := token.NewFileSet()                  // Create a new file set for parsing.
	structsMap := make(map[string][]StructInfo) // Initialize a map to store struct information.

	// Walk through the directory tree to find Go source files.
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Ignore directories, non-Go files, and test files.
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".go") || strings.HasSuffix(info.Name(), "_test.go") {
			return nil
		}

		// Parse the Go source file.
		f, parseErr := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if parseErr != nil {
			return parseErr
		}

		// Ensure the file belongs to the specified package.
		if f.Name.Name != pkgName {
			return nil
		}

		// Iterate over all declarations within the file.
		for _, decl := range f.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok || genDecl.Tok != token.TYPE {
				continue
			}

			// Process type specifications within the declaration.
			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				structType, ok := typeSpec.Type.(*ast.StructType)
				if !ok || !typeSpec.Name.IsExported() {
					continue
				}

				// Retrieve documentation comments for the struct, if any.
				var structComment string
				if genDecl.Doc != nil {
					structComment = genDecl.Doc.Text()
				}

				// Collect information about fields within the struct.
				structFields := make([]FieldInfo, 0)
				for _, field := range structType.Fields.List {
					fieldType := ExprToString(field.Type)
					var fieldComment string
					if field.Doc != nil {
						fieldComment = field.Doc.Text()
					}

					// Add each field to the struct's field list.
					for _, fieldName := range field.Names {
						structFields = append(structFields, FieldInfo{
							Name:    fieldName.Name,
							Type:    fieldType,
							Comment: fieldComment,
						})
					}
				}

				// Store the struct information in the map.
				structInfo := StructInfo{
					Name:    typeSpec.Name.Name,
					Fields:  structFields,
					Comment: structComment,
				}
				structsMap[path] = append(structsMap[path], structInfo)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return structsMap, nil
}
