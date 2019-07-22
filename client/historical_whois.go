package client

const (
	historicalWhoisPath = "historical-whois/data.json"
)

// HistoricalWhois GETs information on objects that are stored in the RIPE DB.
// The result is aligned to a specific object, which is identified by an object type and an object key, which is similar to the Whois data call.
// In contrast to the Whois data call, this data call puts a focus on historical changes of an object and its related objects.
// Historical changes are given in the form of versions, one version - by default the latest version - is presented with details.
// Related objects are separated into referencing and referenced objects.
// Referencing objects are objects that have a reference to the object in focus and referenced objects are referenced from the object in focus.
func (c *Client) HistoricalWhois(resource, version string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["version"] = resource
	return c.execute(historicalWhoisPath, m)
}
