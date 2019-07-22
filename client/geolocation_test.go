package client

import "fmt"

func ExampleClient_Geolocation() {
	c := NewClient().WithSourceApp("go-ripestat")
	resource := "example.com" // prefix, IP range, ASN or hostname
	timestamp := ""
	r, err := c.Geolocation(resource, timestamp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q", r)
	// Output:
	// map[]

}
