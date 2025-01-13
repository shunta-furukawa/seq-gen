package analyzer

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
)

// Call represents a method call with its receiver and method name
type Call struct {
	Receiver string
	Method   string
}

// Analyzer analyzes Go server code to extract method call information
type Analyzer struct {
	calls []Call
}

// AnalyzeFile parses and analyzes a Go file to extract method call information
func (a *Analyzer) AnalyzeFile(filename string) error {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
	if err != nil {
		return fmt.Errorf("failed to parse file: %w", err)
	}

	conf := types.Config{Importer: nil}
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}

	_, err = conf.Check("main", fset, []*ast.File{node}, info)
	if err != nil {
		return fmt.Errorf("failed to type check: %w", err)
	}

	ast.Inspect(node, func(n ast.Node) bool {
		if callExpr, ok := n.(*ast.CallExpr); ok {
			if sel, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
				if x, ok := sel.X.(*ast.Ident); ok {
					a.calls = append(a.calls, Call{
						Receiver: x.Name,
						Method:   sel.Sel.Name,
					})
				}
			}
		}
		return true
	})

	return nil
}

// GetCalls returns the extracted method calls
func (a *Analyzer) GetCalls() []Call {
	return a.calls
}
