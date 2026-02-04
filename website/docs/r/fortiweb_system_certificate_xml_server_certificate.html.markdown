---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_certificate_xml_server_certificate"
description: |-
  Configure FortiWEB system certificate xml server certificate configuration.
---

# fortiweb_system_certificate_xml_server_certificate
Configure FortiWEB system certificate xml server certificate configuration.

## Example Usage
```hcl
resource "fortiweb_system_certificate_xml_server_certificate" "test" {
	mkey = "sign-server"
	certificatefile = "./sign-server.crt"
	keyfile = "./sign-server.key"
	password = "fortinet"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system certificate xml server certificate entry.
* `certificatefile` - The certificate.
* `keyfile` - The secret key.
* `passwd` - The password of secret key.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Certificate Xml Server Certificate can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_certificate_xml_server_certificate.labelname {mkey}
```
