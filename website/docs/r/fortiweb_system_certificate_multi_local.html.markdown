---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_certificate_multi_local"
description: |-
  Configure FortiWEB system certificate multi local configuration.
---

# fortiweb_system_certificate_multi_local
Configure FortiWEB system certificate multi local configuration.

## Example Usage
```hcl
resource "fortiweb_system_certificate_multi_local" "test" {
	mkey = "test"
	rsa_cert = "sign-p12"
	comment = "comment123"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system certificate multi local entry.
* `comment` - A description or other comment. 
* `rsa_cert` - Select the RSA certificate created in system certificate local.
* `dsa_cert` - Select the DSA certificate created in system certificate local.
* `ecc_cert` - Select the ECDSA certificate created in system certificate local.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Certificate Multi Local can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_certificate_multi_local.labelname {mkey}
```
