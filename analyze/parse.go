// analyze/parse.go
package analyze

import (
	"encoding/json"
	"regexp"
	"strings"
)

// SummaryResponse is the expected JSON structure for summary pass
type SummaryResponse struct {
	Purpose    string   `json:"purpose"`
	Behavior   string   `json:"behavior"`
	Invariants []string `json:"invariants"`
	Security   []string `json:"security"`
}

// IssueResponse is a single issue from the LLM
type IssueResponse struct {
	Line       int    `json:"line"`
	Severity   string `json:"severity"`
	Message    string `json:"message"`
	Suggestion string `json:"suggestion,omitempty"`
}

// IssuesResponse is the expected JSON structure for analysis passes
type IssuesResponse struct {
	Issues []IssueResponse `json:"issues"`
}

// ParseSummaryResponse parses the LLM response for a summary pass
func ParseSummaryResponse(response string) (*SummaryResponse, error) {
	cleaned := cleanJSON(response)

	var summary SummaryResponse
	if err := json.Unmarshal([]byte(cleaned), &summary); err != nil {
		return nil, err
	}

	return &summary, nil
}

// ParseIssuesResponse parses the LLM response for an analysis pass
func ParseIssuesResponse(response string) ([]IssueResponse, error) {
	cleaned := cleanJSON(response)

	var issues IssuesResponse
	if err := json.Unmarshal([]byte(cleaned), &issues); err != nil {
		return nil, err
	}

	return issues.Issues, nil
}

// cleanJSON extracts JSON from markdown code blocks if present
func cleanJSON(s string) string {
	s = strings.TrimSpace(s)

	// Remove markdown code blocks
	codeBlockRegex := regexp.MustCompile("(?s)```(?:json)?\\s*(.+?)\\s*```")
	if matches := codeBlockRegex.FindStringSubmatch(s); len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}

	return s
}
