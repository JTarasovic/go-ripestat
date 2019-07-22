package client

const (
	searchcompletePath = "searchcomplete/data.json"
)

// SearchComplete GETs example resource that are directly or indirectly related to the given input.
func (c *Client) SearchComplete(resource, limit string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["limit"] = limit
	return c.execute(searchcompletePath, m)
}
