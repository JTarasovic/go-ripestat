package client

const (
	countryResourceListPath = "country-resource-list/data.json"
)

// CountryResourceList GETs the Internet resources associated with a country, including ASNs, IPv4 ranges and IPv4/6 CIDR prefixes.
// This data is derived from the RIR Statistics files maintained by the various RIRs.
func (c *Client) CountryResourceList(resource, time, v4Format string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["time"] = time
	m["v4_format"] = v4Format
	return c.execute(countryResourceListPath, m)
}
