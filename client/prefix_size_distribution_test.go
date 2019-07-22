package client

import "fmt"

func ExampleClient_PrefixSizeDistribution() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // ASN
	timestamp := ""
	minPeersSeeing := ""
	r, err := c.PrefixSizeDistribution(resource, timestamp, minPeersSeeing)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_time")
	fmt.Printf("%q", data)
	// Output:
	// map["ipv4":[] "ipv6":[] "resource":"64496"]
}
