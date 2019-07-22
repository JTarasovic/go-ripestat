package client

const (
	whoisObjectLastUpdatedPath = "whois-object-last-updated/data.json"
)

// WhoisObjectLastUpdated GETs information of when a certain object was last updated in the whois database.
func (c *Client) WhoisObjectLastUpdated(object, objectType, source, timestamp, compareWithLive string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["object"] = object
	m["type"] = objectType
	m["source"] = source
	m["timestamp"] = timestamp
	m["compare_with_live"] = compareWithLive
	return c.execute(whoisObjectLastUpdatedPath, m)
}
