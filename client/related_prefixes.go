package client

const (
	relatedPrefixesPath = "related-prefixes/data.json"
)

// RelatedPrefixes GETs prefixes that overlap or are adjacent to the specified IP resource.
func (c *Client) RelatedPrefixes(resource string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	return c.execute(relatedPrefixesPath, m)
}
