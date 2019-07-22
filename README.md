# go-ripestat
A client for the RIPEstat DATA API written in Go.

[API Documentation](https://stat.ripe.net/docs/data_api)
[Client Documentation](https://godoc.org/github.com/JTarasovic/go-ripestat/client)


## Installation
```bash
# Go Modules
require github.com/jtarasovic/go-ripestat/client
```


## Usage

```go
import "github.com/jtarasovic/go-ripestat/client"
```

There are examples documented for each call. You can either view it [online](https://godoc.org/github.com/JTarasovic/go-ripestat/client) or in the `client/` source directory.



### Basic Example
```go
package main

import (
    "fmt"
    "github.com/jtarasovic/go-ripestat/client"
)

func main() {
	c := NewClient()
	r, err := c.ExampleResources()
	if err != nil {
        // handle error sanely
		panic(err)
	}
    // result is map[string]interface{} currently
    // see API documentation for returned data
    fmt.Println(r["data"])
}
```


### Using Client Options

```go
package main

import (
    "fmt"
    "github.com/jtarasovic/go-ripestat/client"
)

func main() {
	c := NewClient()
        // helps RIPE identify callers. please include something unique for you application
        .WithSourceApp("my-app")


	resource := "193/23"                    // prefix or IP range
	starttime := "2019-01-01T00:00:00"      // ISO8601 or Unix timestamp
	endtime := "2019-07-01T00:00:00"        // ISO8601 or Unix timestamp

	r, err := c.Blacklist(resource, starttime, endtime)
	if err != nil {
        // handle error sanely
		panic(err)
	}
    // result is map[string]interface{} currently
    // see API documentation for returned data
    fmt.Println(r["data"])
}
```
