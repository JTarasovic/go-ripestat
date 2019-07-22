package client

import "fmt"

func ExampleClient_MaxmindGeoLite() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "193/23" // prefix or IP address
	queryTime := ""
	r, err := c.MaxmindGeoLite(resource, queryTime)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_time")
	fmt.Printf("%q", data)
	// Output:
	// map["earliest_time":"2018-03-27T00:00:00" "latest_time":"2019-07-16T00:00:00" "located_resources":[map["locations":[map["city":"" "country":"NL" "covered_percentage":%!q(float64=100) "latitude":%!q(float64=52.3824) "longitude":%!q(float64=4.8995) "resources":["193.0.0.0/21"]]] "resource":"193.0.0.0/23" "unknown_percentage":%!q(float64=0)]] "parameters":map["query_time":"2019-07-16T00:00:00" "resolution":"raw" "resource":"193.0.0.0/23"] "result_time":"2019-07-16T00:00:00" "unknown_percentage":map["v4":%!q(float64=0)]]
}
