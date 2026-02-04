---
subcategory: "FortiWEB Server Policy"
layout: "fortiadc"
page_title: "FortiWEB: fortiweb_server_policy_vserver"
description: |-
  Configure FortiWEB server policy vserver configuration.
---

# fortiweb_server_policy_vserver
Configure FortiWEB server policy vserver configuration.

## Example Usage
```hcl
resource "fortiweb_server_policy_vserver" "test" {
	mkey = "test"
}
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the server policy vserver entry.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Server Policy Vserver can be imported using any of these accepted formats:
```
$ terraform import fortiweb_server_policy_vserver.labelname {mkey}
```
