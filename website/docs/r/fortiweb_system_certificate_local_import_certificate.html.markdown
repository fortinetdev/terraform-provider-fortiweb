---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_certificate_local_import_certificate"
description: |-
  Configure FortiWEB system certificate local import certificate configuration.
---

# fortiweb_system_certificate_local_import_certificate
Configure FortiWEB system certificate local import certificate configuration.

## Example Usage
```hcl
resource "fortiweb_system_certificate_local_import_certificate" "test1" {
	mkey = "sign-server"
	certificatefile = "./sign-server.crt"
	keyfile = "./sign-server.key"
	password = "fortinet"
}

resource "fortiweb_system_certificate_local_import_certificate" "test2" {
        mkey = "sign-p12"
        certificatewithkeyfile = "./sign-p12.p12"
        password = "fortinet"
}

resource "fortiweb_system_certificate_local_import_certificate" "test3" {
        mkey = "sign-test"
        certificatefile = "./test.crt"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system certificate local entry.
* `certificatefile` - The name of a certificate file.
* `keyfile` - The name of a key file
* `certificatewithkeyfile` - The name of a certificate file.
* `password` - The password for the certificate.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Certificate Local can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_certificate_local_import_certificate.labelname {mkey}
```
