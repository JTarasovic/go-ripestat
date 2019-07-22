package client

const (
	ianaRegistryInfoPath = "iana-registry-info/data.json"
)

// IANARegistryInfo GETs information from various data sources maintained by IANA. These include:
//     IANA 16-bit Autonomous System (AS) Numbers Registry (http://www.iana.org/assignments/as-numbers/as-numbers-1.csv)
//     IANA 32-bit Autonomous System (AS) Numbers Registry (http://www.iana.org/assignments/as-numbers/as-numbers-2.csv)
//     Etc. - the detailed list of data sources included can be seen in the methodology information of the data call. Note: Output format for the methodology information is JSON!
// The data call supports a "resource" parameter which filters all results down to entries that are topologically related to the given resource.
// The data is refreshed once a day to guarantee up-to-date information.
func (c *Client) IANARegistryInfo(resource, bestMatchOnly string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["best_match_only"] = bestMatchOnly
	return c.execute(ianaRegistryInfoPath, m)
}
