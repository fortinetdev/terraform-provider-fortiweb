---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_url_access_url_access_policy_rule"
description: |-
  Configure FortiWEB waf url access url access policy rule configuration.
---

# fortiweb_waf_url_access_url_access_policy_rule
Configure FortiWEB waf url access url access policy rule configuration.

## Example Usage
```hcl
resource "fortiweb_waf_url_access_url_access_policy_rule" "test" {
	sub_mkey = "1"
	mkey = "test"
	url_access_rule_name = "test"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf url access url access policy entry.
* `sub_mkey` - The id of the waf url access url access policy rule entry.
* `url_access_rule_name` - Select the name of the waf url access url access policy rule entry.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf Url Access Url Access Policy Rule can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_url_access_url_access_policy_rule.labelname {mkey}
```
