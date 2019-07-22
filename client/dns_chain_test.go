package client

import "fmt"

func ExampleClient_DNSChain() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "example.com" // Hostname or IP address (Ipv4 or IPv6)
	r, err := c.DNSChain(resource)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_time")
	fmt.Printf("%q", data)
	// Output:
	// map["authoritative_nameservers":["a.iana-servers.net" "b.iana-servers.net"] "forward_nodes":map["example.com":["93.184.216.34" "2606:2800:220:1:248:1893:25c8:1946"]] "nameservers":["193.0.19.101" "193.0.19.102" "2001:67c:2e8:11::c100:1365" "2001:67c:2e8:11::c100:1366" "2001:67c:2e8:11::c100:1367" "193.0.19.103"] "resource":"example.com" "reverse_nodes":map["2606:2800:220:1:248:1893:25c8:1946":[] "93.184.216.34":[]]]
}
