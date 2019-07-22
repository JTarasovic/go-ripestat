package client

import "fmt"

func ExampleClient_AbuseContactFinder() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "193/23"
	r, err := c.AbuseContactFinder(resource)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_time")
	fmt.Printf("%q", data)
	// Output:
	// map["anti_abuse_contacts":map["abuse_c":[map["description":"abuse-c" "email":"abuse@ripe.net" "key":"OPS4-RIPE"]] "emails":[] "extracted_emails":[] "objects_with_remarks":[]] "authorities":["ripe"] "blacklist_info":[] "global_network_info":map["description":"RIPE NCC (Status: ALLOCATED)" "name":"RIPE NCC (Status: ALLOCATED)" "source":"IANA IPv4 Address Space Registry" "source_url":"http://www.iana.org/assignments/ipv4-address-space/ipv4-address-space.csv"] "holder_info":map["name":"RIPE-NCC" "resource":"193.0.0.0 - 193.0.7.255"] "less_specifics":["193.0.0.0-193.0.23.255"] "more_specifics":[] "resource":"193.0.0.0/21" "special_resources":[]]
}
