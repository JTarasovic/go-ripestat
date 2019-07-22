package client

func ExampleClient_RoutingHistory() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // prefix or ASN
	maxRows := ""
	includeFirstHop := ""
	normaliseVisibility := ""
	minPeers := ""
	starttime := ""
	endtime := ""
	r, err := c.RoutingHistory(resource, maxRows, includeFirstHop, normaliseVisibility, minPeers, starttime, endtime)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// by_origin
	// query_endtime
	// query_starttime
	// resource
	// time_granularity
}
