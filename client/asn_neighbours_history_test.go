package client

func ExampleClient_ASNNeighboursHistory() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // ASN
	starttime := ""     // ISO8601 or Unix timestamp for query start
	endtime := ""       // ISO8601 or Unix timestamp for query start
	maxRows := ""       // Defines the limit of neighbours to be included in the result, e.g. max_rows=50 means the result will be truncated to 50 neighbours.
	r, err := c.ASNNeighboursHistory(resource, starttime, endtime, maxRows)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// earliest_time
	// latest_time
	// neighbours
	// query_endtime
	// query_starttime
	// resource
}
