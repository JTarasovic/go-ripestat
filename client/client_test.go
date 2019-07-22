package client

import (
	"fmt"
	"sort"
	"testing"
)

func TestCleanMap(t *testing.T) {
	m := make(map[string]string)
	m["empty"] = ""
	m["something"] = "something"
	m2 := cleanMap(m)
	if _, ok := m2["empty"]; ok {
		t.Errorf("empty value not removed")
	}
}

func printMapKeysSorted(m map[string]interface{}) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k)
	}
}
