---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_snmp_sysinfo"
description: |-
  Configure FortiWEB system snmp sysinfo configuration.
---

# fortiweb_system_snmp_sysinfo
Configure FortiWEB system snmp sysinfo configuration.

## Example Usage
```hcl
resource "fortiweb_system_snmp_sysinfo" "test" {
	status = "enable"
	description = "Description"
	location = "Location"
	contact_info = "Contact"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `status` - Enable to activate the SNMP agent. Valid values: enable, disable.
* `description` - A description of the FortiWeb appliance.
* `contact_info` - The contact information for the administrator or other person responsible for this FortiWeb appliance.
* `location` - The physical location of the FortiWeb appliance. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Snmp Sysinfo can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_snmp_sysinfo.labelname {mkey}
```
