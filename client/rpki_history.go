package client

import "strings"

const (
	rpkiHistoryPath = "rpki-history/data.json"
)

// RPKIHistory GETs a timeseries of RPKI objects.
// Note: The data offered by this data call is provided on a best-effort basis. It covers the time period Jan 2011 to Dec 2018, and has know gaps in between.
// We're working to fill those gaps and provide data from Jan 2019 onwards. Stay tuned!
func (c *Client) RPKIHistory(prefix, prefixFamily, prefixMaskLength, maxLength, asn, trustAnchor, fields string) (map[string]interface{}, error) {
	m := make(map[string]string)
	parseRPKIHistoryPrefix(prefix, m)
	m["prefix_family"] = prefixFamily
	m["prefix_masklen"] = prefixMaskLength
	m["max_length"] = maxLength
	m["asn"] = asn
	m["trust_anchor"] = trustAnchor
	m["fields"] = fields
	return c.execute(rpkiHistoryPath, m)
}

func parseRPKIHistoryPrefix(prefix string, m map[string]string) {
	if strings.Contains(prefix, "=") {
		pairs := strings.Split(prefix, "=")
		m[pairs[0]] = pairs[1]
	} else {
		m["prefix"] = prefix
	}
}
