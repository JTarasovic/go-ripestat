package client

import "fmt"

func ExampleClient_AddressSpaceHierarchy() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "193/21" // prefix or IP range
	r, err := c.AddressSpaceHierarchy(resource)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_time")
	fmt.Printf("%q", data)
	// Output:
	// map["exact":[map["admin-c":"BRD-RIPE" "country":"NL" "created":"2003-03-17T12:15:57Z" "descr":"RIPE Network Coordination CentreAmsterdam, Netherlands" "inetnum":"193.0.0.0/21" "last-modified":"2017-12-04T14:42:31Z" "mnt-by":"RIPE-NCC-MNT" "netname":"RIPE-NCC" "org":"ORG-RIEN1-RIPE" "remarks":"Used for RIPE NCC infrastructure." "source":"RIPE" "status":"ASSIGNED PA" "tech-c":"OPS4-RIPE"]] "less_specific":[map["admin-c":"BRD-RIPE" "country":"NL" "created":"2012-03-09T15:03:38Z" "inetnum":"193.0.0.0-193.0.23.255" "last-modified":"2017-07-19T12:31:22Z" "mnt-by":"RIPE-NCC-HM-MNTRIPE-NCC-MNT" "mnt-lower":"RIPE-NCC-MNT" "mnt-routes":"RIPE-NCC-MNT" "netname":"NL-RIPENCC-OPS-990305" "org":"ORG-RIEN1-RIPE" "remarks":"Amsterdam, Netherlands" "source":"RIPE" "status":"ALLOCATED PA" "tech-c":"OPS4-RIPE"]] "more_specific":[] "parameters":map["resource":"193.0.0.0/21"] "resource":"193.0.0.0/21" "rir":"ripe"]
}
