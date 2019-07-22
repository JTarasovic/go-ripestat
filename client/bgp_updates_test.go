package client

import "fmt"

func ExampleClient_BGPUpdates() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496"  // Prefix, IP address, AS or a list of valid comma-separated resources
	starttime := ""      // ISO8601 or Unix timestamp
	endtime := ""        // ISO8601 or Unix timestamp
	rrcs := ""           // Single-value or comma-separated values of RRC numbers (4 or 0,4,12,15)
	unixTimestamps := "" // format the timestamps in the result as Unix timestamp.
	r, err := c.BGPUpdates(resource, starttime, endtime, rrcs, unixTimestamps)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_starttime")
	delete(data, "query_endtime")
	fmt.Printf("%q", data)
	// Output:
	// map["nr_updates":%!q(float64=0) "resource":"64496" "updates":[]]
}
