---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_ntp_ntpserver"
description: |-
  Configure FortiWEB system ntp ntpserver configuration.
---

# fortiweb_system_ntp_ntpserver
Configure FortiWEB system ntp ntpserver configuration.

## Example Usage
```hcl
resource "fortiweb_system_ntp_ntpserver" "test" {
	sub_mkey = "2"
	mkey = "test"
	authentication = "enable"
	ip_type = "v4"
	key_id = 1
	key_type = "sha1"
	key = "DC2948BC9C28202DD173699014AFFBB7B2D89F5B"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `sub_mkey` - The ID of the system ntp ntpserver entry.
* `mkey` - Specify the IP address or domain name of an NTP server or pool.
* `authentication` - Enable to apply authentication keys to secure the NTP server. Valid values: enable, disable.
* `ip_type` - The IP type. Valid values: v4, v6, both.
* `key_id` - Specify the Key ID. Valid values: 0-65536.
* `key_type` - The key type. Valid values: aes128, aes256, sha1, sha256.
* `key` - Specify the Key in hexadecimal format.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Ntp Ntpserver can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_ntp_ntpserver.labelname {mkey}
```
