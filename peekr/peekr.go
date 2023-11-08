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

// FunctionInfo holds information about a function and its associated comments.
type FunctionInfo struct {
	FileName string
	Function string
	Comments string
	Params   string
	Returns  string
}

// StructInfo holds information about a struct and its fields.
type StructInfo struct {
	Name    string
	Fields  []FieldInfo
	Comment string
}

// FieldInfo holds information about a field within a struct.
type FieldInfo struct {
	Name    string
	Type    string
	Comment string
}

func PackageFunctions(pkgName string) (map[string][]FunctionInfo, error) {
	fset := token.NewFileSet()
	funcMap := make(map[string][]FunctionInfo)

	// Walk the directory tree recursively
	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and test files
		if info.IsDir() || strings.HasSuffix(info.Name(), "_test.go") {
			return nil
		}

		// Process only Go files
		if strings.HasSuffix(info.Name(), ".go") {
			// Parse the Go source file
			f, parseErr := parser.ParseFile(fset, path, nil, parser.ParseComments)
			if parseErr != nil {
				return parseErr
			}

			// Check if the file's package name matches the desired package
			if f.Name.Name != pkgName {
				return nil
			}

			// Use the file name without the extension for grouping
			groupName := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))

			// Process the file's declarations
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

// ListPackageFunctions prints the sorted functions as specified.
func ListPackageFunctions(pkgName string) {
	groupedFunctions, err := PackageFunctions(pkgName)
	if err != nil {
		panic(err)
	}

	var groupNames []string
	for groupName := range groupedFunctions {
		groupNames = append(groupNames, groupName)
	}
	sort.Strings(groupNames)

	header := fmt.Sprintf("\nFunctions in the %s package:", fmt.Sprintf("'%s'", pkgName))
	helpers.TerminalColor(header, helpers.Notice)

	for _, groupName := range groupNames {
		functions := groupedFunctions[groupName]
		helpers.TerminalColor("\nFile: "+groupName+"\n", helpers.Info)
		for _, funcInfo := range functions {
			helpers.TerminalColor("  "+strings.TrimSpace(funcInfo.Comments), helpers.Info)
			signature := fmt.Sprintf("  %s(%s) %s", funcInfo.Function, funcInfo.Params, funcInfo.Returns)
			helpers.TerminalColor(signature, helpers.Debug)
			fmt.Println()
		}
	}
}

// PackageStructs gathers data about structs in the specified package.
func PackageStructs(pkgName string) (map[string][]StructInfo, error) {
	fset := token.NewFileSet()
	structsMap := make(map[string][]StructInfo)

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories, non-Go files, and test files
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".go") || strings.HasSuffix(info.Name(), "_test.go") {
			return nil
		}

		// Parse the Go source file
		f, parseErr := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if parseErr != nil {
			return parseErr
		}

		// Check if the file's package name matches the desired package
		if f.Name.Name != pkgName {
			return nil
		}

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

				var structComment string
				if genDecl.Doc != nil {
					structComment = genDecl.Doc.Text()
				}

				structFields := make([]FieldInfo, 0)
				for _, field := range structType.Fields.List {
					fieldType := exprToString(field.Type)
					var fieldComment string
					if field.Doc != nil {
						fieldComment = field.Doc.Text()
					}

					for _, fieldName := range field.Names {
						structFields = append(structFields, FieldInfo{
							Name:    fieldName.Name,
							Type:    fieldType,
							Comment: fieldComment,
						})
					}
				}

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

// ListPackageStructs formats and prints the structs obtained from PackageStructs using color-coded output.
func ListPackageStructs(pkgName string) error {
	structsMap, err := PackageStructs(pkgName)
	if err != nil {
		return err
	}

	// Sort the groups by name
	var groupNames []string
	for groupName := range structsMap {
		groupNames = append(groupNames, groupName)
	}
	sort.Strings(groupNames)

	header := fmt.Sprintf("\nStructs in the %s package:", fmt.Sprintf("'%s'", pkgName))
	helpers.TerminalColor(header, helpers.Notice)

	for _, groupName := range groupNames {
		structs := structsMap[groupName]
		if len(structs) > 0 {
			helpers.TerminalColor("\nFile: "+groupName+"\n", helpers.Info)
			for _, structInfo := range structs {
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
