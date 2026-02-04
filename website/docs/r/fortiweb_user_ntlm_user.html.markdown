---
subcategory: "FortiWEB User"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_user_ntlm_user"
description: |-
  Configure FortiWEB user ntlm user configuration.
---

# fortiweb_user_ntlm_user
Configure FortiWEB user ntlm user configuration.

## Example Usage
```hcl
resource "fortiweb_user_ntlm_user" "test" {
	mkey = "test"
	server = "172.23.133.1"
	port = "445"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the user ntlm user entry.
* `port` - The port number where the NTLM server listens. Valid values: 1–65535.
* `server` - The IP address of the NTLM server.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

User Ntlm User can be imported using any of these accepted formats:
```
$ terraform import fortiweb_user_ntlm_user.labelname {mkey}
```
