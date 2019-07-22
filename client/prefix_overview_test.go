package client

func ExampleClient_PrefixOverview() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "192/23" // prefix
	minPeersSeeing := ""
	maxRelated := ""
	queryTime := ""
	r, err := c.PrefixOverview(resource, minPeersSeeing, maxRelated, queryTime)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// actual_num_related
	// announced
	// asns
	// block
	// is_less_specific
	// num_filtered_out
	// query_time
	// related_prefixes
	// resource
	// type
}
