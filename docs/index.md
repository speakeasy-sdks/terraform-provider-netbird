---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "repo Provider"
subcategory: ""
description: |-
  NetBird REST API: API to manipulate groups, rules, policies and retrieve information about peers and users
---

# repo Provider

NetBird REST API: API to manipulate groups, rules, policies and retrieve information about peers and users

## Example Usage

```terraform
terraform {
  required_providers {
    netbird = {
      source  = "NetBird/netbird"
      version = "2.4.0"
    }
  }
}

provider "netbird" {
  # Configuration options
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `bearer_auth` (String, Sensitive)
- `server_url` (String) Server URL (defaults to https://api.netbird.io)
- `token_auth` (String, Sensitive)
