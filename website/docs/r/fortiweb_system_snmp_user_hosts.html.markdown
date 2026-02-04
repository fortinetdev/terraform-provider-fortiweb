---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_snmp_user_hosts"
description: |-
  Configure FortiWEB system snmp user hosts configuration.
---

# fortiweb_system_snmp_user_hosts
Configure FortiWEB system snmp user hosts configuration.

## Example Usage
```hcl
resource "fortiweb_system_snmp_user_hosts" "test" {
	mkey = "test"
	sub_mkey = "1"
	ip = "0.0.0.0"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system snmp user.
* `sub_mkey` - The ID of the system snmp user hosts entry.
* `ip` - The IP address of the SNMP manager that can do the following when you enable traps, queries, or both in this community.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Snmp User Hosts can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_snmp_user_hosts.labelname {mkey}
```
