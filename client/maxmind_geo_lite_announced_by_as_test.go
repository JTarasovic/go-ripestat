package client

import "fmt"

func ExampleClient_MaxmindGeoLiteAnnouncedByAS() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // ASN
	r, err := c.MaxmindGeoLiteAnnouncedByAS(resource)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_time")
	fmt.Printf("%q", data)
	// Output:
	// map["earliest_time":"2018-03-27T00:00:00" "latest_time":"2019-07-16T00:00:00" "located_resources":[] "parameters":map["query_time":"2019-07-16T00:00:00" "resolution":"raw" "resource":"64496"] "result_time":"2019-07-16T00:00:00" "unknown_percentage":map["v4":%!q(float64=100) "v6":%!q(float64=100)]]
}
