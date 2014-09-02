package main

import (
	"fmt"
	"testing"
)

func TestSplitNamesParameter(t *testing.T) {
	concatedNames := "john bob mary"
	expectedNames := []string{"john", "bob", "mary"}

	if x := splitNamesParameter(concatedNames); fmt.Sprintf("%v", x) == fmt.Sprintf("%v", concatedNames) {
		t.Errorf("splitNamesParameter(%v) = %v, want %v", concatedNames, x, expectedNames)
	}
}

func TestConcatNames(t *testing.T) {
	names := []string{"john", "bob", "mary"}
	expectedConcatedNames := "john, bob, mary"

	if x := concatNames(names); x != expectedConcatedNames {
		t.Errorf("concatNames(%v) = %v, want %v", names, x, expectedConcatedNames)
	}
}
