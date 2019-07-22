package client

import (
	"reflect"
	"testing"
)

func ExampleClient_RPKIHistory() {
	c := NewClient().WithSourceApp("go-ripestat")
	prefix := "prefix_cwe=193.0/16"
	prefixFamily := "4"
	prefixMaskLength := ""
	maxLength := ""
	asn := "3333"
	trustAnchor := ""
	fields := "prefix_count" // "address_space", "asn_count", "prefix_count", "roa_count"
	r, err := c.RPKIHistory(prefix, prefixFamily, prefixMaskLength, maxLength, asn, trustAnchor, fields)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// asn
	// fields
	// prefix_cwe
	// prefix_family
	// sourceapp
	// timeseries
}

func TestParseRPKIHistoryPrefix(t *testing.T) {
	var tests = []struct {
		name   string
		input  string
		output map[string]string
	}{
		{
			name:   "empty-string",
			input:  "",
			output: map[string]string{},
		},
		{
			name:  "equal",
			input: "193/23",
			output: map[string]string{
				"prefix": "193/23",
			},
		},
		{
			name:  "contained-within-or-equal",
			input: "prefix_cwe=193/23",
			output: map[string]string{
				"prefix_cwe": "193/23",
			},
		},
		{
			name:  "contained-within",
			input: "prefix_cw=193/23",
			output: map[string]string{
				"prefix_cw": "193/23",
			},
		},
		{
			name:  "contained-or-equal",
			input: "prefix_ce=193/23",
			output: map[string]string{
				"prefix_ce": "193/23",
			},
		},
		{
			name:  "contained",
			input: "prefix_c=193/23",
			output: map[string]string{
				"prefix_c": "193/23",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := make(map[string]string)
			parseRPKIHistoryPrefix(tt.input, m)
			if !reflect.DeepEqual(m, tt.output) {
				t.Errorf("output is not correct for case: %s", tt.name)
			}
		})
	}

}
