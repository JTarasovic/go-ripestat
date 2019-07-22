package client

const (
	rpkiValidationStatusPath = "rpki-validation/data.json"
)

// RPKIValidationStatus GETs the RPKI validity state for a combination of prefix and Autonomous System.
// This combination will be used to perform the lookup against the RIPE NCC's RPKI Validator, and then return its RPKI validity state.
func (c *Client) RPKIValidationStatus(resource, prefix string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["prefix"] = prefix
	return c.execute(rpkiValidationStatusPath, m)
}
