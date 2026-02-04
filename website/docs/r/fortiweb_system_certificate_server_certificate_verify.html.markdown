---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_certificate_server_certificate_verify"
description: |-
  Configure FortiWEB system certificate server certificate verify configuration.
---

# fortiweb_system_certificate_server_certificate_verify
Configure FortiWEB system certificate server certificate verify configuration.

## Example Usage
```hcl
resource "fortiweb_system_certificate_server_certificate_verify" "test" {
	mkey = "testservercertverify"
	ca = "test"
	crl = "test"
	ocsp = "testocsprep"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system certificate server certificate verify entry.
* `ca` - The name of an existing CA Group that you want to use to authenticate client certificates.
* `crl` - The name of an existing CRL Group, if any, to use to verify the revocation status of client certificates.
* `ocsp` - the name of an existing Ocsp Group that you want to use to authenticate client certificates.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Certificate Server Certificate Verify can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_certificate_server_certificate_verify.labelname {mkey}
```
