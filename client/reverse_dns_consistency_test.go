package client

import "fmt"

func ExampleClient_ReverseDNSConsistency() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // prefix or ASN
	ipv4 := ""
	ipv6 := ""
	r, err := c.ReverseDNSConsistency(resource, ipv4, ipv6)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_time")
	fmt.Printf("%q", data)
	// Output:
	// map["ipv4":%!q(bool=true) "ipv6":%!q(bool=true) "prefixes":map["ipv4":map[] "ipv6":map[]] "resource":"64496" "source":"routes"]
}
