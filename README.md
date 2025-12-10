# globalrouter-go
Go SDK for Selectel Global Router resources

## Getting started

### Installation

You can install needed `globalrouter-go` packages via `go get` command:

```bash
go get github.com/selectel/globalrouter-go
```

### Authentication

To work with the Selectel global-router Servers API you first need to:

* Create a Selectel account: [registration page](https://my.selectel.ru/registration).
* Create a project in Selectel Cloud Platform [projects](https://my.selectel.ru/vpc/projects).
* Retrieve a token for your account via API or [go-selvpcclient](https://github.com/selectel/go-selvpcclient).

### Endpoints

You can find available endpoints [here](https://docs.selectel.ru/en/api/urls/).

### Usage example

```go
package main

import (
	"context"
	"fmt"

	globalrouter "github.com/selectel/globalrouter-go/pkg/v1"
)

func main() {
	// Auth token.
	token := "gAAAAABeVNzu-..."

	// Global router endpoint to work with.
	endpoint := "https://api.selectel.ru/naas/v1"

	// Create the client.
	client := globalrouter.NewClientV1(
		token,	
		globalrouter.WithAPIUrl("https://api.selectel.ru/naas/v1"),
	)

	// Get all Zones with specified name
	zones, _, _ := client.ListZones(ctx, nil)

	// Print all cloud zones
	for _, zone := range *zones {
		if zone.Service == "vpc" {
			fmt.Printf("Zone name: %s\n", zone.Name)
		}
	}
}
```
