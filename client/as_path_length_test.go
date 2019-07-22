package client

import "fmt"

func ExampleClient_ASPathLength() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // ASN
	sortBy := "number"  // number(default), count, location, geo Sort by the given field. In the case of "geo", sort by approximating a world map on to a circle.
	r, err := c.ASPathLength(resource, sortBy)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_time")
	fmt.Printf("%q", data)
	// Output:
	// map["resource":"64496" "sort_by":"number" "stats":[]]
}
