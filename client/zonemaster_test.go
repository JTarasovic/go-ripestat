package client

func ExampleClient_Zonemaster() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "0.0.193.in-addr.arpa" // The hostname the DNS checks are based on or the ID of the test result
	method := ""                       // This is necessary when retrieving test details and when the "resource" parameter is a test ID
	r, err := c.Zonemaster(resource, method)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// parameters
	// result
}
