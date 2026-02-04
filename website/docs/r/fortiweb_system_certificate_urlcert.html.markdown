---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_certificate_urlcert"
description: |-
  Configure FortiWEB system certificate urlcert configuration.
---

# fortiweb_system_certificate_urlcert
Configure FortiWEB system certificate urlcert configuration.

## Example Usage
```hcl
resource "fortiweb_system_certificate_urlcert" "test" {
	mkey = "testurlcert"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system certificate urlcert entry.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Certificate Urlcert can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_certificate_urlcert.labelname {mkey}
```
