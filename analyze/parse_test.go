// analyze/parse_test.go
package analyze

import (
	"testing"
)

func TestParseSummaryResponse(t *testing.T) {
	response := `{
  "purpose": "Fetches user by ID from database",
  "behavior": "Returns nil, ErrNotFound if user doesn't exist",
  "invariants": ["db must not be nil"],
  "security": ["input id is used in SQL query"]
}`

	summary, err := ParseSummaryResponse(response)
	if err != nil {
		t.Fatalf("ParseSummaryResponse: %v", err)
	}

	if summary.Purpose != "Fetches user by ID from database" {
		t.Errorf("purpose = %q", summary.Purpose)
	}

	if len(summary.Invariants) != 1 {
		t.Errorf("invariants count = %d, want 1", len(summary.Invariants))
	}
}

func TestParseIssuesResponse(t *testing.T) {
	response := `{
  "issues": [
    {
      "line": 42,
      "severity": "high",
      "message": "SQL query built with string concatenation",
      "suggestion": "Use parameterized query"
    }
  ]
}`

	issues, err := ParseIssuesResponse(response)
	if err != nil {
		t.Fatalf("ParseIssuesResponse: %v", err)
	}

	if len(issues) != 1 {
		t.Fatalf("issues count = %d, want 1", len(issues))
	}

	if issues[0].Severity != "high" {
		t.Errorf("severity = %q, want high", issues[0].Severity)
	}

	if issues[0].Line != 42 {
		t.Errorf("line = %d, want 42", issues[0].Line)
	}
}

func TestParseIssuesResponse_WithMarkdown(t *testing.T) {
	// LLMs sometimes wrap JSON in markdown code blocks
	response := "```json\n" + `{
  "issues": [
    {"line": 10, "severity": "medium", "message": "test"}
  ]
}` + "\n```"

	issues, err := ParseIssuesResponse(response)
	if err != nil {
		t.Fatalf("ParseIssuesResponse: %v", err)
	}

	if len(issues) != 1 {
		t.Fatalf("issues count = %d, want 1", len(issues))
	}
}
