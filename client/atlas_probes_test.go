package client

func ExampleClient_AtlasProbes() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "AT" // prefix, ASN or country
	r, err := c.AtlasProbes(resource)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// probes
	// resource
	// stats
}
