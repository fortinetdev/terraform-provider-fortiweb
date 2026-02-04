---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_snmp_user"
description: |-
  Configure FortiWEB system snmp user configuration.
---

# fortiweb_system_snmp_user
Configure FortiWEB system snmp user configuration.

## Example Usage
```hcl
resource "fortiweb_system_snmp_user" "test" {
	mkey = "test"
	status = "enable"
	security_level = "authpriv"
	auth_proto = "sha1"
	auth_pwd = "fortinet"
	priv_proto = "aes"
	priv_pwd = "fortinet"
	query_status = "enable"
	query_port = 161
	trap_status = "enable"
	trapport_local = 163
	trapport_remote = 164
	trapevent = "cpu-high mem-low log-full sys-mode-change intf-ip sys-ha-cluster-status-change sys-ha-member-join sys-ha-member-leave policy-start policy-stop pserver-failed waf-amethod-attack waf-pvalid-attack waf-hidden-fields waf-url-access-attack waf-signature-detection netlink-up-status netlink-down-status power-supply-failure policy-ldap-auth-failure policy-radius-auth-failure"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The ID of the system snmp user entry.
* `status` - Enable to activate the community. Valid values: enable, disable.
* `security_level` - The security level. Valid values: noauthnopriv, authnopriv, authpriv.
* `auth_proto` - Specify the authentication protocol. Valid values: sha1, sha224, sha256, sha384, sha512, md5.
* `auth_pwd` - Specify the authentication password.
* `priv_proto` - Specify the encryption protocol. Valid values: aes, des, aes256.
* `priv_pwd` - Specify the encryption password.
* `query_status` - Enable to respond to queries using the SNMP v3 version of the SNMP protocol. Valid values: enable, disable.
* `query_port` - The port number on which the FortiWeb appliance listens for SNMP v3 queries from the SNMP managers of the community. Valid values: 1–65535.
* `trap_status` - Enable to send traps using the SNMP v3 version of the SNMP protocol. Valid values: enable, disable.
* `trapport_local` - The port number that is the source port number for SNMP v3 trap packets. Valid values: 1–65535.
* `trapport_remote` - he port number that is the destination port number for SNMP v3 trap packets. Valid values: 1–65535.
* `trapevent` - The name of one or more the SNMP events. Valid values: cpu-high, intf-ip, log-full, mem-low, netlink-down-status, netlink-up-status, policy-start, policy-stop, pserver-failed, sys-ha-cluster-status-change, sys-ha-member-join, sys-ha-member-leave, sys-mode-change, waf-amethod-attack, waf-hidden-fields, waf-pvalid-attack, waf-signature-detection, waf-url-access-attack, power-supply-failure.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Snmp User can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_snmp_user.labelname {mkey}
```
