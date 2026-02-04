---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_certificate_ca"
description: |-
  Configure FortiWEB system certificate ca configuration.
---

# fortiweb_system_certificate_ca
Configure FortiWEB system certificate ca configuration.

## Example Usage
```hcl
resource "fortiweb_system_certificate_ca" "test" {
	mkey = "CA_Cert_1"
	srcfile = "./ca.crt"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system certificate ca entry.
* `srcfile` - The name of certificate file.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Certificate Ca can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_certificate_ca.labelname {mkey}
```
