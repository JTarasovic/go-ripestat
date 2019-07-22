package client

import "fmt"

func ExampleClient_BGPUpdateActivity() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496"     // prefix, IP range or AS
	starttime := ""         // ISO8601 or Unix timestamp
	endtime := ""           // ISO8601 or Unix timestamp
	maxSamples := ""        // Positive integer, or 0 to disable
	minSamplingPeriod := "" // The smallest possible time period for each data point. It will be automatically increased to satisfy 'max_points'
	numHours := ""          // If 'starttime' is not given then this parameter will be used to calculate it from the 'endtime' (which defaults to now).
	hideEmptySamples := ""  // If true (default) then samples with 0 updates will not be returned
	r, err := c.BGPUpdateActivity(resource, starttime, endtime, maxSamples, minSamplingPeriod, numHours, hideEmptySamples)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_starttime")
	delete(data, "query_endtime")
	fmt.Printf("%q", data)
	// Output:
	// map["max_samples":%!q(float64=50) "related_prefixes":[] "resource":"64496" "resource_type":"asn" "sampling_period":%!q(float64=21600) "sampling_period_human":"6 hours" "updates":[]]
}
