package client

import "fmt"

func ExampleClient_ASOverview() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // ASN
	r, err := c.ASOverview(resource)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_starttime")
	delete(data, "query_endtime")
	fmt.Printf("%q", data)
	// Output:
	// map["announced":%!q(bool=false) "block":map["desc":"For documentation and sample code; reserved by [RFC5398]" "name":"IANA Special-Purpose Autonomous System (AS) Numbers Registry" "resource":"64496-64511"] "holder":"-Reserved AS-" "resource":"64496" "type":"as"]
}
