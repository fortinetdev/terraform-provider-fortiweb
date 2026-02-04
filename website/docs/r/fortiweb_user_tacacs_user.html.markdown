---
subcategory: "FortiWEB User"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_user_tacacs_user"
description: |-
  Configure FortiWEB user tacacs user configuration.
---

# fortiweb_user_tacacs_user
Configure FortiWEB user tacacs user configuration.

## Example Usage
```hcl
resource "fortiweb_user_tacacs_user" "test" {
	mkey = "test"
	server = "172.23.133.1"
	secret = "123"
	auth_type = "chap"
	type = "specify"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the user tacacs user entry.
* `server` - The IP address or domain name of the TACACS+ server.
* `secret` - The TACACS+ server secret key for the TACACS+ server.
* `auth_type` - Specify a type among MSCHAP, CHAP, PAP, or ASCII when type is Specify. Valid values: auto, ms_chap, chap, pap, ascii. Option MSCHAP may be deprecated in version after 8.0.3 of FortiWeb.
* `type` - Select Auto to automatically assign an authentication type or select Specify to specify a type among MSCHAP, CHAP, PAP, and ASCII. Valid values: auto, specify.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

User Tacacs User can be imported using any of these accepted formats:
```
$ terraform import fortiweb_user_tacacs_user.labelname {mkey}
```
