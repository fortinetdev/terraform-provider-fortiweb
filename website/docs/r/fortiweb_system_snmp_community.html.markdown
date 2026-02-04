---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_snmp_community"
description: |-
  Configure FortiWEB system snmp community configuration.
---

# fortiweb_system_snmp_community
Configure FortiWEB system snmp community configuration.

## Example Usage
```hcl
resource "fortiweb_system_snmp_community" "test" {
	mkey = "1"
	name = "test"
	status = "enable"
	query_v1_status = "enable"
	query_v1_port = 161
	query_v2c_status = "enable"
	query_v2c_port = 162
	trap_v1_status = "enable"
	trap_v1_lport = 163
	trap_v1_rport = 164
	trap_v2c_status = "enable"
	trap_v2c_lport = 165
	trap_v2c_rport = 166
	events = "cpu-high mem-low log-full sys-mode-change intf-ip sys-ha-cluster-status-change sys-ha-member-join sys-ha-member-leave policy-start policy-stop pserver-failed waf-amethod-attack waf-pvalid-attack waf-hidden-fields waf-url-access-attack waf-signature-detection netlink-up-status netlink-down-status power-supply-failure policy-ldap-auth-failure policy-radius-auth-failure"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The ID of the system snmp community entry.
* `name` - The name of the SNMP community to which the FortiWeb appliance and at least one SNMP manager belongs.
* `status` - Enable to activate the community. Valid values: enable, disable.
* `query_v1_status` - Enable to respond to queries using the SNMP v1 version of the SNMP protocol. Valid values: enable, disable.
* `query_v1_port` - The port number on which the FortiWeb appliance will listen for SNMP v1 queries from the SNMP managers of the community. Valid values: 1–65535.
* `query_v2c_status` - Enable to respond to queries using the SNMP v2c version of the SNMP protocol. Valid values: enable, disable.
* `query_v2c_port` - The port number on which the FortiWeb appliance will listen for SNMP v2c queries from the SNMP managers of the community. Valid values: 1–65535.
* `trap_v1_status` - Enable to send traps using the SNMP v1 version of the SNMP protocol. Valid values: enable, disable.
* `trap_v1_lport` - The port number that will be the source port number for SNMP v1 trap packets. Valid values: 1–65535.
* `trap_v1_rport` - The port number that will be the destination port number for SNMP v1 trap packets. Valid values: 1–65535.
* `trap_v2c_status` - Enable to send traps using the SNMP v2c version of the SNMP protocol. Valid values: enable, disable.
* `trap_v2c_lport` - The port number that will be the source port number for SNMP v2c trap packets. Valid values: 1–65535.
* `trap_v2c_rport` - The port number that will be the destination port number for SNMP v2c trap packets. Valid values: 1–65535.
* `events` - One or more of the following SNMP event names in order to cause the FortiWeb appliance to send traps when those events occur. Valid values: cpu-high, intf-ip, log-full, mem-low, netlink-down-status, netlink-up-status, policy-start, policy-stop, pserver-failed, sys-ha-cluster-status-change, sys-ha-member-join, sys-ha-member-leave, sys-mode-change, waf-amethod-attack, waf-hidden-fields, waf-pvalid-attack, waf-signature-detection, power-supply-failure.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Snmp Community can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_snmp_community.labelname {mkey}
```
