package client

func ExampleClient_MlabClients() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "193.0.0.0/16" // IPv4 Prefix, IPv4 address or 2-digit ISO-3166 country code
	starttime := "2013-08-21T07:00"
	endtime := "2013-18-22T12:00"
	r, err := c.MlabClients(resource, starttime, endtime)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// clients
	// nr_clients
	// perc_coverage
	// query_endtime
	// query_starttime
	// resource
}
