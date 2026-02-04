---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_certificate_tsl_ca"
description: |-
  Configure FortiWEB system certificate tsl ca configuration.
---

# fortiweb_system_certificate_tsl_ca
Configure FortiWEB system certificate tsl ca configuration.

## Example Usage
```hcl
resource "fortiweb_system_certificate_tsl_ca" "test" {
	mkey = "test"
	type = "localPC"
	srcfile = "./eu-lotl.xml"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system certificate tsl ca entry.
* `url` - Enter the URL of the TSL CA certificate.
* `srcfile` - TSL file.
* `type` - Select the way to upload a TSL. Valid values: url, localPC.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Certificate Tsl Ca can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_certificate_tsl_ca.labelname {mkey}
```
