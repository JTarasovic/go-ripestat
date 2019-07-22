package client

import "fmt"

func ExampleClient_AddressSpaceUsage() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "193/23" // prefix or IP range
	r, err := c.AddressSpaceUsage(resource)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_time")
	fmt.Printf("%q", data)
	// Output:
	// map["allocations":[map["allocation":"193.0.0.0-193.0.23.255" "asn_name":"NL-RIPENCC-OPS-990305" "assignments":%!q(float64=1) "status":"ALLOCATED PA"]] "assignments":[map["address_range":"193.0.0.0/21" "asn_name":"RIPE-NCC" "parent_allocation":"193.0.0.0-193.0.23.255" "status":"ASSIGNED PA"]] "ip_stats":[map["ips":%!q(float64=512) "status":"ASSIGNED PA"]] "resource":"193.0.0.0/23"]
}
