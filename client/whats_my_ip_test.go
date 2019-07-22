package client

func ExampleClient_WhatsMyIP() {
	c := NewClient().WithSourceApp("go-ripestat")
	r, err := c.WhatsMyIP()
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// ip

}
