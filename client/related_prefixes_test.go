package client

import "fmt"

func ExampleClient_RelatedPrefixes() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "193/23" // prefix or IP range
	r, err := c.RelatedPrefixes(resource)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_time")
	fmt.Printf("%q", data)
	// Output:
	// map["prefixes":[map["asn_name":"RIPE-NCC-AS Reseaux IP Europeens Network Coordination Centre (RIPE NCC), EU" "origin_asn":"3333" "prefix":"193.0.0.0/21" "relationship":"Overlap - Less Specific"] map["asn_name":"HOSTWINDS - Hostwinds LLC., US" "origin_asn":"54290" "prefix":"192.255.128.0/17" "relationship":"Adjacency - Left"]] "resource":"193.0.0.0/23"]
}
