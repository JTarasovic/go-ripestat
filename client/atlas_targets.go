package client

const (
	atlasTargetsPath = "atlas-targets/data.json"
)

// AtlasTargets GETs information on the RIPE Atlas measurements that target an network (ASN), a prefix or a hostname.
// The information is based on data coming from the RIPE Atlas REST API, https://atlas.ripe.net/docs/api/v2/manual/.
func (c *Client) AtlasTargets(resource string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	return c.execute(atlasTargetsPath, m)
}
