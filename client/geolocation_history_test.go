package client

import "fmt"

func ExampleClient_GeolocationHistory() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "example.com" // prefix, IP range, ASN or hostname
	starttime := "1514764800"
	endtime := ""
	r, err := c.GeolocationHistory(resource, starttime, endtime)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_starttime")
	delete(data, "query_endtime")
	fmt.Printf("%q", data)
	// Output:
	// map["months":[map["distributions":[map["city":"Norwell" "country":"US" "percentage":%!q(float64=100)]] "month":"2018-01-01T00:00:00"] map["distributions":[map["city":"Norwell" "country":"US" "percentage":%!q(float64=100)]] "month":"2018-02-01T00:00:00"] map["distributions":[map["city":"Norwell" "country":"US" "percentage":%!q(float64=100)]] "month":"2018-03-01T00:00:00"]] "resource":"example.com"]
}
