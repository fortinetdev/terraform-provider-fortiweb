---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_certificate_crl"
description: |-
  Configure FortiWEB system certificate crl configuration.
---

# fortiweb_system_certificate_crl
Configure FortiWEB system certificate crl configuration.

## Example Usage
```hcl
resource "fortiweb_system_certificate_crl" "test" {
	mkey = "CRL_1"
	type = "localPC"
	srcfile = "./cacrl.crl"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system certificate crl entry.
* `httpurl` - Enter the HTTP URL of the certificate.
* `scepurl` - Enter the SCEP URL of the certificate.
* `srcfile` - Set the certificate. Only certificates in PEM format may be set.
* `type` - Specify how you set the certificate. Valid values: http, scep, localPC.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Certificate Crl can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_certificate_crl.labelname {mkey}
```
