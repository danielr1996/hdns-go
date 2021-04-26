# hdns-go

[![Go Reference](https://pkg.go.dev/badge/github.com/danielr1996/hdns-go.svg)](https://pkg.go.dev/github.com/danielr1996/hdns-go)

A Go library for the [Hetzner DNS API](https://dns.hetzner.com/api-docs/)

> Inspired by [hcloud-go](https://github.com/hetznercloud/hcloud-go)

## Usage

> For the full documentation see https://pkg.go.dev/github.com/danielr1996/hdns-go

Get your API Token at [https://dns.hetzner.com/settings/api-token](https://dns.hetzner.com/settings/api-token)

Create a new Client with your token

```go
package main

import (
	"github.com/danielr1996/hdns-go/client"
)

func main() {
	client := client.New().WithToken("<api token>")
}
```

Receive all zones from the Hetzner DNS API associated with the token

```go
// ...
res, err := c.Zone.GetAll()
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(res)
// ...
```

