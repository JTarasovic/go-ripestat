package client

const (
	addressSpaceUsagePath = "address-space-usage/data.json"
)

// AddressSpaceUsage GETs the usage ratio of this prefix or IP range according to the objects currently present in the RIPE Whois database.
// The data is represented as assignments, allocations and statistic on the distribution of the IPs covered by the queried resource.
func (c *Client) AddressSpaceUsage(resource string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	return c.execute(addressSpaceUsagePath, m)
}
