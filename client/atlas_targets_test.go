package client

import "fmt"

func ExampleClient_AtlasTargets() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "example.com" // prefix, ASN or hostname
	r, err := c.AtlasTargets(resource)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	fmt.Printf("%v", data["authenticated"])
	// Output:
	// false

}
