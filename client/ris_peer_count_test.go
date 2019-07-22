package client

func ExampleClient_RISPeerCount() {
	c := NewClient().WithSourceApp("go-ripestat")
	starttime := ""
	endtime := ""
	v4FullPrefixThreshold := ""
	v6FullPrefixThreshold := ""
	r, err := c.RISPeerCount(starttime, endtime, v4FullPrefixThreshold, v6FullPrefixThreshold)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// endtime
	// peer_count
	// starttime
	// v4_full_prefix_threshold
	// v6_full_prefix_threshold
}
