---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_custom_protection_group_type_list"
description: |-
  Configure FortiWEB waf custom protection group type list configuration.
---

# fortiweb_waf_custom_protection_group_type_list
Configure FortiWEB waf custom protection group type list configuration.

## Example Usage
```hcl
resource "fortiweb_waf_custom_protection_group_type_list" "test" {
	mkey = "test"
	sub_mkey = "1"
	custom_protection_rule = "test"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf custom protection group.
* `sub_mkey` - The id of the waf custom protection group type list entry.
* `custom_protection_rule` - The name of the custom protection rule to associate with the custom protection group.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf Custom Protection Group Type List can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_custom_protection_group_type_list.labelname {mkey}
```
