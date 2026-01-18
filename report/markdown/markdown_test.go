package markdown

import (
	"go/token"
	"strings"
	"testing"

	"github.com/loov/dreamlint/report"
)

func TestWrite(t *testing.T) {
	r := report.NewReport()
	r.Metadata.Modules = []string{"testpkg"}

	r.Units["testpkg.Hello"] = report.UnitReport{
		Functions: []report.FunctionInfo{{
			Package:   "testpkg",
			Name:      "Hello",
			Signature: "func Hello(name string) string",
			Position:  token.Position{Filename: "main.go", Line: 10},
		}},
		Summary: report.FunctionSummary{
			Purpose:  "Returns a greeting",
			Behavior: "Concatenates strings",
		},
	}

	r.AddIssue("testpkg.Hello", report.Issue{
		Position:   token.Position{Filename: "main.go", Line: 12},
		Severity:   report.SeverityCritical,
		Category:   "security",
		Message:    "SQL injection vulnerability",
		Suggestion: "Use parameterized queries",
	})

	md := Write(r)

	// Check key sections exist
	if !strings.Contains(md, "# Code Review Report") {
		t.Error("missing title")
	}

	if !strings.Contains(md, "Critical") {
		t.Error("missing severity")
	}

	if !strings.Contains(md, "SQL injection") {
		t.Error("missing issue message")
	}

	if !strings.Contains(md, "Returns a greeting") {
		t.Error("missing function purpose")
	}
}
