// report/json.go
package report

import (
	"encoding/json"
	"os"
)

// WriteJSON serializes the report to JSON
func WriteJSON(r *Report) ([]byte, error) {
	return json.MarshalIndent(r, "", "  ")
}

// WriteJSONFile writes the report to a JSON file
func WriteJSONFile(r *Report, path string) error {
	data, err := WriteJSON(r)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
