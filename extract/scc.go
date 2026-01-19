package extract

import "sort"

// TarjanSCC computes strongly connected components using Tarjan's algorithm.
// Returns SCCs in reverse topological order (leaves first).
func TarjanSCC(graph map[string][]string) [][]string {
	var (
		index    = 0
		stack    = []string{}
		onStack  = map[string]bool{}
		indices  = map[string]int{}
		lowlinks = map[string]int{}
		sccs     = [][]string{}
	)

	var strongconnect func(v string)
	strongconnect = func(v string) {
		indices[v] = index
		lowlinks[v] = index
		index++
		stack = append(stack, v)
		onStack[v] = true

		for _, w := range graph[v] {
			if _, ok := indices[w]; !ok {
				strongconnect(w)
				if lowlinks[w] < lowlinks[v] {
					lowlinks[v] = lowlinks[w]
				}
			} else if onStack[w] {
				if indices[w] < lowlinks[v] {
					lowlinks[v] = indices[w]
				}
			}
		}

		if lowlinks[v] == indices[v] {
			scc := []string{}
			for {
				w := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				onStack[w] = false
				scc = append(scc, w)
				if w == v {
					break
				}
			}
			sort.Strings(scc) // deterministic order within SCC
			sccs = append(sccs, scc)
		}
	}

	// Get all nodes and sort for deterministic order
	nodes := make([]string, 0, len(graph))
	for v := range graph {
		nodes = append(nodes, v)
	}
	sort.Strings(nodes)

	for _, v := range nodes {
		if _, ok := indices[v]; !ok {
			strongconnect(v)
		}
	}

	return sccs
}
