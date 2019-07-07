package efc

import (
	"go/ast"
	"go/constant"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "efc",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

const Doc = `efc reports print format which does not use "%+v" for error type.
Intend to be applied to projects which uses pkg/errors.
https://godoc.org/github.com/pkg/errors#hdr-Formatted_printing_of_errors
`

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		call := n.(*ast.CallExpr)

		format, idx := formatString(pass, call)
		if idx < 0 {
			return
		}

		argIdx := idx
		errType := types.Universe.Lookup("error").Type()

		for i := 0; i < len(format); i++ {
			if format[i] != '%' {
				continue
			}
			if i+1 < len(format) && format[i+1] == '%' {
				i++
				continue
			}

			argIdx++
			if argIdx >= len(call.Args) {
				break
			}

			arg := call.Args[argIdx]
			typ := pass.TypesInfo.Types[arg].Type

			if types.Identical(typ, errType) {
				if len(format) < i+3 {
					pass.Reportf(call.Pos(), "should use %%+v format for error type")
					break
				}
				if format[i:i+3] != "%+v" {
					pass.Reportf(call.Pos(), "should use %%+v format for error type")
					break
				}
			}
		}
	})

	return nil, nil
}

// copy from go vet printf
// https://github.com/golang/tools/blob/72ffa07ba3db8d09f5215feec0f89464f3028f8e/go/analysis/passes/printf/printf.go#L352
func formatString(pass *analysis.Pass, call *ast.CallExpr) (format string, idx int) {
	for idx := range call.Args {
		if s, ok := stringConstantArg(pass, call, idx); ok {
			return s, idx
		}
		if pass.TypesInfo.Types[call.Args[idx]].Type == types.Typ[types.String] {
			return "", -1
		}
	}
	return "", -1
}

func stringConstantArg(pass *analysis.Pass, call *ast.CallExpr, idx int) (string, bool) {
	if idx >= len(call.Args) {
		return "", false
	}
	arg := call.Args[idx]
	lit := pass.TypesInfo.Types[arg].Value
	if lit != nil && lit.Kind() == constant.String {
		return constant.StringVal(lit), true
	}
	return "", false
}
