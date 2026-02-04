---
layout: "fortiweb"
page_title: "Provider: FortiWEB"
sidebar_current: "docs-fortiweb-index"
description: |-
  The FortiWEB provider interacts with FortiWEB.
---

# FortiWEB Provider

The FortiWEB provider is used to interact with the resources supported by FortiWEB. We need to configure the provider with the proper credentials before it can be used. Please use the navigation on the left to read more details about the available resources.

## Example Usage

```hcl
# Configure the FortiWEB Provider for FortiWEB
provider "fortiweb" {
	hostname = "192.168.52.177"
	insecure = "true"
	vdom	 = "root"
	username = "admin"
	password = "fortinet"
}

# Create a Static Route Item
resource "fortiweb_router_static" "test1" {
	mkey	= "1"
	dst	= "110.2.2.122/32"
	gateway = "2.2.2.2"
	device	= "port1"
}
```

If it is used for testing, you can set `insecure` to "true" and unset `cabundlefile` to quickly set the provider up, for example:

```hcl
provider "fortiweb" {
	hostname = "192.168.52.177"
	insecure = "true"
	username = "admin"
	password = "fortinet"
}
```

Please refer to the Argument Reference below for more help on `insecure` and `cabundlefile`.

## VDOM

If the FortiWEB unit is running in VDOM mode, the `vdom` configuration needs to be added.

Usage:

```hcl
provider "fortiweb" {
	hostname = "192.168.52.177"
	insecure = "false"
	username = "admin"
	password = "fortinet"
	vdom     = "vdomtest"
}

# Create a Static Route Item
resource "fortiweb_router_static" "test1" {
	mkey	= "1"
	dst	= "110.2.2.122/32"
	gateway = "2.2.2.2"
	device	= "port1"
}
```

By default, each resource inherits the provider's global vdom settings, but it can also set its own vdom through the `vdom` of each resource. See the `vdom` argument of each resource for details.



## Release
Check out the FortiWEB provider release notes and additional information from: [the FortiWEB provider releases](https://github.com/fortinetdev/terraform-provider-fortiweb/releases).

## Versioning

The provider can cover FortiWEB 8.0 versions, the configuration of all parameters should be based on the relevant FortiWEB version manual.
