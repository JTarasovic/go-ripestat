package client

func ExampleClient_RISAsns() {
	c := NewClient().WithSourceApp("go-ripestat")
	queryTime := ""
	listASNs := ""
	asnTypes := "" // "o" stands for originating and will show originating ASNs separately. "t" does the same for transitting ASNs (keep in mind the definition of a transit in this case).
	r, err := c.RISAsns(queryTime, listASNs, asnTypes)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// counts
	// earliest_time
	// latest_time
	// list_asns
	// query_time
}
