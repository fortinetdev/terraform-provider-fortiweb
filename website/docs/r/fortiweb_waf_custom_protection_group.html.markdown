---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_custom_protection_group"
description: |-
  Configure FortiWEB waf custom protection group configuration.
---

# fortiweb_waf_custom_protection_group
Configure FortiWEB waf custom protection group configuration.

## Example Usage
```hcl
resource "fortiweb_waf_custom_protection_group" "test" {
	mkey = "test"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf custom protection group entry.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf Custom Protection Group can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_custom_protection_group.labelname {mkey}
```
