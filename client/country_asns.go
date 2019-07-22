package client

const (
	countryASNsPath = "country-asns/data.json"
)

// CountryASNs GETs information on a country's registered and routed ASNs.
// Registered ASNs are based on registration information made public by the Regional Internet Registries.
// The routing information is based on the data collected with the RIPE NCC's RIS system, https://ris.ripe.net.
// The data call supports history, with data points being aligned to times dumps are created in RIS (00:00, 08:00 and 16:00 UTC).
// By default, the data call returns just the number of registered and routed ASNs.
// This is mainly to prevent returning thousands of ASNs.
func (c *Client) CountryASNs(resource, queryTime, levelOfDetail string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["query_time"] = queryTime
	m["lod"] = levelOfDetail
	return c.execute(countryASNsPath, m)
}
