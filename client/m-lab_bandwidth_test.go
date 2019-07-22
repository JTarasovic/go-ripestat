package client

func ExampleClient_MlabBandwidth() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "at" // IPv4 Prefix, IPv4 address or 2-digit ISO-3166 country code
	starttime := ""
	endtime := ""
	r, err := c.MlabBandwidth(resource, starttime, endtime)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// bandwidths
	// query_endtime
	// query_starttime
	// resource
}
