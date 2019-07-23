package client

func ExampleClient_RoutingStatus() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // prefix, IP address or ASN
	timestamp := ""
	minPeersSeeing := ""
	r, err := c.RoutingStatus(resource, timestamp, minPeersSeeing)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// announced_space
	// first_seen
	// last_seen
	// observed_neighbours
	// query_time
	// resource
	// visibility
}
