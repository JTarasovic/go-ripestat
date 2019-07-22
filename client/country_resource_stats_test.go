package client

func ExampleClient_CountryResourceStats() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "at" // 2-digit ISO-3166 country code
	starttime := ""
	endtime := ""
	resolution := "" // 5m, 1h, 1d, 1w
	r, err := c.CountryResourceStats(resource, starttime, endtime, resolution)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// earliest_time
	// hd_latest_time
	// latest_time
	// query_endtime
	// query_starttime
	// resolution
	// resource
	// stats
}
