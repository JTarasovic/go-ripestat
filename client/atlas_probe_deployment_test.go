package client

func ExampleClient_AtlasProbeDeployment() {
	c := NewClient().WithSourceApp("go-ripestat")
	// Due to the ambigious nature of abbreviated identifiers for regions and countries (e.g. me for Middle East and Montenegro) region and country resources should be prefixes with "region_" or "cc_".
	// Looking up a network can be specified on the IP version by using the prefix "asn4_" for IP v4 networks and "asn6_" for IP v6 networks.
	// For mixed results the resources just need to be comma separated.
	resource := "64496" // region, country, network (ASN) or mixed
	starttime := ""     // ISO8601 or Unix timestamp for query start
	endtime := ""       // ISO8601 or Unix timestamp for query start
	r, err := c.AtlasProbeDeployment(resource, starttime, endtime)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	printMapKeysSorted(data)
	// Output:
	// deployments
	// endtime
	// merge
	// query_date
	// resource
	// starttime
}
