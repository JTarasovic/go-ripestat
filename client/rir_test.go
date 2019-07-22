package client

import "fmt"

func ExampleClient_RIR() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // ASN
	starttime := ""
	endtime := ""
	levelOfDetail := "" // 0 - Least detailed output 1 - Default output 2 - Most detailed output
	r, err := c.RIR(resource, starttime, endtime, levelOfDetail)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_starttime")
	delete(data, "query_endtime")
	fmt.Printf("%q", data)
	// Output:
	// map["latest":"2019-07-21T00:00:00" "lod":%!q(float64=1) "resource":"64496" "rirs":[]]
}
