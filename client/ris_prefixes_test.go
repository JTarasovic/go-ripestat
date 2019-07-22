package client

func ExampleClient_RISPrefixes() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // ASN
	queryTime := ""
	listPrefixes := ""
	types := "" // "o" will show originating prefixes and "t" transitting. The combination shows both, which is the default.
	af := ""    // "v4","v6" or "v4,v6"
	noise := "" // "keep" or "filter"
	r, err := c.RISPrefixes(resource, queryTime, listPrefixes, types, af, noise)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// af
	// counts
	// earliest_time
	// latest_time
	// list_prefixes
	// noise
	// query_time
	// resource
	// types
}
