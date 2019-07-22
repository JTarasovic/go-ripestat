package client

const (
	allocationHistoryPath = "allocation-history/data.json"
)

// AllocationHistory GETs information supplied by IANA and RIRs for allocations and direct assignments of prefixes and AS numbers of time.
func (c *Client) AllocationHistory(resource, starttime, endtime string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["starttime"] = starttime
	m["endtime"] = endtime
	return c.execute(announcedPrefixesPath, m)
}
