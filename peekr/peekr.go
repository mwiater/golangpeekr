package peekr

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/mwiater/peekr/helpers"
)

// FunctionInfo holds information about a function and its associated comments.
type FunctionInfo struct {
	FileName string
	Function string
	Comments string
	Params   string // Add a field for the parameters
	Returns  string // Add a field for the return types
}

// StructInfo holds information about a struct and its fields.
type StructInfo struct {
	FileName string
	Struct   string
	Fields   []FieldInfo
}

// FieldInfo holds information about a field within a struct.
type FieldInfo struct {
	Name    string
	Type    string
	Comment string
}

// PackageFunctions lists and sorts the functions in a package, excluding peekr.go and peekr_test.go.
func PackageFunctions(dir string) (map[string][]FunctionInfo, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	funcMap := make(map[string][]FunctionInfo)

	for _, pkg := range pkgs {
		for filePath, f := range pkg.Files {
			// Skip peekr.go and peekr_test.go
			fileName := filepath.Base(filePath)
			if fileName == "peekr.go" || fileName == "peekr_test.go" {
				continue
			}

			// Use the file name without the extension for grouping
			groupName := strings.TrimSuffix(fileName, filepath.Ext(fileName))

			for _, decl := range f.Decls {
				if fn, ok := decl.(*ast.FuncDecl); ok && fn.Name.IsExported() {
					// Get the comments associated with the function
					var comments string
					if fn.Doc != nil {
						comments = Commentify(fn.Doc.Text())
					}

					// Extract the function's parameters and return types
					params := extractFuncParams(fn.Type.Params)
					returns := extractFuncResults(fn.Type.Results)

					funcInfo := FunctionInfo{
						FileName: groupName,
						Function: fn.Name.Name,
						Comments: comments,
						Params:   params,
						Returns:  returns,
					}
					funcMap[groupName] = append(funcMap[groupName], funcInfo)
				}
			}
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

// extractFuncParams takes an *ast.FieldList (parameters) and returns a string representation.
func extractFuncParams(fl *ast.FieldList) string {
	if fl == nil {
		return ""
	}
	var params []string
	for _, field := range fl.List {
		typeString := exprToString(field.Type)
		if len(field.Names) > 0 {
			for _, name := range field.Names {
				params = append(params, fmt.Sprintf("%s %s", name, typeString))
			}
		} else {
			params = append(params, typeString)
		}
	}
	return strings.Join(params, ", ")
}

// extractFuncResults takes an *ast.FieldList (results) and returns a string representation.
func extractFuncResults(fl *ast.FieldList) string {
	if fl == nil {
		return ""
	}
	var results []string
	for _, field := range fl.List {
		typeString := exprToString(field.Type)
		if len(field.Names) > 0 {
			for _, name := range field.Names {
				results = append(results, fmt.Sprintf("%s %s", name, typeString))
			}
		} else {
			results = append(results, typeString)
		}
	}
	if len(results) == 1 {
		return results[0] // Single unnamed return value
	}
	return "(" + strings.Join(results, ", ") + ")"
}

// exprToString takes an ast.Expr (expression) and returns its string representation.
func exprToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return exprToString(t.X) + "." + t.Sel.Name
	case *ast.StarExpr:
		return "*" + exprToString(t.X)
	case *ast.ArrayType:
		return "[]" + exprToString(t.Elt)
	case *ast.InterfaceType:
		if t.Methods != nil && t.Methods.List != nil && len(t.Methods.List) > 0 {
			return "interface{...}"
		}
		return "interface{}"
	default:
		return fmt.Sprintf("%#v", expr)
	}
}

// PrintSortedFunctions prints the sorted functions as specified.
func ListPackageFunctions(dir string) {
	groupedFunctions, err := PackageFunctions(dir)
	if err != nil {
		panic(err)
	}

	var groupNames []string
	for groupName := range groupedFunctions {
		groupNames = append(groupNames, groupName)
	}
	sort.Strings(groupNames)

	cleanDir := path.Clean(dir)
	lastPart := path.Base(cleanDir)

	header := fmt.Sprintf("\nFunctions in the %s package:\n", fmt.Sprintf("'%s'", lastPart))
	helpers.TerminalColor(header, helpers.Notice)

	for _, groupName := range groupNames {
		functions := groupedFunctions[groupName]
		helpers.TerminalColor("[ "+groupName+" ]\n", helpers.Info)
		for _, funcInfo := range functions {
			helpers.TerminalColor("  "+strings.TrimSpace(funcInfo.Comments), helpers.Info)
			signature := fmt.Sprintf("  %s(%s) %s", funcInfo.Function, funcInfo.Params, funcInfo.Returns)
			helpers.TerminalColor(signature, helpers.Debug)
			fmt.Println()
		}
	}
}

// ListPackageStructs lists and prints the structs in a package, excluding peekr.go and peekr_test.go.
func ListPackageStructs(dir string) error {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	cleanDir := path.Clean(dir)
	lastPart := path.Base(cleanDir)

	header := fmt.Sprintf("\nStructs in the %s package:\n", fmt.Sprintf("'%s'", lastPart))
	helpers.TerminalColor(header, helpers.Notice)

	for _, pkg := range pkgs {
		for filePath, f := range pkg.Files {
			// Skip peekr.go and peekr_test.go
			fileName := filepath.Base(filePath)
			if fileName == "peekr.go" || fileName == "peekr_test.go" {
				continue
			}

			// Use the file name without the extension for grouping
			groupName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
			structsPrinted := false

			for _, decl := range f.Decls {
				genDecl, ok := decl.(*ast.GenDecl)
				if !ok || genDecl.Tok != token.TYPE {
					continue
				}

				for _, spec := range genDecl.Specs {
					typeSpec, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}

					structType, ok := typeSpec.Type.(*ast.StructType)
					if !ok || !typeSpec.Name.IsExported() {
						continue
					}

					if !structsPrinted {
						helpers.TerminalColor("[ "+groupName+" ]\n", helpers.Info)
						structsPrinted = true
					}

					var structComment string
					if genDecl.Doc != nil {
						structComment = Commentify(genDecl.Doc.Text())
					}

					if structComment != "" {
						helpers.TerminalColor("  "+strings.TrimSpace(structComment), helpers.Info)
					}
					//helpers.TerminalColor("  Struct: "+strings.TrimSpace(typeSpec.Name.Name), helpers.Alert)

					for _, field := range structType.Fields.List {
						fieldType := exprToString(field.Type)
						for _, fieldName := range field.Names {
							field := fmt.Sprintf("%s %s\n", fieldName.Name, fieldType)
							helpers.TerminalColor("  "+strings.TrimSpace(field), helpers.Debug)
						}
					}
					fmt.Println()
				}
			}
		}
	}

	return nil
}

// Commentify takes a string and returns it as a commented block of text,
// without adding comment syntax to the final newline if it exists.
func Commentify(str string) string {
	lines := strings.Split(str, "\n")
	for i, line := range lines {
		if i == len(lines)-1 && line == "" {
			continue // Skip the empty last line
		}
		lines[i] = "  // " + line
	}
	return strings.Join(lines, "\n")
}
