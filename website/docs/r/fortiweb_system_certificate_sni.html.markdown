---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_certificate_sni"
description: |-
  Configure FortiWEB system certificate sni configuration.
---

# fortiweb_system_certificate_sni
Configure FortiWEB system certificate sni configuration.

## Example Usage
```hcl
resource "fortiweb_system_certificate_sni" "test" {
	mkey = "testsni"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system certificate sni entry.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Certificate Sni can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_certificate_sni.labelname {mkey}
```
