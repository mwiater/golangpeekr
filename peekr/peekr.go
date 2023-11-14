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

	"github.com/mwiater/peekr/config"
	"github.com/mwiater/peekr/helpers"
)

var Logger = config.GetLogger()

// Info is a common interface for items that can be printed by commonOutput.
type Info interface {
	GetFileName() string
}

// Implement the interface for FunctionInfo.
func (fi FunctionInfo) GetFileName() string {
	return fi.FileName
}

// Implement the interface for StructInfo (assuming similar fields and formatting needs).
func (si StructInfo) GetFileName() string {
	return si.Name // Assuming StructInfo has a Name field for the file name
}

// FunctionInfo holds metadata about a function within a Go source file.
// It includes the file name, function signature, associated comments,
// parameter list, and return types.
type FunctionInfo struct {
	FileName string
	Function string
	Comments string
	Params   string
	Returns  string
}

// StructInfo holds metadata about a struct type within a Go source file.
// It includes the struct name, slice of its fields, and associated comments.
type StructInfo struct {
	Name     string
	FileName string
	Fields   []FieldInfo
	Comment  string
}

// FieldInfo holds metadata about a field within a struct.
// It includes the field name, field type, and associated comments.
type FieldInfo struct {
	Name    string
	Type    string
	Comment string
}

// commonOutput handles the shared output logic.
func commonOutput(pkgName string, infoMap map[string][]Info, infoType string) {
	var groupNames []string
	for groupName := range infoMap {
		groupNames = append(groupNames, groupName)
	}
	sort.Strings(groupNames)

	if len(groupNames) > 0 {
		header := fmt.Sprintf("\n%s in the %s package:", infoType, fmt.Sprintf("'%s'", pkgName))
		helpers.TerminalColor(header, helpers.Info)
		for _, groupName := range groupNames {
			infos := infoMap[groupName]
			helpers.TerminalColor("\nFile: "+groupName+"\n", helpers.Cyan)
			for _, info := range infos {
				switch v := info.(type) {
				case FunctionInfo:
					helpers.TerminalColor(v.Comments, helpers.Cyan)
					signature := fmt.Sprintf("  %s(%s) %s", v.Function, v.Params, v.Returns)
					helpers.TerminalColor(signature, helpers.Debug)
				case StructInfo:
					helpers.TerminalColor(v.Comment, helpers.Cyan)

					maxLength := 0
					for _, field := range v.Fields {
						if len(field.Name) > maxLength {
							maxLength = len(field.Name)
						}
					}

					for _, field := range v.Fields {
						formattedField := fmt.Sprintf("  %-*s  %s", maxLength+2, field.Name, field.Type)
						helpers.TerminalColor(formattedField, helpers.Debug)
					}

				default:
					fmt.Println("Unknown type")
				}
				fmt.Println()
			}
		}
	} else {
		header := fmt.Sprintf("\nNo %s in the %s package:", infoType, fmt.Sprintf("'%s'", pkgName))
		helpers.TerminalColor(header, helpers.Error)
	}
}

// ListPackageFunctions prints a color-coded list of functions from the specified package.
// It retrieves function metadata using PackageFunctions and formats the output.
func ListPackageFunctions(dir, pkgName string) {
	functionMap, err := PackageFunctions(dir, pkgName)
	if err != nil {
		Logger.Error(err.Error())
		os.Exit(1)
	}

	infoMap := make(map[string][]Info)
	for k, v := range functionMap {
		var infos []Info
		for _, fi := range v {
			infos = append(infos, fi)
		}
		infoMap[k] = infos
	}

	commonOutput(pkgName, infoMap, "Functions")
}

// ListPackageStructs prints a color-coded list of structs from the specified package.
// It retrieves struct metadata using PackageStructs and formats the output.
func ListPackageStructs(dir, pkgName string) {
	structsMap, err := PackageStructs(dir, pkgName)
	if err != nil {
		Logger.Error(err.Error())
		os.Exit(1)
	}

	infoMap := make(map[string][]Info)
	for k, v := range structsMap {
		var infos []Info
		for _, si := range v {
			infos = append(infos, si)
		}
		infoMap[k] = infos
	}

	commonOutput(pkgName, infoMap, "Structs")
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
		if strings.Contains(err.Error(), "EOF") {
			return funcMap, nil
		} else {
			return nil, err
		}

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
					structComment = Commentify(genDecl.Doc.Text())
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
