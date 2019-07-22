package client

func ExampleClient_Whois() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // ASN or IP Address
	r, err := c.Whois(resource)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// authorities
	// irr_records
	// query_time
	// records
	// resource
}
