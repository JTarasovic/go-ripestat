package client

func ExampleClient_RIRPrefixSizeDistribution() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "192/23" // prefix
	queryTime := ""
	r, err := c.RIRPrefixSizeDistribution(resource, queryTime)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// query_time
	// resource
	// rirs
}
