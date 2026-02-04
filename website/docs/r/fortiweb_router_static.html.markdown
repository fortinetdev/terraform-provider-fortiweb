---
subcategory: "FortiWEB Router"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_router_static"
description: |-
  Configure FortiWEB router static configuration.
---

# fortiweb_router_static
Configure FortiWEB router static configuration.

## Example Usage
```hcl
resource "fortiweb_router_static" "test" {
	mkey = "1"
	dst = "0.0.0.0/0"
	gateway = "192.0.133.1"
	device = "port1"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The ID of the router static entry.
* `dst` - The destination IP address and netmask of traffic that will be subject to this route, separated with a space.
* `gateway` - The IP address of a next-hop router.
* `device` - The name of the network interface device, such as port1, through which traffic subject to this route will be outbound. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Router Static can be imported using any of these accepted formats:
```
$ terraform import fortiweb_router_static.labelname {mkey}
```
