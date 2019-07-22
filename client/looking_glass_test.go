package client

import "fmt"

func ExampleClient_LookingGlass() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "192/23" // A prefix or an IP address.
	r, err := c.LookingGlass(resource)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_time")
	delete(data, "latest_time")
	fmt.Printf("%q", data)
	// Output:
	// map["parameters":map["resource":"192.0.0.0/23"] "rrcs":[]]
}
