package client

import "fmt"

func ExampleClient_BGPState() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496"  // Prefix, IP address, AS or a list of valid comma-separated resources
	timestamp := ""      // ISO8601 or Unix timestamp
	rrcs := ""           // Single-value or comma-separated values of RRC numbers (4 or 0,4,12,15)
	unixTimestamps := "" // format the timestamps in the result as Unix timestamp.
	r, err := c.BGPState(resource, timestamp, rrcs, unixTimestamps)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_time")
	fmt.Printf("%q", data)
	// Output:
	// map["bgp_state":[] "nr_routes":%!q(float64=0) "resource":"64496"]

}
