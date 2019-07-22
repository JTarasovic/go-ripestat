package client

import "fmt"

func ExampleClient_ReverseDNSIP() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "193.0.6.139" // IP address
	r, err := c.ReverseDNSIP(resource)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_time")
	fmt.Printf("%q", data)
	// Output:
	// map["error":"" "resource":"193.0.6.139" "result":["www.ripe.net"]]
}
