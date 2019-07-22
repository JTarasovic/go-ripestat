package client

func ExampleClient_SpeedcheckerBandwidthMeasurements() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "193/23" // prefix
	starttime := ""
	endtime := ""
	r, err := c.SpeedcheckerBandwidthMeasurements(resource, starttime, endtime)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// earliest_time
	// endtime
	// latest_time
	// measurements
	// resource
	// starttime
	// statistics
}
