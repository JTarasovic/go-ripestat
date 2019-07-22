package client

import "fmt"

func ExampleClient_MlabActivityCount() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "at" // IPv4 Prefix, IPv4 address or 2-digit ISO-3166 country code
	starttime := ""
	endtime := ""
	r, err := c.MlabActivityCount(resource, starttime, endtime)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_starttime")
	delete(data, "query_endtime")
	fmt.Printf("%q", data)
	// Output:
	// map["nr_ips":%!q(float64=1487) "perc_coverage":<nil> "resource":"at"]
}
