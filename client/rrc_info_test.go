package client

func ExampleClient_RRCInfo() {
	c := NewClient().WithSourceApp("go-ripestat")
	r, err := c.RRCInfo()
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// parameters
	// rrcs
}
