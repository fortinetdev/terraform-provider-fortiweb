---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_cors_protection_rule_allowed_methods_list"
description: |-
  Configure FortiWEB waf cors protection rule allowed methods list configuration.
---

# fortiweb_waf_cors_protection_rule_allowed_methods_list
Configure FortiWEB waf cors protection rule allowed methods list configuration.

## Example Usage
```hcl
resource "fortiweb_waf_cors_protection_rule_allowed_methods_list" "test" {
	sub_mkey = "1"
	mkey = "test"
	method = "post"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf cors protection rule.
* `sub_mkey` - The ID of the waf cors protection rule allowed methods list entry.
* `method` - The name of allowed method.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf Cors Protection Rule Allowed Methods List can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_cors_protection_rule_allowed_methods_list.labelname {mkey}
```
