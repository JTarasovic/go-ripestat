package client

import (
	"fmt"
)

func ExampleClient_AnnouncedPrefixes() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496"  // ASN
	starttime := ""      // ISO8601 or Unix timestamp for query start
	endtime := ""        // ISO8601 or Unix timestamp for query start
	minPeersSeeing := "" // Minimum number of RIS peers seeing the prefix for it to be included in the results.
	r, err := c.AnnouncedPrefixes(resource, starttime, endtime, minPeersSeeing)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_starttime")
	delete(data, "query_endtime")
	delete(data, "latest_time")
	fmt.Printf("%q", data)
	// Output:
	// map["earliest_time":"2000-08-01T00:00:00" "prefixes":[] "resource":"64496"]
}
