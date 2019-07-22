package client

import "fmt"

func ExampleClient_ASNNeighbours() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // ASN
	starttime := ""     // ISO8601 or Unix timestamp for query start
	r, err := c.ASNNeighbours(resource, starttime)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_starttime")
	delete(data, "query_endtime")
	fmt.Printf("%q", data)
	// Output:
	// map["earliest_time":"2000-08-01T00:00:00" "latest_time":"2019-07-22T00:00:00" "neighbour_counts":map["left":%!q(float64=1) "right":%!q(float64=8) "uncertain":%!q(float64=0) "unique":%!q(float64=9)] "neighbours":[map["asn":%!q(float64=25091) "power":%!q(float64=16) "type":"left" "v4_peers":%!q(float64=68) "v6_peers":%!q(float64=0)] map["asn":%!q(float64=15954) "power":%!q(float64=2) "type":"right" "v4_peers":%!q(float64=50) "v6_peers":%!q(float64=0)] map["asn":%!q(float64=263152) "power":%!q(float64=2) "type":"right" "v4_peers":%!q(float64=2) "v6_peers":%!q(float64=0)] map["asn":%!q(float64=267750) "power":%!q(float64=2) "type":"right" "v4_peers":%!q(float64=2) "v6_peers":%!q(float64=0)] map["asn":%!q(float64=28135) "power":%!q(float64=2) "type":"right" "v4_peers":%!q(float64=6) "v6_peers":%!q(float64=0)] map["asn":%!q(float64=29117) "power":%!q(float64=2) "type":"right" "v4_peers":%!q(float64=2) "v6_peers":%!q(float64=0)] map["asn":%!q(float64=38147) "power":%!q(float64=2) "type":"right" "v4_peers":%!q(float64=2) "v6_peers":%!q(float64=0)] map["asn":%!q(float64=63956) "power":%!q(float64=2) "type":"right" "v4_peers":%!q(float64=2) "v6_peers":%!q(float64=0)] map["asn":%!q(float64=7642) "power":%!q(float64=2) "type":"right" "v4_peers":%!q(float64=2) "v6_peers":%!q(float64=0)]] "resource":"64496"]
}
