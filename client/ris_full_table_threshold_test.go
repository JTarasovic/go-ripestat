package client

func ExampleClient_RISFullTableThreshold() {
	c := NewClient().WithSourceApp("go-ripestat")
	queryTime := ""
	r, err := c.RISFullTableThreshold(queryTime)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// earliest_time
	// latest_time
	// parameters
	// result_time
	// v4
	// v6
}
