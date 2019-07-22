package client

func ExampleClient_PrefixCount() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // ASN
	starttime := ""
	endtime := ""
	minPeersSeeing := ""
	resolution := "" // 8h, 2d, 12d
	r, err := c.PrefixCount(resource, starttime, endtime, minPeersSeeing, resolution)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// ipv4
	// ipv6
	// query_endtime
	// query_starttime
	// resolution
	// resource
}
