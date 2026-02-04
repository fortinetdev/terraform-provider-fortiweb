---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_cors_protection_rule_exposed_headers_list"
description: |-
  Configure FortiWEB waf cors protection rule exposed headers list configuration.
---

# fortiweb_waf_cors_protection_rule_exposed_headers_list
Configure FortiWEB waf cors protection rule exposed headers list configuration.

## Example Usage
```hcl
resource "fortiweb_waf_cors_protection_rule_exposed_headers_list" "test" {
	sub_mkey = "1"
	mkey = "test"
	header = "test"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf cors protection rule.
* `sub_mkey` - The ID of the waf cors protection rule exposed headers list entry.
* `header` - The name of exposed header.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf Cors Protection Rule Exposed Headers List can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_cors_protection_rule_exposed_headers_list.labelname {mkey}
```
