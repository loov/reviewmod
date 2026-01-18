// report/json_test.go
package report

import (
	"encoding/json"
	"go/token"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	r := NewReport()
	r.Metadata.Modules = []string{"testpkg"}
	r.Metadata.TotalUnits = 1

	r.Units["testpkg.Hello"] = UnitReport{
		Functions: []FunctionInfo{{
			Package:   "testpkg",
			Name:      "Hello",
			Signature: "func Hello(name string) string",
			Position:  token.Position{Filename: "main.go", Line: 10},
		}},
		Summary: FunctionSummary{
			Purpose:  "Returns a greeting",
			Behavior: "Concatenates 'Hello, ' with name",
		},
	}

	r.AddIssue("testpkg.Hello", Issue{
		Position: token.Position{Filename: "main.go", Line: 12},
		Severity: SeverityMedium,
		Category: "cleanliness",
		Message:  "Consider using fmt.Sprintf",
	})

	data, err := WriteJSON(r)
	if err != nil {
		t.Fatalf("WriteJSON: %v", err)
	}

	// Verify it's valid JSON
	var parsed Report
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if parsed.Summary.TotalIssues != 1 {
		t.Errorf("total issues = %d, want 1", parsed.Summary.TotalIssues)
	}
}
