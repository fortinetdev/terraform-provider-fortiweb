---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_certificate_ocsp_responder"
description: |-
  Configure FortiWEB system certificate ocsp responder configuration.
---

# fortiweb_system_certificate_ocsp_responder
Configure FortiWEB system certificate ocsp responder configuration.

## Example Usage
```hcl
resource "fortiweb_system_certificate_ocsp_responder" "test" {
	mkey = "testocsprep"
	ocsp_signing_certs = "OCSP_SIGN_CERTS_1"
	ocsp_url = "0.0.0.0"
	timeout = "15"
	caching = "enable"
	caching_ttl = 3600
	comment = "testtest1234"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system certificate ocsp responder entry.
* `ocsp_url` - Enter URL of the OCSP server corresponding to the specified CA certificate. 
* `ocsp_signing_certs` - Ocsp-Signing-Certs setting.
* `timeout` - Timeout setting.
* `comment` - Comment setting.
* `caching` - Caching setting.
* `caching_ttl` - Caching-Ttl setting.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Certificate Ocsp Responder can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_certificate_ocsp_responder.labelname {mkey}
```
