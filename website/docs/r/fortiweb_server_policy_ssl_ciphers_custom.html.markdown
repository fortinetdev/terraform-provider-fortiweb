---
subcategory: "FortiWEB Server Policy"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_server_policy_ssl_ciphers_custom"
description: |-
  Configure FortiWEB server policy ssl ciphers custom configuration.
---

# fortiweb_server_policy_ssl_ciphers_custom
Configure FortiWEB server policy ssl ciphers custom configuration.

## Example Usage
```hcl
resource "fortiweb_server_policy_ssl_ciphers_custom" "test" {
	mkey = "test"
	tls_v10 = "enable"
	tls_v11 = "enable"
	tls_v12 = "enable"
	tls_v13 = "disable"
	ssl_cipher = "custom"
	ssl_custom_cipher = "ECDHE-ECDSA-AES256-GCM-SHA384 ECDHE-RSA-AES256-GCM-SHA384 ECDHE-ECDSA-CHACHA20-POLY1305 ECDHE-RSA-CHACHA20-POLY1305 ECDHE-ECDSA-AES128-GCM-SHA256 ECDHE-RSA-AES128-GCM-SHA256 ECDHE-ECDSA-AES256-SHA384 ECDHE-RSA-AES256-SHA384 ECDHE-ECDSA-AES128-SHA256 ECDHE-RSA-AES128-SHA256 ECDHE-ECDSA-AES256-SHA ECDHE-RSA-AES256-SHA ECDHE-ECDSA-AES128-SHA ECDHE-RSA-AES128-SHA AES256-GCM-SHA384 AES128-GCM-SHA256 AES256-SHA256 AES128-SHA256" 
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the server policy ssl ciphers custom entry.
* `tls_v10` - Specifies whether clients can connect securely to FortiWeb using the TLS 1.0 cryptographic protocol. Valid values: enable, disable.
* `tls_v11` - Specifies whether clients can connect securely to FortiWeb using the TLS 1.1 cryptographic protocol. Valid values: enable, disable.
* `tls_v12` - Specifies whether clients can connect securely to FortiWeb using the TLS 1.2 cryptographic protocol. Valid values: enable, disable.
* `tls_v13` - Specifies whether clients can connect securely to FortiWeb using the TLS 1.3 cryptographic protocol. Valid values: enable, disable.
* `ssl_cipher` - Specify whether the set of cipher suites that FortiWeb allows creates a medium-security, high-security, or custom configuration. Valid values: medium, high, custom.
* `ssl_custom_cipher` - Specify one or more cipher suites that FortiWeb allows. Valid values: ECDHE-ECDSA-AES256-GCM-SHA384, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-ECDSA-CHACHA20-POLY1305, ECDHE-RSA-CHACHA20-POLY1305, ECDHE-ECDSA-AES128-GCM-SHA256, ECDHE-RSA-AES128-GCM-SHA256, ECDHE-ECDSA-AES256-SHA384, ECDHE-RSA-AES256-SHA384, ECDHE-ECDSA-AES128-SHA256, ECDHE-RSA-AES128-SHA256, ECDHE-ECDSA-AES256-SHA, ECDHE-RSA-AES256-SHA, ECDHE-ECDSA-AES128-SHA, ECDHE-RSA-AES128-SHA, AES256-GCM-SHA384, AES128-GCM-SHA256, AES256-SHA256, AES128-SHA256.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Server Policy Ssl Ciphers Custom can be imported using any of these accepted formats:
```
$ terraform import fortiweb_server_policy_ssl_ciphers_custom.labelname {mkey}
```
