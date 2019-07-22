package client

const (
	risFullTableThresholdPath = "ris-full-table-threshold/data.json"
)

// RISFullTableThreshold GETs the cut-off threshold for the number of prefixes that a BGP full-table peer requires to have.
// Peers to RIS (http://ris.ripe.net) that share less than this amount of prefixes are not considered full-table peers and hence are not considered in calculations like routing visibility.
// The threshold is obviously different between address families (IPv4 and IPv6) and time.
// For this reason the data call also supports historical lookups.
func (c *Client) RISFullTableThreshold(queryTime string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["query_time"] = queryTime
	return c.execute(risFullTableThresholdPath, m)
}
