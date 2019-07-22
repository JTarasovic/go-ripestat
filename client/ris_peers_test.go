package client

func ExampleClient_RISPeers() {
	c := NewClient().WithSourceApp("go-ripestat")
	queryTime := ""
	r, err := c.RISPeers(queryTime)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// earliest_time
	// latest_time
	// parameters
	// peers
}
