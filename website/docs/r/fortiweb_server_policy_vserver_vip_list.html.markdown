---
subcategory: "FortiWEB Server Policy"
layout: "fortiadc"
page_title: "FortiWEB: fortiweb_server_policy_vserver_vip_list"
description: |-
  Configure FortiWEB server policy vserver vip list configuration.
---

# fortiweb_server_policy_vserver_vip_list
Configure FortiWEB server policy vserver vip list configuration.

## Example Usage
```hcl
resource "fortiweb_server_policy_vserver_vip_list" "test" {
        mkey = "test"
        sub_mkey = "1"
        interface = "port2"
        status = "enable"
        use_interface_ip = "enable"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the server policy vserver.
* `sub_mkey` - The ID of the server policy vserver vip list entry.
* `interface` - Enter the name of the network interface or bridge.
* `status` - Enable to accept traffic destined for this virtual server. Valid values: enable, disable.
* `use_interface_ip` - For FortiWeb-VM on Microsoft Azure, specify whether the virtual server uses the IP address of the specified interface, instead of an IP specified by vip or vip6. Valid values: enable, disable.
* `vip` - Enter the name of the vip (Virtual IP) item.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Server Policy Vserver Vip List can be imported using any of these accepted formats:
```
$ terraform import fortiweb_server_policy_vserver_vip_list.labelname {mkey}
```
