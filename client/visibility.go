package client

const (
	visibilityPath = "visibility/data.json"
)

// Visibility GETs information on the visibility of a resource as observed from RIS (http://ris.ripe.net).
// Historical lookups are supported - a query has to be aligned to the times (00:00, 08:00 and 16:00 UTC) when RIS data has been collected.
// By default, the data call returns the latest data.
func (c *Client) Visibility(resource, queryTime, include string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["query_time"] = queryTime
	m["include"] = include
	return c.execute(visibilityPath, m)
}
