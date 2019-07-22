package client

func ExampleClient_ReverseDNS() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "193/23" // prefix
	r, err := c.ReverseDNS(resource)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// delegations
	// query_time
	// resource
}
