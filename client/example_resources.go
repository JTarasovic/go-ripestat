package client

const (
	exampleResourcesPath = "example-resources/data.json"
)

// ExampleResources GETs sample resources for ASN, IPv4 and IPv6 resources.
// All are taken from routing data.
func (c *Client) ExampleResources() (map[string]interface{}, error) {
	return c.execute(exampleResourcesPath, nil)
}
