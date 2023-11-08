package peekr

import (
	"fmt"
	"go/ast"
	"strings"
)

// ExtractFuncParams converts the parameters of a function from an *ast.FieldList to a string.
// Each parameter is represented by its name followed by its type, separated by spaces.
// If a parameter is unnamed, only the type is included in the string.
func ExtractFuncParams(fl *ast.FieldList) string {
	if fl == nil {
		return ""
	}
	var params []string
	for _, field := range fl.List {
		typeString := ExprToString(field.Type)
		if len(field.Names) > 0 {
			// If the field has names, create a string for each name with the type.
			for _, name := range field.Names {
				params = append(params, fmt.Sprintf("%s %s", name, typeString))
			}
		} else {
			// If the field is unnamed, append only the type string.
			params = append(params, typeString)
		}
	}
	return strings.Join(params, ", ")
}

// ExtractFuncResults converts the result types of a function from an *ast.FieldList to a string.
// If there is only one unnamed result, it returns just the type string.
// For multiple or named results, it returns a parenthesized list separated by commas.
func ExtractFuncResults(fl *ast.FieldList) string {
	if fl == nil {
		return ""
	}
	var results []string
	for _, field := range fl.List {
		typeString := ExprToString(field.Type)
		if len(field.Names) > 0 {
			// If the field has names, create a string for each name with the type.
			for _, name := range field.Names {
				results = append(results, fmt.Sprintf("%s %s", name, typeString))
			}
		} else {
			// If the field is unnamed, append only the type string.
			results = append(results, typeString)
		}
	}

	// Format the results based on the number and naming of the return values.
	if len(results) == 1 {
		return results[0] // Single unnamed return value
	}
	return "(" + strings.Join(results, ", ") + ")"
}

// ExprToString converts an AST expression to its string representation.
// It handles different types of expressions like identifiers, selector expressions,
// pointer types, array types, and interfaces.
func ExprToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return ExprToString(t.X) + "." + t.Sel.Name
	case *ast.StarExpr:
		return "*" + ExprToString(t.X)
	case *ast.ArrayType:
		return "[]" + ExprToString(t.Elt)
	case *ast.InterfaceType:
		if t.Methods != nil && t.Methods.List != nil && len(t.Methods.List) > 0 {
			return "interface{...}"
		}
		return "interface{}"
	default:
		return fmt.Sprintf("%#v", expr)
	}
}

// Commentify formats a given string as a block of code comments.
// It ensures that the final newline does not have comment syntax if it's empty.
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
