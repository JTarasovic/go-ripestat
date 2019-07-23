package client

func ExampleClient_ASNNeighbours() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // ASN
	starttime := ""     // ISO8601 or Unix timestamp for query start
	r, err := c.ASNNeighbours(resource, starttime)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// earliest_time
	// latest_time
	// neighbour_counts
	// neighbours
	// query_endtime
	// query_starttime
	// resource
}
