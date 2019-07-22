package client

import "fmt"

func ExampleClient_RoutingStatus() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // prefix, IP address or ASN
	timestamp := ""
	minPeersSeeing := ""
	r, err := c.RoutingStatus(resource, timestamp, minPeersSeeing)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_starttime")
	delete(data, "query_endtime")
	fmt.Printf("%q", data)
	// Output:
	// map["announced_space":map["v4":map["ips":%!q(float64=0) "prefixes":%!q(float64=0)] "v6":map["48s":%!q(float64=0) "prefixes":%!q(float64=0)]] "first_seen":map["origin":"64496" "prefix":"46.107.100.0/24" "time":"2010-09-23T00:00:00"] "last_seen":map["origin":"64496" "prefix":"198.57.8.0/24" "time":"2014-10-21T00:00:00"] "observed_neighbours":%!q(float64=9) "query_time":"2019-07-22T08:00:00" "resource":"64496" "visibility":map["v4":map["ris_peers_seeing":%!q(float64=0) "total_ris_peers":%!q(float64=240)] "v6":map["ris_peers_seeing":%!q(float64=0) "total_ris_peers":%!q(float64=230)]]]
}
