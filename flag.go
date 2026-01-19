package main

import "fmt"

// stringSlice is a flag.Value that collects multiple string values
type stringSlice []string

func (s *stringSlice) String() string {
	if s == nil {
		return ""
	}
	return fmt.Sprintf("%v", *s)
}

func (s *stringSlice) Set(value string) error {
	*s = append(*s, value)
	return nil
}
