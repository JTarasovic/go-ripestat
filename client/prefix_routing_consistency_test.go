package client

import "fmt"

func ExampleClient_PrefixRoutingConsistency() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "193/23" // prefix
	r, err := c.PrefixRoutingConsistency(resource)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_starttime")
	delete(data, "query_endtime")
	fmt.Printf("%q", data)
	// Output:
	// map["resource":"193.0.0.0/23" "routes":[map["asn_name":"RIPE-NCC-AS - Reseaux IP Europeens Network Coordination Centre (RIPE NCC)" "in_bgp":%!q(bool=true) "in_whois":%!q(bool=true) "irr_sources":["RIPE" "RIPE"] "origin":%!q(float64=3333) "prefix":"193.0.0.0/21"]]]
}
