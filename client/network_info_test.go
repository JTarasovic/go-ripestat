package client

import "fmt"

func ExampleClient_NetworkInfo() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "140.78.90.50" // Any IP address one wants to get network info for
	r, err := c.NetworkInfo(resource)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_starttime")
	delete(data, "query_endtime")
	fmt.Printf("%q", data)
	// Output:
	// map["asns":["1205"] "prefix":"140.78.0.0/16"]
}
