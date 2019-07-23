package client

func ExampleClient_RISFirstLastSeen() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // prefix or ASN
	// This parameter defines additional data to be included.
	// "more_specific" includes more specific IP ranges, which only works for prefix lookups. By default "more_specific" is not set as it makes the lookup slower.
	// "low_visibility_flag" includes the flag to indicate low visibility. By default it is not included.
	include := ""
	r, err := c.RISFirstLastSeen(resource, include)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// include
	// latest_time
	// resource
	// resources
	// stats
}
