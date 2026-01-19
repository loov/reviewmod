package extract

import (
	"reflect"
	"testing"
)

func TestTarjanSCC_Simple(t *testing.T) {
	// A -> B -> C (no cycles)
	graph := map[string][]string{
		"A": {"B"},
		"B": {"C"},
		"C": {},
	}

	sccs := TarjanSCC(graph)

	// Each node is its own SCC, order: C, B, A (reverse topo)
	expected := [][]string{{"C"}, {"B"}, {"A"}}
	if !reflect.DeepEqual(sccs, expected) {
		t.Errorf("got %v, want %v", sccs, expected)
	}
}

func TestTarjanSCC_Cycle(t *testing.T) {
	// A -> B -> C -> A (single SCC)
	graph := map[string][]string{
		"A": {"B"},
		"B": {"C"},
		"C": {"A"},
	}

	sccs := TarjanSCC(graph)

	if len(sccs) != 1 {
		t.Fatalf("got %d SCCs, want 1", len(sccs))
	}
	if len(sccs[0]) != 3 {
		t.Errorf("got SCC size %d, want 3", len(sccs[0]))
	}
}

func TestTarjanSCC_Mixed(t *testing.T) {
	// D -> A -> B -> C -> B (B-C cycle), A also -> C
	graph := map[string][]string{
		"D": {"A"},
		"A": {"B", "C"},
		"B": {"C"},
		"C": {"B"},
	}

	sccs := TarjanSCC(graph)

	// Should have: {B,C} as one SCC, then A, then D
	if len(sccs) != 3 {
		t.Fatalf("got %d SCCs, want 3", len(sccs))
	}

	// First SCC should be the cycle {B, C}
	if len(sccs[0]) != 2 {
		t.Errorf("first SCC size %d, want 2", len(sccs[0]))
	}
}
