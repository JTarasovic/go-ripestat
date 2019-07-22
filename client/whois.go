package client

const (
	whoisPath = "whois/data.json"
)

// Whois GETs whois information from the relevant Regional Internet Registry and Routing Registry.
func (c *Client) Whois(resource string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	return c.execute(whoisPath, m)
}
