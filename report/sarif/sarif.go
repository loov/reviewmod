// Package sarif provides SARIF format output for reports.
package sarif

// SARIF format version 2.1.0
// https://docs.oasis-open.org/sarif/sarif/v2.1.0/sarif-v2.1.0.html

// Report is the root SARIF object
type Report struct {
	Schema  string `json:"$schema"`
	Version string `json:"version"`
	Runs    []Run  `json:"runs"`
}

// Run represents a single analysis run
type Run struct {
	Tool    Tool     `json:"tool"`
	Results []Result `json:"results"`
}

// Tool describes the analysis tool
type Tool struct {
	Driver Driver `json:"driver"`
}

// Driver describes the tool driver
type Driver struct {
	Name           string `json:"name"`
	Version        string `json:"version"`
	InformationURI string `json:"informationUri,omitempty"`
	Rules          []Rule `json:"rules,omitempty"`
}

// Rule describes a rule/category
type Rule struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	ShortDescription Message `json:"shortDescription"`
}

// Result represents a single issue
type Result struct {
	RuleID    string     `json:"ruleId"`
	Level     string     `json:"level"`
	Message   Message    `json:"message"`
	Locations []Location `json:"locations,omitempty"`
	Fixes     []Fix      `json:"fixes,omitempty"`
}

// Message holds a text message
type Message struct {
	Text string `json:"text"`
}

// Location describes where an issue occurs
type Location struct {
	PhysicalLocation PhysicalLocation `json:"physicalLocation"`
}

// PhysicalLocation describes file location
type PhysicalLocation struct {
	ArtifactLocation ArtifactLocation `json:"artifactLocation"`
	Region           *Region          `json:"region,omitempty"`
}

// ArtifactLocation describes the file
type ArtifactLocation struct {
	URI string `json:"uri"`
}

// Region describes the position in the file
type Region struct {
	StartLine   int `json:"startLine"`
	StartColumn int `json:"startColumn,omitempty"`
}

// Fix describes a suggested fix
type Fix struct {
	Description Message `json:"description"`
}
