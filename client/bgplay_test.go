package client

import "fmt"

func ExampleClient_BGPlay() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496"  // Prefix, IP address, AS or a list of valid comma-separated resources
	starttime := ""      // ISO8601 or Unix timestamp
	endtime := ""        // ISO8601 or Unix timestamp
	rrcs := ""           // Single-value or comma-separated values of RRC numbers (4 or 0,4,12,15)
	unixTimestamps := "" // format the timestamps in the result as Unix timestamp.
	r, err := c.BGPlay(resource, starttime, endtime, rrcs, unixTimestamps)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_starttime")
	delete(data, "query_endtime")
	fmt.Printf("%q", data)
	// Output:
	// map["events":[] "initial_state":[] "nodes":[] "resource":"64496" "sources":[] "targets":[]]
}
