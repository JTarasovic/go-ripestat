package client

const (
	risPrefixesPath = "ris-prefixes/data.json"
)

// RISPrefixes GETs information on prefixes related to an ASN.
// The data call distinguishes prefixes in the originated and transitted ASN.
// Note that this distinction is purely based on the perspective of the RIPE NCC's RIS system and does NOT imply the underlying (business) relationships between networks!
// The data call supports history, with each data point being aligned to times a dump is created in RIS (00:00, 08:00 and 16:00 UTC).
// By default, the data call returns just the count of prefixes related to the looked up ASN.
// This is mainly to prevent returning thousands of prefixes. See parameter settings below to further tailor the output to your needs.
func (c *Client) RISPrefixes(resource, queryTime, listPrefixes, types, af, noise string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["query_time"] = queryTime
	m["list_prefixes"] = listPrefixes
	m["types"] = types
	m["af"] = af
	m["noise"] = noise
	return c.execute(risPrefixesPath, m)
}
