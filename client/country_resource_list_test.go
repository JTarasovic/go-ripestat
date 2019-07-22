package client

func ExampleClient_CountryResourceList() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "at" // 2-digit ISO-3166 country code
	time := ""
	v4Format := ""
	r, err := c.CountryResourceList(resource, time, v4Format)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data["resources"].(map[string]interface{}))
	// Output:
	// asn
	// ipv4
	// ipv6
}
