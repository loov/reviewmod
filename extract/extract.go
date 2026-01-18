// extract/extract.go
package extract

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"

	"golang.org/x/tools/go/packages"
)

// FunctionInfo holds information about a single function
type FunctionInfo struct {
	Package   string
	Name      string
	Receiver  string
	Signature string
	Body      string
	Godoc     string
	Position  token.Position
}

// ExtractFunctions loads packages and extracts all function information
func ExtractFunctions(dir string, patterns ...string) ([]*FunctionInfo, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedSyntax |
			packages.NeedTypes |
			packages.NeedTypesInfo,
		Dir: dir,
	}

	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		return nil, fmt.Errorf("load packages: %w", err)
	}

	var funcs []*FunctionInfo

	for _, pkg := range pkgs {
		if len(pkg.Errors) > 0 {
			return nil, fmt.Errorf("package %s has errors: %v", pkg.PkgPath, pkg.Errors)
		}

		for _, file := range pkg.Syntax {
			for _, decl := range file.Decls {
				fn, ok := decl.(*ast.FuncDecl)
				if !ok {
					continue
				}

				info := &FunctionInfo{
					Package:  pkg.PkgPath,
					Name:     fn.Name.Name,
					Position: pkg.Fset.Position(fn.Pos()),
				}

				// Extract receiver
				if fn.Recv != nil && len(fn.Recv.List) > 0 {
					var buf strings.Builder
					printer.Fprint(&buf, pkg.Fset, fn.Recv.List[0].Type)
					info.Receiver = buf.String()
				}

				// Extract signature
				info.Signature = formatSignature(pkg.Fset, fn)

				// Extract body
				if fn.Body != nil {
					var buf strings.Builder
					printer.Fprint(&buf, pkg.Fset, fn.Body)
					info.Body = buf.String()
				}

				// Extract godoc
				if fn.Doc != nil {
					info.Godoc = fn.Doc.Text()
				}

				funcs = append(funcs, info)
			}
		}
	}

	return funcs, nil
}

func formatSignature(fset *token.FileSet, fn *ast.FuncDecl) string {
	var buf strings.Builder
	buf.WriteString("func ")

	if fn.Recv != nil && len(fn.Recv.List) > 0 {
		buf.WriteString("(")
		printer.Fprint(&buf, fset, fn.Recv.List[0].Type)
		buf.WriteString(") ")
	}

	buf.WriteString(fn.Name.Name)
	printer.Fprint(&buf, fset, fn.Type)

	return buf.String()
}
