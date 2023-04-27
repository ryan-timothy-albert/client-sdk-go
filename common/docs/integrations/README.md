# Integrations

## Overview

View and manage your available integrations in Codat.

### Available Operations

* [Get](#get) - Get integration
* [GetBranding](#getbranding) - Get branding
* [List](#list) - List integrations

## Get

Get single integration, by platformKey

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
	"github.com/codatio/client-sdk-go/common/pkg/models/operations"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetIntegrationRequest{
        PlatformKey: "gbol",
    }

    res, err := s.Integrations.Get(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Integration != nil {
        // handle response
    }
}
```

## GetBranding

Get branding for platform.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
	"github.com/codatio/client-sdk-go/common/pkg/models/operations"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetIntegrationsBrandingRequest{
        PlatformKey: "gbol",
    }

    res, err := s.Integrations.GetBranding(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Branding != nil {
        // handle response
    }
}
```

## List

List your available integrations

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
	"github.com/codatio/client-sdk-go/common/pkg/models/operations"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListIntegrationsRequest{
        OrderBy: codatcommon.String("-modifiedDate"),
        Page: 1,
        PageSize: codatcommon.Int(100),
        Query: codatcommon.String("veritatis"),
    }

    res, err := s.Integrations.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Integrations != nil {
        // handle response
    }
}
```