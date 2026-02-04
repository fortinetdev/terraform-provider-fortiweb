---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_certificate_letsencrypt"
description: |-
  Configure FortiWEB system certificate letsencrypt configuration.
---

# fortiweb_system_certificate_letsencrypt
Configure FortiWEB system certificate letsencrypt configuration.

## Example Usage
```hcl
resource "fortiweb_system_certificate_letsencrypt" "test" {
	mkey = "testlets"
	domain = "1.1.1.1"
	validation_method = "HTTP-01"
	key_type = "RSA-2048"
	renew_period = 30
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system certificate letsencrypt entry.
* `domain` - The domain name of your application.
* `validation_method` - Valid values: HTTP-01, TLS-ALPN-01, DNS-01.
* `key_type` - Key Type. Valid values: RSA-2048, RSA-3072, RSA-4096.
* `renew_period` - Set how soon FortiWeb obtains the TLS certificate from Let’s Encrypt. Valid values: 1-60.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Certificate Letsencrypt can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_certificate_letsencrypt.labelname {mkey}
```
