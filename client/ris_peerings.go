package client

const (
	risPeeringsPath = "ris-peerings/data.json"
)

// RISPeerings GETs routes for advertisements of a given IP resource, or that are originated from a given ASN, as seen by the RIPE NCC route collectors.
// Historical lookups are supported - a query has to be aligned to the times (00:00, 08:00 and 16:00 UTC) when RIS data has been collected.
// By default, the data call returns the latest data.
func (c *Client) RISPeerings(resource, queryTime string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["query_time"] = queryTime
	return c.execute(risPeeringsPath, m)
}
