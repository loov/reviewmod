// types.go
package main

import "go/token"

// Severity levels for issues
type Severity string

const (
	SeverityCritical Severity = "critical"
	SeverityHigh     Severity = "high"
	SeverityMedium   Severity = "medium"
	SeverityLow      Severity = "low"
	SeverityInfo     Severity = "info"
)

// FunctionInfo holds information about a single function
type FunctionInfo struct {
	Package   string         `json:"package"`
	Name      string         `json:"name"`
	Receiver  string         `json:"receiver,omitempty"`
	Signature string         `json:"signature"`
	Body      string         `json:"body"`
	Godoc     string         `json:"godoc,omitempty"`
	Position  token.Position `json:"position"`
}

// AnalysisUnit is the atomic unit of analysis (single func or SCC)
type AnalysisUnit struct {
	ID        string          `json:"id"`
	Functions []*FunctionInfo `json:"functions"`
	Callees   []string        `json:"callees"`
}

// ExternalFunc holds shallow info about external dependencies
type ExternalFunc struct {
	Package    string   `json:"package"`
	Name       string   `json:"name"`
	Signature  string   `json:"signature"`
	Godoc      string   `json:"godoc,omitempty"`
	Invariants []string `json:"invariants,omitempty"`
	Pitfalls   []string `json:"pitfalls,omitempty"`
}

// Summary describes a function's behavior for callers
type Summary struct {
	UnitID      string   `json:"unit_id"`
	Purpose     string   `json:"purpose"`
	Behavior    string   `json:"behavior"`
	Invariants  []string `json:"invariants,omitempty"`
	Security    []string `json:"security,omitempty"`
	ContentHash string   `json:"content_hash"`
}

// Issue represents a problem found during analysis
type Issue struct {
	UnitID     string         `json:"unit_id"`
	Position   token.Position `json:"position"`
	Severity   Severity       `json:"severity"`
	Category   string         `json:"category"`
	Message    string         `json:"message"`
	Snippet    string         `json:"snippet,omitempty"`
	Suggestion string         `json:"suggestion,omitempty"`
}
