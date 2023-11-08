package peekr

import (
	"fmt"
	"go/ast"
	"strings"
)

// ExtractFuncParams takes an *ast.FieldList (parameters) and returns a string representation.
func ExtractFuncParams(fl *ast.FieldList) string {
	if fl == nil {
		return ""
	}
	var params []string
	for _, field := range fl.List {
		typeString := ExprToString(field.Type)
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

// ExtractFuncResults takes an *ast.FieldList (results) and returns a string representation.
func ExtractFuncResults(fl *ast.FieldList) string {
	if fl == nil {
		return ""
	}
	var results []string
	for _, field := range fl.List {
		typeString := ExprToString(field.Type)
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

// ExprToString takes an ast.Expr (expression) and returns its string representation.
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
