---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_certificate_sni_members"
description: |-
  Configure FortiWEB system certificate sni members configuration.
---

# fortiweb_system_certificate_sni_members
Configure FortiWEB system certificate sni members configuration.

## Example Usage
```hcl
resource "fortiweb_system_certificate_sni_members" "test1" {
	mkey = "test"
	sub_mkey = "1"
	domain_type = "plain"
	domain = "0.0.0.0"
	certificate_type = "disable"
	local_cert = "sign-p12"
	multi_local_cert = "disable"
	inter_group = "testCAGroup"
	verify = "testcertverify"
}

resource "fortiweb_system_certificate_sni_members" "test2" {
        mkey = "test"
        sub_mkey = "2"
        domain_type = "plain"
        domain = "1.1.1.1"
        certificate_type = "disable"
        multi_local_cert = "enable"
	multi_local_cert_group = "test"
        inter_group = "testCAGroup"
        verify = "testcertverify"
}

resource "fortiweb_system_certificate_sni_members" "test3" {
        mkey = "test"
        sub_mkey = "3"
        domain_type = "plain"
        domain = "2.2.2.2"
        certificate_type = "enable"
        multi_local_cert = "disable"
	lets_certificate = "test"
        inter_group = "testCAGroup"
        verify = "testcertverify"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system certificate sni.
* `sub_mkey` - The ID of the system certificate sni members entry.
* `domain` - The domain of the secure website (HTTPS) that uses the certificate specified by local-cert.
* `domain_type` - Specify plain/regular to match multiple domains to certificates using a literal domain/regular expression specified in domain.
* `certificate_type` - Enable allow FortiWeb to automatically retrieve CA certificates from Let's Encrypt. Valid values: enable, disable.
* `local_cert` - The name of the server certificate that FortiWeb uses to encrypt or decrypt SSL-secured connections for the website specified by domain.
* `multi_local_cert` - Enable this option to allow FortiWeb to use multiple local certificates. Valid values: enable, disable.
* `multi_local_cert_group` - Select the multi-certificate you have created.
* `lets_certificate` - Select the Letsencrypt certificate you have created.
* `inter_group` - The name of a group of intermediate certificate authority (CA) certificates, if any, that FortiWeb presents to validate the CA signature of the certificate specified by local-cert.
* `verify` - The name of a certificate verifier, if any, that FortiWeb uses when an HTTP client presents its personal certificate.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Certificate Sni Members can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_certificate_sni_members.labelname {mkey}
```
