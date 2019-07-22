package client

func ExampleClient_ExampleResources() {
	c := NewClient().WithSourceApp("go-ripestat")
	r, err := c.ExampleResources()
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// asn
	// ipv4
	// ipv6
	// range4
}
