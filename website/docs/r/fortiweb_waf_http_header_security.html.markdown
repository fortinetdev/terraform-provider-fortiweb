---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_http_header_security"
description: |-
  Configure FortiWEB waf http header security configuration.
---

# fortiweb_waf_http_header_security
Configure FortiWEB waf http header security configuration.

## Example Usage
```hcl
resource "fortiweb_waf_http_header_security" "test" {
	mkey = "test"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf http header security entry.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf Http Header Security can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_http_header_security.labelname {mkey}
```
