---
subcategory: "FortiWEB User"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_user_radius_user"
description: |-
  Configure FortiWEB user radius user configuration.
---

# fortiweb_user_radius_user
Configure FortiWEB user radius user configuration.

## Example Usage
```hcl
resource "fortiweb_user_radius_user" "test" {
	mkey = "test"
	server = "172.23.133.1"
	secret = "123"
	server_port = "1812"
	secondary_server = "172.23.133.3"
	secondary_secret = "12345"
	secondary_server_port = "1812"
	auth_type = "default"
	nas_ip = "8.8.8.8"
	
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the user radius user entry.
* `server` - The IP address or domain name of the RADIUS server to query for users.
* `server_port` - The port number where the RADIUS server listens. Valid values: 1–65535.
* `secret` - The RADIUS server secret key for the primary RADIUS server.
* `secondary_server` - The IP address or domain name of the secondary RADIUS server.
* `secondary_server_port` - The port number where the secondary RADIUS server listens. Valid values: 1–65535.
* `secondary_secret` - The RADIUS server secret key for the secondary RADIUS server.
* `nas_ip` - The NAS IP address and called station ID. 
* `auth_type` - The authentication method. Valid values: default, chap, ms_chap, ms_chap_v2, pap.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

User Radius User can be imported using any of these accepted formats:
```
$ terraform import fortiweb_user_radius_user.labelname {mkey}
```
