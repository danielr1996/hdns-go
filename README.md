# hdns-go

[![Go Reference](https://pkg.go.dev/badge/github.com/danielr1996/hdns-go.svg)](https://pkg.go.dev/github.com/danielr1996/hdns-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/danielr1996/hdns-go)](https://goreportcard.com/report/github.com/danielr1996/hdns-go)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/danielr1996/hdns-go?style=plastic)

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

## Developing
> `hdns-go` uses [Task](https://taskfile.dev/#/installation) to run the tests.
>
> Install Task or execute the commands listed in `Taskfile.yml` manually.

To run the tests run `task test`, to view the coverage results run `task test:view`
