package client

import "fmt"

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
	delete(data, "query_starttime")
	delete(data, "query_endtime")
	fmt.Printf("%q", data)
	// Output:
	// map["include":[] "latest_time":"2019-07-22T08:00:00" "resource":["64496"] "resources":[map["first":map["time":"2018-05-17T16:00:00"] "last":map["time":"2019-07-22T08:00:00"] "resource":"64496" "type":"t"] map["first":map["time":"2010-09-23T00:00:00"] "last":map["time":"2019-03-21T00:00:00"] "resource":"64496" "type":"o"]] "stats":map["count":%!q(float64=2)]]
}
