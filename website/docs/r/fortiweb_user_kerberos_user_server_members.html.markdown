---
subcategory: "FortiWEB User"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_user_kerberos_user_server_members"
description: |-
  Configure FortiWEB user kerberos user server members configuration.
---

# fortiweb_user_kerberos_user_server_members
Configure FortiWEB user kerberos user server members configuration.

## Example Usage
```hcl
resource "fortiweb_user_kerberos_user_server_members" "test" {
	mkey = "test"
	sub_mkey = "2"
	server = "172.23.133.1"
	port = "89"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the user kerberos user.
* `id` - The id of the user kerberos user server members entry.
* `server` - The IP address of the KDC.
* `port` - The port the KDC uses to listen for requests.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

User Kerberos User Server Members can be imported using any of these accepted formats:
```
$ terraform import fortiweb_user_kerberos_user_server_members.labelname {mkey}
```
