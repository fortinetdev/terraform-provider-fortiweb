---
subcategory: "FortiWEB User"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_user_admin_usergrp_members"
description: |-
  Configure FortiWEB user admin usergrp members configuration.
---

# fortiweb_user_admin_usergrp_members
Configure FortiWEB user admin usergrp members configuration.

## Example Usage
```hcl
resource "fortiweb_user_admin_usergrp_members" "test" {
	mkey = "test"
	sub_mkey = "2"
	type = "tacacs+"
	tacacs_name = "testacacs"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the user admin usergrp
* `sub_mkey` - The id of the user admin usergrp members entry.
* `type` - Select the protocol used for the query, LDAP, RADIUS, PKI or TACACS+.
* `ldap_name` - The name of an existing LDAP account query.
* `radius_name` - The name of an existing RADIUS account query.
* `pki_name` - The name of an existing PKI user.
* `tacacs_name` - The name of an existing TACACS+.
* `group_name` - The name of the remote authentication group.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

User Admin Usergrp Members can be imported using any of these accepted formats:
```
$ terraform import fortiweb_user_admin_usergrp_members.labelname {mkey}
```
