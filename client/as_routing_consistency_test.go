package client

func ExampleClient_ASRoutingConsistency() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // ASN
	r, err := c.ASRoutingConsistency(resource)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// authority
	// exports
	// imports
	// prefixes
	// query_endtime
	// query_starttime
	// resource
}
