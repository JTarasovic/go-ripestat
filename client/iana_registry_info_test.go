package client

import "fmt"

func ExampleClient_IANARegistryInfo() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "64496" // IP address, prefix or ASN
	bestMatchOnly := ""
	r, err := c.IANARegistryInfo(resource, bestMatchOnly)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_starttime")
	delete(data, "query_endtime")
	delete(data, "load_time")
	fmt.Printf("%q", data)
	// Output:
	// map["resource":"64496" "resources":[map["description":"For documentation and sample code; reserved by [RFC5398]" "details":map["Reason for Reservation":"For documentation and sample code; reserved by [RFC5398]" "Reference":"[RFC5398]"] "resource":"64496-64511" "source":"IANA Special-Purpose Autonomous System (AS) Numbers Registry" "source_url":"http://www.iana.org/assignments/iana-as-numbers-special-registry/special-purpose-as-numbers.csv" "type_properties":["asn" "range"]] map["description":"Reserved for use in documentation and sample code (Reference: RFC5398)" "details":map["Description":"Reserved for use in documentation and sample code" "RDAP":"" "Reference":"[RFC5398]" "Registration\n        Date":"2008-12-03" "WHOIS":""] "resource":"64496-64511" "source":"IANA 16-bit Autonomous System (AS) Numbers Registry" "source_url":"http://www.iana.org/assignments/as-numbers/as-numbers-1.csv" "type_properties":["asn" "range"]] map["description":"See Sub-registry 16-bit AS numbers (Reference: RFC1930)" "details":map["Description":"See Sub-registry 16-bit AS numbers" "RDAP":"" "Reference":"[RFC1930]" "Registration\n        Date":"" "WHOIS":""] "resource":"0-65535" "source":"IANA 32-bit Autonomous System (AS) Numbers Registry" "source_url":"http://www.iana.org/assignments/as-numbers/as-numbers-2.csv" "type_properties":["asn" "range"]]] "returned":%!q(float64=3)]
}
