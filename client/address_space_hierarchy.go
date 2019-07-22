package client

const (
	addressSpaceHierarchyPath = "address-space-hierarchy/data.json"
)

// AddressSpaceHierarchy GETs address space objects (inetnum, inet6num...) from the RIPE DB (whois) related (exact, more- and less-specific) to the queried resource.
// Less- and more-specific results are first-level only, further levels would have to be retrieved iteratively
// Due to performance issues on the backend side, data calls are limited to a maximum prefix size of /16 for IPv4 resources!
func (c *Client) AddressSpaceHierarchy(resource string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	return c.execute(addressSpaceHierarchyPath, m)
}
