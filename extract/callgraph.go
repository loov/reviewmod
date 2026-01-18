// extract/callgraph.go
package extract

import (
	"fmt"

	"golang.org/x/tools/go/callgraph/cha"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

// BuildCallgraph builds a callgraph using CHA analysis
// Returns a map from function ID to list of callee IDs
func BuildCallgraph(dir string, patterns ...string) (map[string][]string, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedSyntax |
			packages.NeedTypes |
			packages.NeedTypesInfo |
			packages.NeedImports |
			packages.NeedDeps,
		Dir: dir,
	}

	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		return nil, fmt.Errorf("load packages: %w", err)
	}

	for _, pkg := range pkgs {
		if len(pkg.Errors) > 0 {
			return nil, fmt.Errorf("package %s has errors: %v", pkg.PkgPath, pkg.Errors)
		}
	}

	// Build SSA
	prog, _ := ssautil.AllPackages(pkgs, ssa.SanityCheckFunctions)
	prog.Build()

	// Build callgraph using CHA
	cg := cha.CallGraph(prog)

	// Convert to our format
	graph := make(map[string][]string)

	for fn, node := range cg.Nodes {
		if fn == nil {
			continue
		}

		callerID := funcID(fn)
		if callerID == "" {
			continue
		}

		// Initialize entry even if no callees
		if _, ok := graph[callerID]; !ok {
			graph[callerID] = []string{}
		}

		for _, edge := range node.Out {
			if edge.Callee.Func == nil {
				continue
			}

			calleeID := funcID(edge.Callee.Func)
			if calleeID == "" {
				continue
			}

			// Avoid duplicates
			if !contains(graph[callerID], calleeID) {
				graph[callerID] = append(graph[callerID], calleeID)
			}
		}
	}

	return graph, nil
}

func funcID(fn *ssa.Function) string {
	if fn.Pkg == nil {
		return ""
	}

	pkg := fn.Pkg.Pkg.Path()
	name := fn.Name()

	// Handle methods
	if recv := fn.Signature.Recv(); recv != nil {
		return fmt.Sprintf("%s.(%s).%s", pkg, recv.Type().String(), name)
	}

	return fmt.Sprintf("%s.%s", pkg, name)
}

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
