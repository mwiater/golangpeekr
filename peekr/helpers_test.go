package peekr

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractFuncParams(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "No params",
			input:    "package test\nfunc foo() {}",
			expected: "",
		},
		{
			name:     "One param",
			input:    "package test\nfunc foo(x int) {}",
			expected: "x int",
		},
		{
			name:     "Multiple params",
			input:    "package test\nfunc foo(x int, y string) {}",
			expected: "x int, y string",
		},
		{
			name:     "Unnamed params",
			input:    "package test\nfunc foo(int, string) {}",
			expected: "int, string",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fset := token.NewFileSet()
			file, err := parser.ParseFile(fset, "", tc.input, 0)
			if err != nil {
				t.Fatalf("Failed to parse input: %s", err)
			}

			fn := file.Decls[0].(*ast.FuncDecl)
			actual := ExtractFuncParams(fn.Type.Params)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestExtractFuncResults(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "No results",
			input:    "package test\nfunc foo() {}",
			expected: "",
		},
		{
			name:     "One result",
			input:    "package test\nfunc foo() int { return 0 }",
			expected: "int",
		},
		{
			name:     "Multiple results",
			input:    "package test\nfunc foo() (int, string) { return 0, \"\" }",
			expected: "(int, string)",
		},
		{
			name:     "Named results",
			input:    "package test\nfunc foo() (x int, y string) { return 0, \"\" }",
			expected: "(x int, y string)",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fset := token.NewFileSet()
			file, err := parser.ParseFile(fset, "", tc.input, 0)
			if err != nil {
				t.Fatalf("Failed to parse input: %s", err)
			}

			fn := file.Decls[0].(*ast.FuncDecl)
			actual := ExtractFuncResults(fn.Type.Results)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestExprToString(t *testing.T) {
	testCases := []struct {
		name     string
		input    ast.Expr
		expected string
	}{
		{
			name:     "Identifier",
			input:    &ast.Ident{Name: "int"},
			expected: "int",
		},
		// Add more cases for other types like *ast.SelectorExpr, *ast.StarExpr, etc.
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := ExprToString(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestCommentify(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Single line comment",
			input:    "Test comment",
			expected: "  // Test comment",
		},
		{
			name:     "Multi-line comment",
			input:    "Line one\nLine two\n",
			expected: "  // Line one\n  // Line two",
		},
		{
			name:     "Empty comment",
			input:    "",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Commentify(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
