---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_certificate_offline_sni_members"
description: |-
  Configure FortiWEB system certificate offline sni members configuration.
---

# fortiweb_system_certificate_offline_sni_members
Configure FortiWEB system certificate offline sni members configuration.

## Example Usage
```hcl
resource "fortiweb_system_certificate_offline_sni_members" "test" {
	mkey = "testofflinesni"
	sub_mkey = "1"
	domain_type = "plain"
	domain = "0.0.0.0"
	local_cert = "sign-p12"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system certificate offline sni.
* `sub_mkey` - The ID of the system certificate offline sni members entry.
* `domain_type` - Specify plain/regular to match multiple domains to certificates using a literal domain/regular expression specified in domain. Valid values: plain, regular.
* `domain` - The domain of the secure website (HTTPS) that uses the certificate specified by local-cert.
* `local_cert` - The name of the server certificate that FortiWeb uses to encrypt or decrypt SSL-secured connections for the website specified by domain. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Certificate Offline Sni Members can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_certificate_offline_sni_members.labelname {mkey}
```
