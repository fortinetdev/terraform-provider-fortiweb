---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_vip"
description: |-
  Configure FortiWEB system vip configuration.
---

# fortiweb_system_vip
Configure FortiWEB system vip configuration.

## Example Usage
```hcl
resource "fortiweb_system_vip" "test" {
	mkey = "test"
	vip = "192.0.0.1/24"
	vip6 = "::/0"
	interface = "port3"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system vip entry.
* `vip` -  The IPv4 address and subnet of the virtual IP. Valid example: 0.0.0.0/0.
* `vip6` - The IPv6 address and subnet of the virtual IP. Valid example: ::/0.
* `interface` - The name of the network interface or bridge the virtual IP is bound to and where traffic destined for the virtual IP arrives.
* `domains` - The ADOM you want to create this virtual IP in.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Vip can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_vip.labelname {mkey}
```
