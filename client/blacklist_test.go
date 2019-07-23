package client

import "fmt"

func ExampleClient_Blacklist() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "193/23" // prefix or IP range
	starttime := ""      // ISO8601 or Unix timestamp
	endtime := ""        // ISO8601 or Unix timestamp
	r, err := c.Blacklist(resource, starttime, endtime)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_starttime")
	delete(data, "query_endtime")
	fmt.Printf("%q", data)

	// Output:
	// map["resource":"193.0.0.0/23" "sources":map["uceprotect-level1":[map["details":"no details" "prefix":"193.0.0.1/32" "timelines":[map["endtime":"2016-09-06T08:10:00" "starttime":"2016-08-31T00:10:00"] map["endtime":"2016-12-01T16:10:00" "starttime":"2016-11-25T00:10:00"] map["endtime":"2016-12-13T00:10:00" "starttime":"2016-12-06T08:10:00"] map["endtime":"2017-03-25T16:10:00" "starttime":"2017-03-20T00:10:00"]]] map["details":"no details" "prefix":"193.0.1.102/32" "timelines":[map["endtime":"2018-10-28T08:10:00" "starttime":"2018-10-21T16:10:00"] map["endtime":"2018-11-15T00:10:00" "starttime":"2018-11-08T08:10:00"] map["endtime":"2019-04-04T08:10:00" "starttime":"2019-03-28T16:10:00"] map["endtime":"2019-04-18T08:10:00" "starttime":"2019-04-11T16:10:00"] map["endtime":"2019-06-18T00:10:00" "starttime":"2019-06-11T08:10:00"]]] map["details":"no details" "prefix":"193.0.1.101/32" "timelines":[map["endtime":"2018-10-19T16:10:00" "starttime":"2018-10-13T00:10:00"]]] map["details":"no details" "prefix":"193.0.0.170/32" "timelines":[map["endtime":"2019-01-11T00:10:00" "starttime":"2018-12-04T00:10:00"] map["endtime":"2019-01-26T08:10:00" "starttime":"2019-01-12T08:10:00"] map["endtime":"2019-02-12T16:12:00" "starttime":"2019-01-28T00:10:00"]]] map["details":"no details" "prefix":"193.0.1.241/32" "timelines":[map["endtime":"2018-09-15T08:10:00" "starttime":"2018-09-08T16:10:00"] map["endtime":"2019-01-18T00:10:00" "starttime":"2019-01-11T08:10:00"]]]]]]
}
