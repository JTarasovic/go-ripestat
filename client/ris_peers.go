package client

const (
	risPeersPath = "ris-peers/data.json"
)

// RISPeers GETs information on the peers of RIS - ASN, IP address and number of shared routes.
// The data is grouped by RIS collectors.
// Historical lookups are supported - a query has to be aligned to the times (00:00, 08:00 and 16:00 UTC) when RIS data has been collected.
// By default, the data call returns the latest data.
func (c *Client) RISPeers(queryTime string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["query_time"] = queryTime
	return c.execute(risPeersPath, m)
}
