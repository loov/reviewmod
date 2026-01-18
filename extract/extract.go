// extract/extract.go
package extract

import (
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

// ExtractFunctions extracts all function information from loaded packages.
func ExtractFunctions(p *Packages) []*FunctionInfo {
	return ExtractFunctionsFromPackages(p.Pkgs)
}

// ExtractFunctionsFromPackages extracts function information from a slice of packages.
func ExtractFunctionsFromPackages(pkgs []*packages.Package) []*FunctionInfo {
	var funcs []*FunctionInfo

	for _, pkg := range pkgs {
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

	return funcs
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
