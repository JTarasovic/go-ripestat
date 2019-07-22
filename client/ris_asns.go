package client

const (
	risASNsPath = "ris-asns/data.json"
)

// RISAsns GETs high-level information on ASNs in RIS, including:
//     - total number of ASNs
//     - listing of all ASNs
// The data call supports history, with each data point being aligned to times a dump is created in RIS (00:00, 08:00 and 16:00 UTC).
// By default, the data call returns the total number of ASNs; more details can be obtained using parameters.
// Note the term "transit" related to this data call means any ASN that is seen in the AS paths, collected by RIS, that is not the origin of a route.
func (c *Client) RISAsns(queryTime, listASNs, asnTypes string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["query_time"] = queryTime
	m["list_asns"] = listASNs
	m["asn_types"] = asnTypes
	return c.execute(risASNsPath, m)
}
