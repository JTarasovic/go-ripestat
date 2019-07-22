package client

import "fmt"

func ExampleClient_RIRGeo() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // IP resource or ASN
	queryTime := ""
	r, err := c.RIRGeo(resource, queryTime)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_time")
	fmt.Printf("%q", data)
	// Output:
	// map["earliest_time":"2005-02-18T00:00:00" "latest_time":"2019-07-21T00:00:00" "located_resources":[] "parameters":map["query_time":"2019-07-21T00:00:00" "resource":"64496"] "result_time":""]
}
