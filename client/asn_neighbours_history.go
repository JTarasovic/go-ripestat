package client

const (
	asNNeighboursHistoryPath = "asn-neighbours-history/data.json"
)

// ASNNeighboursHistory GETs information about neigbouring ASNs of a queried ASN extended with history.
func (c *Client) ASNNeighboursHistory(resource, starttime, endtime, maxrows string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["starttime"] = starttime
	m["endtime"] = endtime
	m["max_rows"] = maxrows
	return c.execute(asNNeighboursHistoryPath, m)
}
