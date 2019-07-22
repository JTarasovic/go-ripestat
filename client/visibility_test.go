package client

func ExampleClient_Visibility() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // IP address or ASN
	queryTime := ""
	include := "" // optionally, "peers_seeing" includes details on peers that are seeing a resource as only the peers that are not seeing a resource.
	r, err := c.Visibility(resource, queryTime, include)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// include
	// lastest_time
	// query_time
	// related_prefixes
	// resource
	// visibilities
}
