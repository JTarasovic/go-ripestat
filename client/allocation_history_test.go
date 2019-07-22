package client

import (
	"fmt"
)

func ExampleClient_AllocationHistory() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496"             // ASN
	starttime := "2011-10-10T12:00" // ISO8601 or Unix timestamp for query start
	endtime := ""                   // ISO8601 or Unix timestamp for query start
	r, err := c.AllocationHistory(resource, starttime, endtime)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_starttime")
	delete(data, "query_endtime")
	fmt.Printf("%q", data)
	// Output:
	// map["earliest_time":"2000-08-01T00:00:00" "latest_time":"2019-07-22T08:00:00" "prefixes":[map["prefix":"2001:c78:ff11::/48" "timelines":[map["endtime":"2011-12-31T16:00:00" "starttime":"2011-10-10T16:00:00"] map["endtime":"2012-04-02T16:00:00" "starttime":"2012-02-01T00:00:00"] map["endtime":"2012-08-17T00:00:00" "starttime":"2012-04-11T08:00:00"]]] map["prefix":"2001:c78:ff12::/47" "timelines":[map["endtime":"2011-12-31T16:00:00" "starttime":"2011-10-10T16:00:00"] map["endtime":"2012-08-17T00:00:00" "starttime":"2012-02-01T00:00:00"]]] map["prefix":"198.57.8.0/24" "timelines":[map["endtime":"2013-03-27T00:00:00" "starttime":"2013-01-22T00:00:00"] map["endtime":"2013-05-02T16:00:00" "starttime":"2013-03-27T16:00:00"] map["endtime":"2013-07-31T00:00:00" "starttime":"2013-05-03T16:00:00"] map["endtime":"2013-08-04T00:00:00" "starttime":"2013-07-31T16:00:00"] map["endtime":"2013-08-10T16:00:00" "starttime":"2013-08-04T16:00:00"] map["endtime":"2013-09-20T00:00:00" "starttime":"2013-08-11T08:00:00"] map["endtime":"2013-10-03T00:00:00" "starttime":"2013-09-20T16:00:00"] map["endtime":"2013-10-17T00:00:00" "starttime":"2013-10-03T16:00:00"] map["endtime":"2014-01-31T16:00:00" "starttime":"2013-10-17T16:00:00"] map["endtime":"2014-03-12T16:00:00" "starttime":"2014-03-01T00:00:00"] map["endtime":"2014-05-07T08:00:00" "starttime":"2014-03-13T08:00:00"] map["endtime":"2014-06-04T16:00:00" "starttime":"2014-05-08T00:00:00"] map["endtime":"2014-10-21T00:00:00" "starttime":"2014-06-06T00:00:00"]]]] "resource":"64496"]
}
