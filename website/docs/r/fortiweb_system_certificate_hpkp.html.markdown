---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_certificate_hpkp"
description: |-
  Configure FortiWEB system certificate hpkp configuration.
---

# fortiweb_system_certificate_hpkp
Configure FortiWEB system certificate hpkp configuration.

## Example Usage
```hcl
resource "fortiweb_system_certificate_hpkp" "test" {
	mkey = "test"
	pin_sha256 = "qqq aaa"
	max_age = 1296000
	subdomains = "enable"
	report_only = "enable"
	report_uri = "aaa.com"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system certificate hpkp profile entry.
* `pin_sha256` - Enter a Base64 encoded SPKI fingerprint.
* `max_age` - Enter an interval in which the client will use the SPKI fingerprint to attempt to access the server. Valid values: 0-31536000.
* `subdomains` - Enable to apply the public key pinning rule to all of the server's subdomains. Valid values: enable, disable.
* `report_only` - Enable so that FortiWeb sends reports to the specified Report URI.
* `report_uri` - Enter a URI to which FortiWeb will send pin validation failures.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Certificate Hpkp can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_certificate_hpkp.labelname {mkey}
```
