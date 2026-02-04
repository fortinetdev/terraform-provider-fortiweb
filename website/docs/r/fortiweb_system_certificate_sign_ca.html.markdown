---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_certificate_sign_ca"
description: |-
  Configure FortiWEB system certificate sign ca configuration.
---

# fortiweb_system_certificate_sign_ca
Configure FortiWEB system certificate sign ca configuration.

## Example Usage
```hcl
resource "fortiweb_system_certificate_sign_ca" "test1" {
	mkey = "sign-server"
	upfile = "./sign-server.crt"
	keyfile = "./sign-server.key"
	passwd = "fortinet"
	comment = "test123"
}

resource "fortiweb_system_certificate_sign_ca" "test2" {
	mkey = "sign-p12"
	upfile = "./sign-p12.p12"
	passwd = "fortinet"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system certificate sign ca entry.
* `upfile` - The name of a certificate file.
* `keyfile` - The name of a key file.
* `passwd` - The password for the certificate.
* `comment` -  	Enter a description or other comment.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Certificate Sign Ca can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_certificate_sign_ca.labelname {mkey}
```
