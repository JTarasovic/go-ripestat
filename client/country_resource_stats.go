package client

const (
	countryResourceStatsPath = "country-resource-stats/data.json"
)

// CountryResourceStats GETs statistics on Internet resources for a country - this includes:
//     number of ASNs seen in routing data and registration data
//     number of prefixes in routing data and registration data (split into IPv4 and IPv6)
//     amount of IPv4 space seen in routing data as well as registration data
// The results can be restricted to a specific time period as well the granularity is variable but can be set explicitly.
func (c *Client) CountryResourceStats(resource, starttime, endtime, resolution string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["starttime"] = starttime
	m["endtime"] = endtime
	m["resolution"] = resolution
	return c.execute(countryResourceStatsPath, m)
}
