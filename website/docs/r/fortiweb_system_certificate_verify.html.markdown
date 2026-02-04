---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_certificate_verify"
description: |-
  Configure FortiWEB system certificate verify configuration.
---

# fortiweb_system_certificate_verify
Configure FortiWEB system certificate verify configuration.

## Example Usage
```hcl
resource "fortiweb_system_certificate_verify" "test" {
	mkey = "testcertverify"
	ca = "test"
	crl = "test"
	ocsp = "testocsprep"
	publish_dn = "enable"
	strictly_need_cert = "disable"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system certificate verify entry.
* `ca` - The name of an existing CA Group that you want to use to authenticate client certificates.
* `crl` - The name of an existing CRL Group, if any, to use to verify the revocation status of client certificates.
* `ocsp` - Specificy the name of an OCSP group.
* `publish_dn` - Enable to list only certificates related to the specified CA Group. Valid values: enable, disable.
* `strictly_need_cert` - Enable to strictly require verifying the client certificate. Valid values: enable, disable.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Certificate Verify can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_certificate_verify.labelname {mkey}
```
