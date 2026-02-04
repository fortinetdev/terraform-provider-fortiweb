---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_certificate_letsencrypt_san_list"
description: |-
  Configure FortiWEB system certificate letsencrypt san list configuration.
---

# fortiweb_system_certificate_letsencrypt_san_list
Configure FortiWEB system certificate letsencrypt san list configuration.

## Example Usage
```hcl
resource "fortiweb_system_certificate_letsencrypt_san_list" "test" {
	mkey = "test"
	sub_mkey = "1"
	san_dns = "1.1.1.1"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system certificate letsencrypt.
* `sub_mkey` - The ID of the system certificate letsencrypt san list entry.
* `san_dns` - Enter domain names.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Certificate Letsencrypt San List can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_certificate_letsencrypt_san_list.labelname {mkey}
```
