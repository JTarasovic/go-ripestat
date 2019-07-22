package client

import "fmt"

func ExampleClient_ASRoutingConsistency() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // ASN
	r, err := c.ASRoutingConsistency(resource)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_starttime")
	delete(data, "query_endtime")
	fmt.Printf("%q", data)
	// Output:
	// map["authority":"ripe" "exports":[map["in_bgp":%!q(bool=true) "in_whois":%!q(bool=false) "peer":%!q(float64=38147)] map["in_bgp":%!q(bool=true) "in_whois":%!q(bool=false) "peer":%!q(float64=267750)] map["in_bgp":%!q(bool=true) "in_whois":%!q(bool=false) "peer":%!q(float64=28135)] map["in_bgp":%!q(bool=true) "in_whois":%!q(bool=false) "peer":%!q(float64=263152)] map["in_bgp":%!q(bool=true) "in_whois":%!q(bool=false) "peer":%!q(float64=15954)] map["in_bgp":%!q(bool=true) "in_whois":%!q(bool=false) "peer":%!q(float64=25091)] map["in_bgp":%!q(bool=true) "in_whois":%!q(bool=false) "peer":%!q(float64=63956)] map["in_bgp":%!q(bool=true) "in_whois":%!q(bool=false) "peer":%!q(float64=7642)] map["in_bgp":%!q(bool=true) "in_whois":%!q(bool=false) "peer":%!q(float64=29117)]] "imports":[map["in_bgp":%!q(bool=true) "in_whois":%!q(bool=false) "peer":%!q(float64=38147)] map["in_bgp":%!q(bool=true) "in_whois":%!q(bool=false) "peer":%!q(float64=267750)] map["in_bgp":%!q(bool=true) "in_whois":%!q(bool=false) "peer":%!q(float64=28135)] map["in_bgp":%!q(bool=true) "in_whois":%!q(bool=false) "peer":%!q(float64=263152)] map["in_bgp":%!q(bool=true) "in_whois":%!q(bool=false) "peer":%!q(float64=15954)] map["in_bgp":%!q(bool=true) "in_whois":%!q(bool=false) "peer":%!q(float64=25091)] map["in_bgp":%!q(bool=true) "in_whois":%!q(bool=false) "peer":%!q(float64=63956)] map["in_bgp":%!q(bool=true) "in_whois":%!q(bool=false) "peer":%!q(float64=7642)] map["in_bgp":%!q(bool=true) "in_whois":%!q(bool=false) "peer":%!q(float64=29117)]] "prefixes":[map["in_bgp":%!q(bool=false) "in_whois":%!q(bool=true) "irr_sources":["ARIN"] "prefix":"2604:e840::/32"] map["in_bgp":%!q(bool=false) "in_whois":%!q(bool=true) "irr_sources":["ALTDB"] "prefix":"2804:17d4::/32"]] "resource":"64496"]
}
