package client

import "fmt"

func ExampleClient_HistoricalWhois() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // This is a prefix (v4/v6), an AS number, or a string of the format "object-type:object-key" for looking up generic database objects.
	version := ""       // Given as a numerical value, the value must match exactly the historical version number. Given as a time-based value, the version that was valid at the given time will be returned.
	r, err := c.HistoricalWhois(resource, version)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_starttime")
	delete(data, "query_endtime")
	delete(data, "latest_time")
	fmt.Printf("%q", data)
	// Output:
	// map["access":"standard" "database":"RIPE" "num_versions":%!q(float64=0) "objects":[] "referenced_by":map[] "referencing":[] "resource":"64496" "suggestions":[] "terms_and_conditions":"" "version":%!q(float64=64496) "versions":[]]
}
