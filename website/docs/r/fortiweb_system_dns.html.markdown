---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_dns"
description: |-
  Configure FortiWEB system dns configuration.
---

# fortiweb_system_dns
Configure FortiWEB system dns configuration.

## Example Usage
```hcl
resource "fortiweb_system_dns" "test" {
	primary = "8.8.8.8"
	secondary = "0.0.0.0"
	domain = ""
}
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `primary` - The IP address of the primary DNS server.
* `secondary` - The IP address of the secondary DNS server.
* `domain` - The name of the local domain to which the FortiWeb appliance belongs, if any. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Dns can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_dns.labelname {mkey}
```
