# hdns-go
A Go library for the Hetzner DNS API

> Inspired by [hcloud-go](https://github.com/hetznercloud/hcloud-go)

## Usage
Get your API Token at [https://dns.hetzner.com/settings/api-token]()

```go
package main

import (
	"fmt"
	"github.com/danielr1996/hdns-go/src/hdns"
)

func main(){
	client := hdns.NewClient().WithToken("<api token>")
	zones := client.Zones()
	fmt.Printf("%",zones)
}
```
