package client

func ExampleClient_CountryASNs() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "nl"    // The country has to be provided as an ISO-3166-1 alpha-2 country code
	queryTime := ""     // Defines the time of the lookup. This value needs to be or will be aligned to the RIS dump times!
	levelOfDetail := "" // Defines the level of detail in which the data is being returned. Levels are: 0 - Least detailed output 1 - Most detailed output
	r, err := c.CountryASNs(resource, queryTime, levelOfDetail)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	//countries
	// latest_time
	// lod
	// query_time
	// resource
}
