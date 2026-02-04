---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_snmp_community_hosts"
description: |-
  Configure FortiWEB system snmp community hosts configuration.
---

# fortiweb_system_snmp_community_hosts
Configure FortiWEB system snmp community hosts configuration.

## Example Usage
```hcl
resource "fortiweb_system_snmp_community_hosts" "test1" {
	mkey = "1"
	sub_mkey = "1"
	ip = "0.0.0.0"
}

resource "fortiweb_system_snmp_community_hosts" "test2" {
        mkey = "1"
        sub_mkey = "2"
        ip = "2.2.2.2"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The ID of the system snmp community.
* `id` - The ID of the system snmp community hosts entry.
* `ip` - The IP address of the SNMP manager that, if traps and/or queries are enabled in this community.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Snmp Community Hosts can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_snmp_community_hosts.labelname {mkey}
```
