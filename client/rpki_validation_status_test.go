package client

import "fmt"

func ExampleClient_RPKIValidationStatus() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "3333" // ASN
	prefix := "193/21"
	r, err := c.RPKIValidationStatus(resource, prefix)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_starttime")
	delete(data, "query_endtime")
	fmt.Printf("%q", data)
	// Output:
	// map["prefix":"193.0.0.0/21" "resource":"3333" "status":"valid" "validating_roas":[map["max_length":%!q(float64=21) "origin":"AS3333" "prefix":"193.0.0.0/21" "source":"RIPE NCC RPKI Root" "validity":"valid"]]]
}
