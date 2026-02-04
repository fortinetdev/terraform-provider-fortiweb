---
subcategory: "FortiWEB Router"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_router_policy"
description: |-
  Configure FortiWEB router policy configuration.
---

# fortiweb_router_policy
Configure FortiWEB router policy configuration.

## Example Usage
```hcl
resource "fortiweb_router_policy" "test" {
	mkey = "1"
	iif = "port3"
	src = "0.0.0.0/0"
	dst = "0.0.0.0/0"
	action = "forward-traffic"
	oif = "port4"
	gateway = "0.0.0.0"
	priority = 200
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The ID of the router policy entry.
* `src` -  The source IP address and netmask to match, separated with a space.
* `dst` - The destination IP address and netmask to match, separated with a space.
* `action` - Action setting. Valid values: forward-traffic, stop-policy-routing.
* `gateway` - The IP address of a next-hop router.
* `iif` - The name of the interface, such as port1, on which FortiWeb receives packets it applies this routing policy to.
* `oif` - The name of the interface, such as port2, through which FortiWeb routes packets that match the specified IP address information.
* `priority` - Priority setting. Valid values: 1-200.
* `fwmark` - The Fwmark value specified in Firewall Fwmark Policy. Valid values: 0-255.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Router Policy can be imported using any of these accepted formats:
```
$ terraform import fortiweb_router_policy.labelname {mkey}
```
