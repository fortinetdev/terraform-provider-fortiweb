---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_admin"
description: |-
  Configure FortiWEB system admin configuration.
---

# fortiweb_system_admin
Configure FortiWEB system admin configuration.

## Example Usage
```hcl
resource "fortiweb_system_admin" "test" {
	mkey = "test"
	type = "local-user"
	trusthostv4 = "0.0.0.0/0 "
	trusthostv6 = "::/0 "
	access_profile = "admin_no_access"
	force_password_change = "disable"
	password = "test"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system admin entry.
* `type` - Valid values: local-user, remote-user.
* `admin_usergrp` - Enter the name of the remote authentication group whose settings the FortiWeb appliance will use to connect to a remote authentication server when authenticating login attempts for this account.
* `wildcard` - Used when administrator accounts authenticate via a RADIUS query. Valid values: enable, disable.
* `trusthostv4` - To allow login attempts from any IP address, enter 0.0.0.0/0.
* `trusthostv6` - To allow login attempts from any IP address, enter ::/0. 
* `access_profile` - Enter the name of an access profile that gives the permissions for this administrator account.
* `force_password_change` - Enable force password change for next login. Valid values: enable. disable.
* `password` - Enter a password for the administrator account.
* `fortiai` - Enable to use FortiAI User. Valid values: enable, disable.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Admin can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_admin.labelname {mkey}
```
