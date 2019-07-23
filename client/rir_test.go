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
	delete(data, "latest")
	fmt.Printf("%q", data)
	// Output:
	// map["lod":%!q(float64=1) "resource":"64496" "rirs":[]]
}
