package client

func ExampleClient_RIRGeo() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // IP resource or ASN
	queryTime := ""
	r, err := c.RIRGeo(resource, queryTime)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// earliest_time
	// latest_time
	// located_resources
	// parameters
	// result_time
}
