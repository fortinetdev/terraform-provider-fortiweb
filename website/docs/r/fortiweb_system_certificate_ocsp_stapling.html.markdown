---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_certificate_ocsp_stapling"
description: |-
  Configure FortiWEB system certificate ocsp stapling configuration.
---

# fortiweb_system_certificate_ocsp_stapling
Configure FortiWEB system certificate ocsp stapling configuration.

## Example Usage
```hcl
resource "fortiweb_system_certificate_ocsp_stapling" "test" {
	mkey = "test"
	certificate = "CA_Cert_1"
	local_cert = "sign-p12"
	ocsp_url = "0.0.0.0"
	comment = "test"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system certificate ocsp stapling entry.
* `certificate` - A CA certificate that has been imported in FortiWeb.
* `local_cert` - The local certificate of the server certificate to be queried.
* `ocsp_url` - Enter URL of the OCSP server corresponding to the specified CA certificate. 
* `comment` - Optionally, enter a comment for the OCSP group.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Certificate Ocsp Stapling can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_certificate_ocsp_stapling.labelname {mkey}
```
