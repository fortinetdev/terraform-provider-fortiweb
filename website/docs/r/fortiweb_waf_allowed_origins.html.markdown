---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_allowed_origins"
description: |-
  Configure FortiWEB waf allowed origins configuration.
---

# fortiweb_waf_allowed_origins
Configure FortiWEB waf allowed origins configuration.

## Example Usage
```hcl
resource "fortiweb_waf_allowed_origins" "test" {
	mkey = "test"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf allowed origins entry.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf Allowed Origins can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_allowed_origins.labelname {mkey}
```
