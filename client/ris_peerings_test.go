package client

func ExampleClient_RISPeerings() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "192/23" // prefix
	queryTime := ""
	r, err := c.RISPeerings(resource, queryTime)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// peerings
	// query_endtime
	// query_starttime
	// resource
}
